package seeds

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"gorm.io/gorm"
)

const (
	MinAddressID       = 1
	MaxAddressID       = 100
	MinOrderTotalPrice = 1
	MaxOrderTotalPrice = 10000
	MinDiscount        = 1
	MaxDiscount        = 100
	MinPaymentMethodID = 1
	MaxPaymentMethodID = 10
	MinCount           = 1
	MaxCount           = 10
)

func CreateOrderDomain(db *gorm.DB, options ...func(*domain.ShopOrder)) error {
	gofakeit.Seed(0)

	user := domain.User{}
	if err := db.Order("RANDOM()").Take(&user).Error; err != nil {
		return err
	}
	address := domain.Address{}
	if err := db.Order("RANDOM()").Take(&address).Error; err != nil {
		return err
	}
	productItem := domain.ProductItem{}
	if err := db.Order("RANDOM()").Take(&productItem).Error; err != nil {
		return err
	}

	variation := domain.Variation{}
	if err := db.Order("RANDOM()").Take(&variation).Error; err != nil {
		return err
	}

	variationOption := domain.VariationOption{}
	if err := db.Order("RANDOM()").Take(&variationOption).Error; err != nil {
		return err
	}

	paymentMethod := domain.PaymentMethod{}
	if err := db.Order("RANDOM()").Take(&paymentMethod).Error; err != nil {
		return err
	}
	// order domain
	order := domain.ShopOrder{
		UserID:          user.ID,
		OrderDate:       gofakeit.Date(),
		AddressID:       address.ID,
		OrderTotalPrice: uint(gofakeit.Number(MinOrderTotalPrice, MaxOrderTotalPrice)),
		Discount:        uint(gofakeit.Number(MinDiscount, MaxDiscount)),
		PaymentMethodID: paymentMethod.ID,
	}
	if err := db.Create(&order).Error; err != nil {
		return err
	}

	// ShopOrderProductItem domain
	shopOrderProductItem := domain.ShopOrderProductItem{
		ShopOrderID:   order.ID,
		ProductItemID: productItem.ID,
		Count:         uint(gofakeit.Number(MinCount, MaxCount)),
	}
	if err := db.Create(&shopOrderProductItem).Error; err != nil {
		return err
	}

	// ShopOrderVariation domain
	shopOrderVariation := domain.ShopOrderVariation{
		ShopOrderID:       order.ID,
		VariationID:       variation.ID,
		VariationOptionID: variationOption.ID,
	}
	if err := db.Create(&shopOrderVariation).Error; err != nil {
		return err
	}

	return nil
}

func CreateOrdersDomain(db *gorm.DB, count int) error {
	for i := 0; i < count; i++ {
		if err := CreateOrderDomain(db); err != nil {
			return err
		}
	}

	return nil
}
