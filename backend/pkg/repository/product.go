package repository

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{
		DB: db,
	}
}

func (c *productDatabase) Transactions(ctx echo.Context, trxFn func(repo interfaces.ProductRepository) error) error {
	trx := c.DB.Begin()

	repo := NewProductRepository(trx)

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

// To check the category name exist.
func (c *productDatabase) IsCategoryNameExist(ctx echo.Context, name string) (exist bool, err error) {
	err = c.DB.Table("categories").Where("name = ? AND category_id IS NULL", name).First(&exist).Error

	return
}

// Save Category.
func (c *productDatabase) SaveCategory(ctx echo.Context, categoryName string) (err error) {
	err = c.DB.Table("categories").Create(&domain.Category{Name: categoryName}).Error

	return err
}

// Find all main category(its not have a category_id).
func (c *productDatabase) FindAllMainCategories(ctx echo.Context,
	pagination request.Pagination,
) (categories []response.Category, err error) {
	limit := pagination.Count
	offset := (pagination.PageNumber - 1) * limit

	err = c.DB.Table("categories").Limit(int(limit)).Offset(int(offset)).Find(&categories).Error

	return
}

// Find all variations which related to given category id.
func (c *productDatabase) FindAllVariationsByCategoryID(ctx echo.Context,
	categoryID uint,
) (variations []response.Variation, err error) {
	err = c.DB.Table("variations").Where("category_id = ?", categoryID).Find(&variations).Error

	return
}

// Find all variation options which related to given variation id.
func (c productDatabase) FindAllVariationOptionsByVariationID(ctx echo.Context,
	variationID uint,
) (variationOptions []response.VariationOption, err error) {
	err = c.DB.Table("variation_options").Where("variation_id = ?", variationID).Find(&variationOptions).Error

	return
}

// To check a variation exist for the given category.
func (c *productDatabase) IsVariationNameExistForCategory(ctx echo.Context,
	name string, categoryID uint,
) (exist bool, err error) {
	err = c.DB.Table("variations").Where("name = ? AND category_id = ?", name, categoryID).First(&exist).Error

	return
}

// To check a variation value exist for the given variation.
func (c *productDatabase) IsVariationValueExistForVariation(ctx echo.Context,
	value string, variationID uint,
) (exist bool, err error) {
	err = c.DB.Table("variation_options").Where("value = ? AND variation_id = ?", value, variationID).First(&exist).Error

	return
}

// Save Variation for category.
func (c *productDatabase) SaveVariation(ctx echo.Context, categoryID uint, variationName string) error {
	err := c.DB.Table("variations").Create(&domain.Variation{CategoryID: categoryID, Name: variationName}).Error

	return err
}

// add variation option.
func (c *productDatabase) SaveVariationOption(ctx echo.Context, variationID uint, variationValue string) error {
	err := c.DB.Table("variation_options").Create(&domain.VariationOption{VariationID: variationID, Value: variationValue}).Error

	return err
}

// find product by id.
func (c *productDatabase) FindProductByID(ctx echo.Context, productID uint) (product response.Product, err error) {
	err = c.DB.Table("products").Where("id = ?", productID).Find(&product).Error

	return
}

func (c *productDatabase) IsProductNameExistForOtherProduct(ctx echo.Context,
	name string, productID uint,
) (exist bool, err error) {
	err = c.DB.Table("products").Where("name = ? AND id != ?", name, productID).First(&exist).Error

	return
}

func (c *productDatabase) IsProductNameExist(ctx echo.Context, productName string) (exist bool, err error) {
	err = c.DB.Table("products").Where("name = ?", productName).First(&exist).Error

	return
}

// to add a new product in database.
func (c *productDatabase) SaveProduct(ctx echo.Context, product domain.Product) error {
	product.CreatedAt = time.Now()
	err := c.DB.Table("products").Create(&product).Error

	return err
}

// update product.
func (c *productDatabase) UpdateProduct(ctx echo.Context, product domain.Product) error {
	product.UpdatedAt = time.Now()

	err := c.DB.Table("products").Where("id = ?", product.ID).Updates(&product).Error

	return err
}

// get all products from database.
func (c *productDatabase) FindAllProducts(ctx echo.Context, pagination request.Pagination, categoryID *uint, brandID *uint) (products []response.Product, err error) {
	limit := int(pagination.Count)
	offset := (int(pagination.PageNumber) - 1) * limit

	db := c.DB.Table("products p").
		Select("p.id, p.name, p.description, p.price, p.discount_price, p.image, p.category_id, sc.name AS category_name, p.brand_id, b.name AS brand_name, p.created_at, p.updated_at").
		Joins("INNER JOIN categories sc ON p.category_id = sc.id").
		Joins("INNER JOIN brands b ON b.id = p.brand_id").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset)

	if categoryID != nil {
		db = db.Where("p.category_id = ?", *categoryID)
	}

	if brandID != nil {
		db = db.Where("p.brand_id = ?", *brandID)
	}

	err = db.Scan(&products).Error

	return
}

