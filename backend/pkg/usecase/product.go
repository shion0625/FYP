package usecase

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	"github.com/shion0625/FYP/backend/pkg/service/cloud"
	usecaseInterfaces "github.com/shion0625/FYP/backend/pkg/usecase/interfaces"
	"github.com/shion0625/FYP/backend/pkg/utils"
)

type productUseCase struct {
	productRepo  interfaces.ProductRepository
	cloudService cloud.CloudService
}

// to get a new instance of productUseCase.
func NewProductUseCase(productRepo interfaces.ProductRepository, cloudService cloud.CloudService) usecaseInterfaces.ProductUseCase {
	return &productUseCase{
		productRepo:  productRepo,
		cloudService: cloudService,
	}
}

const numGoroutines = 2

func (p *productUseCase) FindAllCategories(ctx echo.Context, pagination request.Pagination) ([]response.Category, error) {
	categories, err := p.productRepo.FindAllMainCategories(ctx, pagination)
	if err != nil {
		return nil, fmt.Errorf("unable to find all main categories: %w", err)
	}

	return categories, nil
}

// Save category.
func (p *productUseCase) SaveCategory(ctx echo.Context, categoryName string) error {
	categoryExist, err := p.productRepo.IsCategoryNameExist(ctx, categoryName)
	if err != nil {
		return fmt.Errorf("unable to check if category already exists: %w", err)
	}

	if categoryExist {
		return ErrCategoryAlreadyExist
	}

	err = p.productRepo.SaveCategory(ctx, categoryName)
	if err != nil {
		return fmt.Errorf("unable to save category: %w", err)
	}

	return nil
}

