package repository

import (
	"fmt"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type orderDatabase struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &orderDatabase{
		DB: db,
	}
}

func (c *orderDatabase) Transactions(ctx echo.Context, trxFn func(repo interfaces.OrderRepository) error) error {
	trx := c.DB.Begin()

	repo := NewOrderRepository(trx)

	if err := trxFn(repo); err != nil {
		trx.Rollback()

		return err
	}

	if err := trx.Commit().Error; err != nil {
		trx.Rollback()

		return err
	}

	return nil
}

func (c *orderDatabase) UpdateProductItemStock(ctx echo.Context, productItemID uint, purchaseQuantity uint) (newStock uint, err error) {
	var productItem domain.ProductItem
	err = c.DB.Table("product_items").Select("qty_in_stock").Where("id = ?", productItemID).First(&productItem).Error

	if err != nil {
		return 0, err
	}

	if productItem.QtyInStock < purchaseQuantity {
		return 0, fmt.Errorf("%s not enough stock", productItem.Name)
	}

	newStock = productItem.QtyInStock - purchaseQuantity
	err = c.DB.Table("product_items").Where("id = ?", productItem.ID).Update("qty_in_stock", newStock).Error

	return newStock, err
}

func (c *orderDatabase) SaveOrder(ctx echo.Context, userID string, payOrder request.PayOrder) error {
	// Create a new ShopOrder from the PayOrder request
	shopOrder := domain.ShopOrder{
		UserID:          userID,
		OrderDate:       time.Now(),
		OrderTotalPrice: payOrder.TotalFee,
		AddressID:       payOrder.AddressID,
		PaymentMethodID: payOrder.PaymentMethodID,
	}

	// Save the ShopOrder to the database
	if err := c.DB.Create(&shopOrder).Error; err != nil {
		return err
	}

	shopOrderProductItems := make([]domain.ShopOrderProductItem, 0, len(payOrder.ProductItemInfo))

	for _, productItemInfo := range payOrder.ProductItemInfo {
		shopOrderProductItem := domain.ShopOrderProductItem{
			ShopOrderID:   shopOrder.ID,
			ProductItemID: productItemInfo.ProductItemID,
			Count:         productItemInfo.Count,
		}
		shopOrderProductItems = append(shopOrderProductItems, shopOrderProductItem)

		shopOrderVariations := make([]domain.ShopOrderVariation, 0, len(*productItemInfo.VariationValues))

		for _, variationValue := range *productItemInfo.VariationValues {
			shopOrderVariation := domain.ShopOrderVariation{
				ShopOrderID:       shopOrder.ID,
				ProductItemID:     productItemInfo.ProductItemID,
				VariationID:       variationValue.VariationID,
				VariationOptionID: variationValue.VariationOptionID,
			}
			shopOrderVariations = append(shopOrderVariations, shopOrderVariation)
		}

		if err := c.DB.Create(&shopOrderVariations).Error; err != nil {
			return err
		}
	}

	if err := c.DB.Create(&shopOrderProductItems).Error; err != nil {
		return err
	}

	return nil
}

func (c *orderDatabase) PayOrder(ctx echo.Context, paymentMethodID uint) error {
	paymentMethod := domain.PaymentMethod{
		ID: paymentMethodID,
	}

	if err := c.DB.First(&paymentMethod).Error; err != nil {
		return err
	}

	// Perform payment processing using the PaymentMethod
	// This is a placeholder, replace with actual payment processing logic
	if paymentMethod.CreditNumber == "" {
		return fmt.Errorf("invalid payment method")
	}

	return nil
}

func (o *orderDatabase) GetShopOrders(ctx echo.Context, userID string, pagination request.Pagination) (orders []response.Order, err error) {
	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	var shopOrders []domain.ShopOrder
	if err := o.DB.Preload("Address").Preload("PaymentMethod").Table("shop_orders").Limit(int(limit)).Offset(int(offset)).Where("user_id= ?", userID).Find(&shopOrders).Error; err != nil {
		return nil, err
	}

	fmt.Print(&shopOrders)

	for _, shopOrder := range shopOrders {
		var productItemInfos []response.ProductItemInfo

		var shopOrderProductItems []domain.ShopOrderProductItem
		if err := o.DB.Where("shop_order_id = ?", shopOrder.ID).Find(&shopOrderProductItems).Error; err != nil {
			return nil, err
		}

		for _, productItem := range shopOrderProductItems {
			var variationValues []response.VariationValues

			var shopOrderVariations []domain.ShopOrderVariation
			if err := o.DB.Where("shop_order_id = ? AND product_item_id = ?", shopOrder.ID, productItem.ProductItemID).Find(&shopOrderVariations).Error; err != nil {
				return nil, err
			}

			for _, variation := range shopOrderVariations {
				variationValues = append(variationValues, response.VariationValues{
					VariationID:       variation.VariationID,
					Name:              variation.Variation.Name,
					VariationOptionID: variation.VariationOptionID,
					Value:             variation.VariationOption.Value,
				})
			}

			productItemInfos = append(productItemInfos, response.ProductItemInfo{
				ProductItemID:   productItem.ProductItemID,
				Count:           productItem.Count,
				VariationValues: &variationValues,
			})
		}

		orders = append(orders, response.Order{
			UserID:          shopOrder.UserID,
			ShopOrderId:     strconv.Itoa(int(shopOrder.ID)),
			ProductItemInfo: productItemInfos,
			Address: response.Address{
				ID:   shopOrder.Address.ID,
				Name: shopOrder.Address.Name,
			},
			TotalFee: shopOrder.OrderTotalPrice,
			PaymentMethod: response.PaymentMethod{
				ID:           shopOrder.PaymentMethod.ID,
				CreditNumber: shopOrder.PaymentMethod.CreditNumber,
			},
		})
	}

	//nolint:nakedret
	return
}
