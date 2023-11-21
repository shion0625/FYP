package seeds

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"gorm.io/gorm"
)

func CreateProductDomain(db *gorm.DB, options ...func(*domain.User)) error {
	const (
		MaxSentenceLength = 200
		MaxPrice          = 1000
		ImageSize         = 100
		MaxQtyInStock     = 100
	)

	gofakeit.Seed(0)

	// product
	category := domain.Category{
		Name: gofakeit.Company(),
	}
	if err := db.Create(&category).Error; err != nil {
		return err
	}

	brand := domain.Brand{
		Name: gofakeit.Company(),
	}
	if err := db.Create(&brand).Error; err != nil {
		return err
	}

	product := domain.Product{
		Name:        gofakeit.AppName(),
		Description: gofakeit.Sentence(MaxSentenceLength),
		CategoryID:  category.ID,
		Category:    category,
		BrandID:     brand.ID,
		Brand:       brand,
		Price:       uint(gofakeit.Price(1, MaxPrice)),
		Image:       gofakeit.ImageURL(ImageSize, ImageSize),
	}
	if err := db.Create(&product).Error; err != nil {
		return err
	}

	productItem := domain.ProductItem{
		ProductID:  product.ID,
		Product:    product,
		QtyInStock: uint(gofakeit.Number(1, MaxQtyInStock)),
		Price:      uint(gofakeit.Price(1, MaxPrice)),
		SKU:        gofakeit.UUID(),
	}
	if err := db.Create(&productItem).Error; err != nil {
		return err
	}

	variation := domain.Variation{
		CategoryID: category.ID,
		Category:   category,
		Name:       gofakeit.Color(),
	}
	if err := db.Create(&variation).Error; err != nil {
		return err
	}

	variationOption := domain.VariationOption{
		VariationID: variation.ID,
		Variation:   variation,
		Value:       gofakeit.Color(),
	}
	if err := db.Create(&variationOption).Error; err != nil {
		return err
	}

	productConfiguration := domain.ProductConfiguration{
		ProductItemID:     productItem.ID,
		ProductItem:       productItem,
		VariationOptionID: variationOption.ID,
		VariationOption:   variationOption,
	}
	if err := db.Create(&productConfiguration).Error; err != nil {
		return err
	}

	productImage := domain.ProductImage{
		ProductItemID: productItem.ID,
		ProductItem:   productItem,
		Image:         gofakeit.ImageURL(ImageSize, ImageSize),
	}
	if err := db.Create(&productImage).Error; err != nil {
		return err
	}

	offer := domain.Offer{
		Name:         gofakeit.Company(),
		Description:  gofakeit.Sentence(MaxSentenceLength),
		DiscountRate: uint(gofakeit.Number(1, MaxQtyInStock)),
		StartDate:    gofakeit.Date(),
		EndDate:      gofakeit.Date(),
	}
	if err := db.Create(&offer).Error; err != nil {
		return err
	}

	offerProduct := domain.OfferProduct{
		OfferID:   offer.ID,
		Offer:     offer,
		ProductID: product.ID,
		Product:   product,
	}
	if err := db.Create(&offerProduct).Error; err != nil {
		return err
	}

	return nil
}

func CreateProductsDomain(db *gorm.DB, count int) error {
	for i := 0; i < count; i++ {
		if err := CreateProductDomain(db); err != nil {
			return err
		}
	}

	return nil
}