// to add new variation for a category.
func (p *productUseCase) SaveVariation(ctx echo.Context, categoryID uint, variationNames []string) error {
	err := p.productRepo.Transactions(ctx, func(repo interfaces.ProductRepository) error {
		for _, variationName := range variationNames {
			variationExist, err := repo.IsVariationNameExistForCategory(ctx, variationName, categoryID)
			if err != nil {
				return fmt.Errorf("unable to check if variation already exists: %w", err)
			}

			if variationExist {
				return fmt.Errorf("variation name %s: %w", variationName, ErrVariationAlreadyExist)
			}

			err = p.productRepo.SaveVariation(ctx, categoryID, variationName)
			if err != nil {
				return fmt.Errorf("unable to save variation: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("unable to save variation: %w", err)
	}

	return nil
}

// to add new variation value for variation.
func (p *productUseCase) SaveVariationOption(ctx echo.Context, variationID uint, variationOptionValues []string) error {
	err := p.productRepo.Transactions(ctx, func(repo interfaces.ProductRepository) error {
		for _, variationValue := range variationOptionValues {
			valueExist, err := repo.IsVariationValueExistForVariation(ctx, variationValue, variationID)
			if err != nil {
				return fmt.Errorf("unable to check if variation already exists: %w", err)
			}
			if valueExist {
				return fmt.Errorf("variation option value %s: %w", variationValue, ErrVariationOptionAlreadyExist)
			}

			err = repo.SaveVariationOption(ctx, variationID, variationValue)
			if err != nil {
				return fmt.Errorf("unable to save variation option: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("unable to save variation option: %w", err)
	}

	return nil
}

func (p *productUseCase) FindAllVariationsAndItsValues(ctx echo.Context, categoryID uint) ([]response.Variation, error) {
	variations, err := p.productRepo.FindAllVariationsByCategoryID(ctx, categoryID)
	if err != nil {
		return nil, fmt.Errorf("unable to find all variations of category: %w", err)
	}

	// get all variation values of each variations
	for i, variation := range variations {
		variationOption, err := p.productRepo.FindAllVariationOptionsByVariationID(ctx, variation.ID)
		if err != nil {
			return nil, fmt.Errorf("unable to get variation option: %w", err)
		}

		variations[i].VariationOptions = variationOption
	}

	return variations, nil
}

// to get all product.
func (p *productUseCase) FindAllProducts(ctx echo.Context, pagination request.Pagination, categoryID *uint, brandID *uint) ([]response.Product, error) {
	products, err := p.productRepo.FindAllProducts(ctx, pagination, categoryID, brandID)
	if err != nil {
		return nil, fmt.Errorf("unable to get product details from database: %w", err)
	}

	for i := range products {
		url, err := p.cloudService.GetFileUrl(ctx, products[i].Image)
		if err != nil {
			continue
		}

		products[i].Image = url
	}

	return products, nil
}

func (p *productUseCase) GetProduct(ctx echo.Context, productID uint) (response.Product, error) {
	product, err := p.productRepo.FindProductByID(ctx, productID)
	if err != nil {
		return response.Product{}, fmt.Errorf("unable to get product from database: %w", err)
	}

	url, err := p.cloudService.GetFileUrl(ctx, product.Image)
	if err != nil {
		return response.Product{}, fmt.Errorf("unable to get image url from cloud service: %w", err)
	}

	product.Image = url

	return product, nil
}

// to add new product.
func (p *productUseCase) SaveProduct(ctx echo.Context, product request.Product) error {
	productNameExist, err := p.productRepo.IsProductNameExist(ctx, product.Name)
	if err != nil {
		return fmt.Errorf("unable to check if product name already exists: %w", err)
	}

	if productNameExist {
		return fmt.Errorf("product name %s: %w", product.Name, ErrProductAlreadyExist)
	}

	uploadID, err := p.cloudService.SaveFile(ctx, product.ImageFileHeader)
	if err != nil {
		return fmt.Errorf("unable to save image on cloud storage: %w", err)
	}

	err = p.productRepo.SaveProduct(ctx, domain.Product{
		Name:        product.Name,
		Description: product.Description,
		CategoryID:  product.CategoryID,
		BrandID:     product.BrandID,
		Price:       product.Price,
		Image:       uploadID,
	})
	if err != nil {
		return fmt.Errorf("unable to save product: %w", err)
	}

	return nil
}

// for add new productItem for a specific product.
func (p *productUseCase) SaveProductItem(ctx echo.Context, productID uint, productItem request.ProductItem) error {
	variationCount, err := p.productRepo.FindVariationCountForProduct(ctx, productID)
	if err != nil {
		return fmt.Errorf("unable to get variation count of product from database: %w", err)
	}

	if len(productItem.VariationOptionIDs) != int(variationCount) {
		return ErrNotEnoughVariations
	}

	// check the given all combination already exist (Color:Red with Size:M)
	productItemExist, err := p.isProductVariationCombinationExist(productID, productItem.VariationOptionIDs)
	if err != nil {
		return err
	}

	if productItemExist {
		return ErrProductItemAlreadyExist
	}

	err = p.productRepo.Transactions(ctx, func(trxRepo interfaces.ProductRepository) error {
		sku, err := utils.GenerateSKU()
		if err != nil {
			return fmt.Errorf("unable to generate SKU: %w", err)
		}

		newProductItem := domain.ProductItem{
			ProductID:  productID,
			QtyInStock: productItem.QtyInStock,
			Price:      productItem.Price,
			SKU:        sku,
		}

		productItemID, err := trxRepo.SaveProductItem(ctx, newProductItem)
		if err != nil {
			return fmt.Errorf("unable to save product item: %w", err)
		}

		errChan := make(chan error, numGoroutines)

		newCtx, cancel := context.WithCancel(ctx.Request().Context()) // for any of one of goroutine get error then cancel the working of other also
		defer cancel()

		go func() {
			// save all product configurations based on given variation option id
			for _, variationOptionID := range productItem.VariationOptionIDs {
				select {
				case <-newCtx.Done():
					return
				default:
					err = trxRepo.SaveProductConfiguration(ctx, productItemID, variationOptionID)
					if err != nil {
						errChan <- fmt.Errorf("unable to save product_item configuration: %w", err)

						return
					}
				}
			}
			errChan <- nil
		}()

		go func() {
			// save all images for the given product item
			for _, imageFile := range productItem.ImageFileHeaders {
				select {
				case <-newCtx.Done():
					return
				default:
					// upload image on cloud
					uploadID, err := p.cloudService.SaveFile(ctx, imageFile)
					if err != nil {
						errChan <- fmt.Errorf("unable to upload image to cloud: %w", err)

						return
					}
					// save upload id on database
					err = trxRepo.SaveProductItemImage(ctx, productItemID, uploadID)
					if err != nil {
						errChan <- fmt.Errorf("unable to save image for product item on database: %w", err)

						return
					}
				}
			}
			errChan <- nil
		}()

		// wait for the both go routine to complete
		for i := 1; i <= 2; i++ {
			select {
			case <-ctx.Request().Context().Done():
				return nil
			case err := <-errChan:
				if err != nil { // if any of the goroutine send error then return the error
					return err
				}
				// no error then continue for the next check of select
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("unable to save product item: %w", err)
	}

	return nil
}

func (p *productUseCase) isProductVariationCombinationExist(productID uint, variationOptionIDs []uint) (exist bool, err error) {
	setOfIds := map[uint]int{}

	for _, variationOptionID := range variationOptionIDs {
		productItemIds, err := p.productRepo.FindAllProductItemIDsByProductIDAndVariationOptionID(context.TODO(),
			productID, variationOptionID)
		if err != nil {
			return false, fmt.Errorf("unable to find product item ids from database using product id and variation option id: %w", err)
		}

		if len(productItemIds) == 0 {
			return false, nil
		}

		for _, productItemID := range productItemIds {
			setOfIds[productItemID]++
			// if any of the ids count is equal to array length it means product item id of this is the existing product item of this configuration
			if setOfIds[productItemID] >= len(variationOptionIDs) {
				return true, nil
			}
		}
	}

	return false, nil
}

// for get all productItem for a specific product.
func (p *productUseCase) FindAllProductItems(ctx echo.Context, productID uint) ([]response.ProductItems, error) {
	productItems, err := p.productRepo.FindAllProductItems(ctx, productID)
	completeProductItems := make([]response.ProductItems, len(productItems))

	for i, item := range productItems {
		completeProductItems[i] = response.ProductItems{
			ID:               item.ID,
			Name:             item.Name,
			ItemName:         item.ItemName,
			Price:            item.Price,
			DiscountPrice:    item.DiscountPrice,
			SKU:              item.SKU,
			QtyInStock:       item.QtyInStock,
			CategoryName:     item.CategoryName,
			MainCategoryName: item.MainCategoryName,
			BrandID:          item.BrandID,
			BrandName:        item.BrandName,
			Images:           item.Images,
			VariationValues:  []response.ProductVariationValue{},
		}
	}

	if err != nil {
		return completeProductItems, fmt.Errorf("unable to find all product items: %w", err)
	}

	errChan := make(chan error, numGoroutines)
	newCtx, cancel := context.WithCancel(ctx.Request().Context())

	defer cancel()

	go func() {
		// get all variation values of each product items
		for i := range completeProductItems {
			select { // checking each time echo is cancelled or not
			case <-ctx.Request().Context().Done():
				return
			default:
				variationValues, err := p.productRepo.FindAllVariationValuesOfProductItem(ctx, completeProductItems[i].ID)
				if err != nil {
					errChan <- fmt.Errorf("unable to find variation values product item: %w", err)

					return
				}

				completeProductItems[i].VariationValues = variationValues
			}
		}
		errChan <- nil
	}()

	go func() {
		// get all images of each product items
		for i := range completeProductItems {
			select { // checking each time echo is cancelled or not
			case <-newCtx.Done():
				return
			default:
				images, err := p.productRepo.FindAllProductItemImages(ctx, completeProductItems[i].ID)

				imageUrls := make([]string, len(images))

				for j := range images {
					url, err := p.cloudService.GetFileUrl(ctx, images[j])
					if err != nil {
						errChan <- fmt.Errorf("unable to get image url from cloud service: %w", err)
					}

					imageUrls[j] = url
				}

				if err != nil {
					errChan <- fmt.Errorf("unable to find images of product item: %w", err)

					return
				}

				completeProductItems[i].Images = imageUrls
			}
		}
		errChan <- nil
	}()

	// wait for the two routine to complete
	for i := 1; i <= 2; i++ {
		select {
		case <-ctx.Request().Context().Done():
			return nil, nil
		case err := <-errChan:
			// no error then continue for the next check
			if err != nil {
				return nil, err
			}
		}
	}

	return completeProductItems, nil
}

func (p *productUseCase) UpdateProduct(ctx echo.Context, updateDetails domain.Product) error {
	nameExistForOther, err := p.productRepo.IsProductNameExistForOtherProduct(ctx, updateDetails.Name, updateDetails.ID)
	if err != nil {
		return fmt.Errorf("unable to check if product name already exists for other product: %w", err)
	}

	if nameExistForOther {
		return fmt.Errorf("product name %s: %w", updateDetails.Name, ErrProductAlreadyExist)
	}

	// p.productRepo.FindProductByID(ctx, updateDetails.ID)

	err = p.productRepo.UpdateProduct(ctx, updateDetails)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	return nil
}
