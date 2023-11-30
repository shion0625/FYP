package repository

import (
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
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
	err = c.DB.Model(&productItem).Update("qty_in_stock", newStock).Error

	return newStock, err
}

func (c *orderDatabase) SaveOrder(ctx echo.Context, payOrder request.PayOrder) error {
	// Create a new ShopOrder from the PayOrder request
	shopOrder := domain.ShopOrder{
		UserID:          payOrder.UserID,
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
	}

	if err := c.DB.Create(&shopOrderProductItems).Error; err != nil {
		return err
	}

	shopOrderVariations := make([]domain.ShopOrderVariation, 0, len(*payOrder.VariationValue))

	for _, variationValue := range *payOrder.VariationValue {
		shopOrderVariation := domain.ShopOrderVariation{
			ShopOrderID:       shopOrder.ID,
			VariationID:       variationValue.VariationID,
			VariationOptionID: variationValue.VariationOptionID,
		}
		shopOrderVariations = append(shopOrderVariations, shopOrderVariation)
	}

	if err := c.DB.Create(&shopOrderVariations).Error; err != nil {
		return err
	}

	return nil
}

func (c *orderDatabase) PayOrder(ctx echo.Context, payOrder request.PayOrder) error {
	paymentMethod := domain.PaymentMethod{
		ID: payOrder.PaymentMethodID,
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