// to get productItem id.
func (c *productDatabase) FindProductItemByID(ctx echo.Context, productItemID uint) (productItem domain.ProductItem, err error) {
	err = c.DB.Table("product_items").Where("id = ?", productItemID).Find(&productItem).Error

	return productItem, err
}

// to get how many variations are available for a product.
func (c *productDatabase) FindVariationCountForProduct(ctx echo.Context, productID uint) (variationCount uint, err error) {
	var count int64
	err = c.DB.Table("variations v").
		Joins("INNER JOIN categories c ON c.id = v.category_id").
		Joins("INNER JOIN products p ON p.category_id = v.category_id").
		Where("p.id = ?", productID).Count(&count).Error
	variationCount = uint(count)

	return
}

// To find all product item ids which related to the given product id and variation option id.
func (c *productDatabase) FindAllProductItemIDsByProductIDAndVariationOptionID(ctx context.Context, productID,
	variationOptionID uint,
) (productItemIDs []uint, err error) {
	err = c.DB.Table("product_items pi").
		Joins("INNER JOIN product_configurations pc ON pi.id = pc.product_item_id").
		Where("pi.product_id = ? AND variation_option_id = ?", productID, variationOptionID).Find(&productItemIDs).Error

	return
}

func (c *productDatabase) SaveProductConfiguration(ctx echo.Context, productItemID, variationOptionID uint) error {
	err := c.DB.Table("product_configurations").Create(&domain.ProductConfiguration{ProductItemID: productItemID, VariationOptionID: variationOptionID}).Error

	return err
}

func (c *productDatabase) SaveProductItem(ctx echo.Context, productItem domain.ProductItem) (productItemID uint, err error) {
	productItem.CreatedAt = time.Now()
	err = c.DB.Table("product_items").Create(&productItem).Scan(&productItemID).Error

	return
}

// for get all products items for a product.
func (c *productDatabase) FindAllProductItems(ctx echo.Context,
	productID uint,
) (productItems []response.ProductItemsDB, err error) {
	// first find all product_items
	err = c.DB.Table("product_items pi").
		Select("p.name, pi.id,  pi.product_id,pi.name AS item_name, pi.price, pi.discount_price, pi.qty_in_stock, pi.sku, p.category_id, sc.name AS category_name, p.brand_id, b.name AS brand_name").
		Joins("INNER JOIN products p ON p.id = pi.product_id").
		Joins("INNER JOIN categories sc ON p.category_id = sc.id").
		Joins("INNER JOIN brands b ON b.id = p.brand_id").
		Where("pi.product_id = ?", productID).Find(&productItems).Error

	return
}

// Find all variation and value of a product item.
func (c *productDatabase) FindAllVariationValuesOfProductItem(ctx echo.Context,
	productItemID uint,
) (productVariationsValues []response.ProductVariationValue, err error) {
	err = c.DB.Table("product_configurations pc").
		Select("v.id AS variation_id, v.name, vo.id AS variation_option_id, vo.value").
		Joins("INNER JOIN variation_options vo ON vo.id = pc.variation_option_id").
		Joins("INNER JOIN variations v ON v.id = vo.variation_id").
		Where("pc.product_item_id = ?", productItemID).Find(&productVariationsValues).Error

	return
}

// To save image for product item.
func (c *productDatabase) SaveProductItemImage(ctx echo.Context, productItemID uint, image string) error {
	err := c.DB.Table("product_images").Create(&domain.ProductImage{ProductItemID: productItemID, Image: image}).Error

	return err
}

// To find all images of a product item.
func (c *productDatabase) FindAllProductItemImages(ctx echo.Context, productItemID uint) (images []string, err error) {
	err = c.DB.Table("product_images").Select("image").Where("product_item_id = ?", productItemID).Find(&images).Error

	return
}
