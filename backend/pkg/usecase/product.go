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
		return nil, fmt.Errorf("failed find all main categories %w", err)
	}

	return categories, nil
}

// Save category.
func (p *productUseCase) SaveCategory(ctx echo.Context, categoryName string) error {
	categoryExist, err := p.productRepo.IsCategoryNameExist(ctx, categoryName)
	if err != nil {
		return fmt.Errorf("failed to check category already exist: %w", err)
	}

	if categoryExist {
		return ErrCategoryAlreadyExist
	}

	err = p.productRepo.SaveCategory(ctx, categoryName)
	if err != nil {
		return fmt.Errorf("failed to save category: %w", err)
	}

	return nil
}

// to add new variation for a category.
func (p *productUseCase) SaveVariation(ctx echo.Context, categoryID uint, variationNames []string) error {
	err := p.productRepo.Transactions(ctx, func(repo interfaces.ProductRepository) error {
		for _, variationName := range variationNames {
			variationExist, err := repo.IsVariationNameExistForCategory(ctx, variationName, categoryID)
			if err != nil {
				return fmt.Errorf("failed to check variation already exist: %w", err)
			}

			if variationExist {
				return fmt.Errorf("variation name %s: %w", variationName, ErrVariationAlreadyExist)
			}

			err = p.productRepo.SaveVariation(ctx, categoryID, variationName)
			if err != nil {
				return fmt.Errorf("failed to save variation: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to save variation: %w", err)
	}

	return nil
}

// to add new variation value for variation.
func (p *productUseCase) SaveVariationOption(ctx echo.Context, variationID uint, variationOptionValues []string) error {
	err := p.productRepo.Transactions(ctx, func(repo interfaces.ProductRepository) error {
		for _, variationValue := range variationOptionValues {
			valueExist, err := repo.IsVariationValueExistForVariation(ctx, variationValue, variationID)
			if err != nil {
				return fmt.Errorf("failed to check variation already exist: %w", err)
			}
			if valueExist {
				return fmt.Errorf("variation option value %s: %w", variationValue, ErrVariationOptionAlreadyExist)
			}

			err = repo.SaveVariationOption(ctx, variationID, variationValue)
			if err != nil {
				return fmt.Errorf("failed to save variation option: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to save variation option: %w", err)
	}

	return nil
}

func (p *productUseCase) FindAllVariationsAndItsValues(ctx echo.Context, categoryID uint) ([]response.Variation, error) {
	variations, err := p.productRepo.FindAllVariationsByCategoryID(ctx, categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to find all variations of category: %w", err)
	}

	// get all variation values of each variations
	for i, variation := range variations {
		variationOption, err := p.productRepo.FindAllVariationOptionsByVariationID(ctx, variation.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to get variation option: %w", err)
		}

		variations[i].VariationOptions = variationOption
	}

	return variations, nil
}

// to get all product.
func (p *productUseCase) FindAllProducts(ctx echo.Context, pagination request.Pagination, categoryID *uint, brandID *uint) ([]response.Product, error) {
	products, err := p.productRepo.FindAllProducts(ctx, pagination, categoryID, brandID)
	if err != nil {
		return nil, fmt.Errorf("failed to get product details from database: %w", err)
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
		return response.Product{}, fmt.Errorf("failed to get product from database: %w", err)
	}

	url, err := p.cloudService.GetFileUrl(ctx, product.Image)

	if err != nil {
		return response.Product{}, fmt.Errorf("failed to get image url from could service: %w", err)
	}

	product.Image = url

	return product, nil
}

// to add new product.
func (p *productUseCase) SaveProduct(ctx echo.Context, product request.Product) error {
	productNameExist, err := p.productRepo.IsProductNameExist(ctx, product.Name)
	if err != nil {
		return fmt.Errorf("failed to check product name already exist: %w", err)
	}

	if productNameExist {
		return fmt.Errorf("product name %s: %w", product.Name, ErrProductAlreadyExist)
	}

	uploadID, err := p.cloudService.SaveFile(ctx, product.ImageFileHeader)
	if err != nil {
		return fmt.Errorf("failed to save image on cloud storage: %w", err)
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
		return fmt.Errorf("failed to save product: %w", err)
	}

	return nil
}

// for add new productItem for a specific product.
func (p *productUseCase) SaveProductItem(ctx echo.Context, productID uint, productItem request.ProductItem) error {
	variationCount, err := p.productRepo.FindVariationCountForProduct(ctx, productID)
	if err != nil {
		return fmt.Errorf("failed to get variation count of product from database: %w", err)
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
			return fmt.Errorf("failed to generate SKU: %w", err)
		}

		newProductItem := domain.ProductItem{
			ProductID:  productID,
			QtyInStock: productItem.QtyInStock,
			Price:      productItem.Price,
			SKU:        sku,
		}

		productItemID, err := trxRepo.SaveProductItem(ctx, newProductItem)
		if err != nil {
			return fmt.Errorf("failed to save product item: %w", err)
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
						errChan <- fmt.Errorf("failed to save product_item configuration: %w", err)

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
						errChan <- fmt.Errorf("failed to upload image to cloud: %w", err)

						return
					}
					// save upload id on database
					err = trxRepo.SaveProductItemImage(ctx, productItemID, uploadID)
					if err != nil {
						errChan <- fmt.Errorf("failed to save image for product item on database: %w", err)

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
		return fmt.Errorf("failed to save product item: %w", err)
	}

	return nil
}

// step 1 : get product_id and all variation id as function parameter
// step 2 : initialize an map for storing product item id and its count(map[uint]int)
// step 3 : loop through the variation option ids
// step 4 : then find all product items ids with given product id and the loop variation option id
// step 5 : if the product item array length is zero means the configuration not exist return false
// step 6 : then loop through the product items ids array(got from database)
// step 7 : add each id on the map and increment its count
// step 8 : check if any of the product items id's count is greater than the variation options ids length then return true
// step 9 : if the loop exist means product configuration is not exist.
func (p *productUseCase) isProductVariationCombinationExist(productID uint, variationOptionIDs []uint) (exist bool, err error) {
	setOfIds := map[uint]int{}

	for _, variationOptionID := range variationOptionIDs {
		productItemIds, err := p.productRepo.FindAllProductItemIDsByProductIDAndVariationOptionID(context.TODO(),
			productID, variationOptionID)
		if err != nil {
			return false, fmt.Errorf("failed to find product item ids from database using product id and variation option id: %w", err)
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
			ProductID:        item.ProductID,
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
		return completeProductItems, fmt.Errorf("failed to find all product items: %w", err)
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
					errChan <- fmt.Errorf("failed to find variation values product item: %w", err)

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
					fmt.Println(images[j])
					url, err := p.cloudService.GetFileUrl(ctx, images[j])
					if err != nil {
						errChan <- fmt.Errorf("failed to get image url from could service: %w", err)
					}

					imageUrls[j] = url
				}

				if err != nil {
					errChan <- fmt.Errorf("failed to find images of product item: %w", err)

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
		return fmt.Errorf("failed to check product name already exist for other product: %w", err)
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
