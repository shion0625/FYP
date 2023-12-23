package handler

import (
	"errors"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/interfaces"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/usecase"
	usecaseInterface "github.com/shion0625/FYP/backend/pkg/usecase/interfaces"
	"github.com/shion0625/FYP/backend/pkg/utils"
)

type ProductHandler struct {
	productUseCase usecaseInterface.ProductUseCase
}

func NewProductHandler(productUsecase usecaseInterface.ProductUseCase) interfaces.ProductHandler {
	return &ProductHandler{
		productUseCase: productUsecase,
	}
}

func (p *ProductHandler) GetAllCategories(ctx echo.Context) error {
	pagination := request.GetPagination(ctx)

	categories, err := p.productUseCase.FindAllCategories(ctx, pagination)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to retrieve categories", err, nil)

		return nil
	}

	if len(categories) == 0 {
		response.SuccessResponse(ctx, http.StatusOK, "No categories found", nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully retrieved all categories", categories)

	return nil
}

func (p *ProductHandler) SaveCategory(ctx echo.Context) error {
	var body request.Category

	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, nil)

		return nil
	}

	if err := ctx.Validate(body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request data", err, nil)

		return nil
	}

	err := p.productUseCase.SaveCategory(ctx, body.Name)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, usecase.ErrCategoryAlreadyExist) {
			statusCode = http.StatusConflict
		}

		response.ErrorResponse(ctx, statusCode, "Failed to add category", err, nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusCreated, "Successfully added category", nil)

	return nil
}

func (p *ProductHandler) SaveVariation(ctx echo.Context) error {
	categoryIDStr := ctx.Param("category_id")

	categoryID, err := utils.ParseStringToUint32(categoryIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindParamFailMessage, err, nil)

		return nil
	}

	var body request.Variation

	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, nil)

		return nil
	}

	if err := ctx.Validate(body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request data", err, nil)

		return nil
	}

	err = p.productUseCase.SaveVariation(ctx, categoryID, body.Names)

	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, usecase.ErrVariationAlreadyExist) {
			statusCode = http.StatusConflict
		}

		response.ErrorResponse(ctx, statusCode, "Failed to add variation", err, nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusCreated, "Successfully added variations", nil)

	return nil
}

func (p *ProductHandler) SaveVariationOption(ctx echo.Context) error {
	variationIDStr := ctx.Param("variation_id")

	variationID, err := utils.ParseStringToUint32(variationIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindParamFailMessage, err, nil)

		return nil
	}

	var body request.VariationOption

	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, nil)

		return nil
	}

	if err := ctx.Validate(body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request data", err, nil)

		return nil
	}

	err = p.productUseCase.SaveVariationOption(ctx, variationID, body.Values)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, usecase.ErrVariationOptionAlreadyExist) {
			statusCode = http.StatusConflict
		}

		response.ErrorResponse(ctx, statusCode, "Failed to add variation options", err, nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusCreated, "Successfully added variation options", nil)

	return nil
}

func (c *ProductHandler) GetAllVariations(ctx echo.Context) error {
	categoryIDStr := ctx.Param("category_id")

	categoryID, err := utils.ParseStringToUint32(categoryIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindParamFailMessage, err, nil)

		return nil
	}

	variations, err := c.productUseCase.FindAllVariationsAndItsValues(ctx, categoryID)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to Get variations and its values", err, nil)

		return nil
	}

	if len(variations) == 0 {
		response.SuccessResponse(ctx, http.StatusOK, "No variations found", nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully retrieved all variations and its values", variations)

	return nil
}

func (p *ProductHandler) SaveProduct(ctx echo.Context) error {
	var body request.Product

	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, nil)

		return nil
	}

	if err := ctx.Validate(body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request data", err, nil)

		return nil
	}

	err := p.productUseCase.SaveProduct(ctx, body)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, usecase.ErrProductAlreadyExist) {
			statusCode = http.StatusConflict
		}

		response.ErrorResponse(ctx, statusCode, "Failed to add product", err, nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusCreated, "Successfully product added", nil)

	return nil
}

func (p *ProductHandler) GetAllProductsAdmin() func(ctx echo.Context) error {
	return p.getAllProducts()
}

func (p *ProductHandler) GetAllProductsUser() func(ctx echo.Context) error {
	return p.getAllProducts()
}

// Get products is common for user and admin so this function is to get the common Get all products func for them.
func (p *ProductHandler) getAllProducts() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		categoryID, err := utils.GetParamID(ctx, "category_id")
		if err != nil {
			response.ErrorResponse(ctx, http.StatusBadRequest, BindParamFailMessage, err, nil)

			return nil
		}

		brandID, err := utils.GetParamID(ctx, "brand_id")
		if err != nil {
			response.ErrorResponse(ctx, http.StatusBadRequest, BindParamFailMessage, err, nil)

			return nil
		}

		pagination := request.GetPagination(ctx)

		products, err := p.productUseCase.FindAllProducts(ctx, pagination, categoryID, brandID)
		if err != nil {
			response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to Get all products", err, nil)

			return nil
		}

		if len(products) == 0 {
			response.SuccessResponse(ctx, http.StatusOK, "No products found", nil)

			return nil
		}

		response.SuccessResponse(ctx, http.StatusOK, "Successfully found all products", products)

		return nil
	}
}

func (p *ProductHandler) GetProduct(ctx echo.Context) error {
	productIDStr := ctx.Param("product_id")

	productID, err := utils.ParseStringToUint32(productIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindParamFailMessage, err, nil)

		return nil
	}

	products, err := p.productUseCase.GetProduct(ctx, productID)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to Get all products", err, nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully found all products", products)

	return nil
}

func (c *ProductHandler) UpdateProduct(ctx echo.Context) error {
	var body request.UpdateProduct

	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, nil)

		return nil
	}

	if err := ctx.Validate(body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request data", err, nil)

		return nil
	}

	var product domain.Product
	if err := copier.Copy(&product, &body); err != nil {
		response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to copy product data", err, nil)

		return nil
	}

	if err := c.productUseCase.UpdateProduct(ctx, product); err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, usecase.ErrProductAlreadyExist) {
			statusCode = http.StatusConflict
		}

		response.ErrorResponse(ctx, statusCode, "Failed to update product", err, nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusOK, "Successfully product updated", nil)

	return nil
}

func (p *ProductHandler) SaveProductItem(ctx echo.Context) error {
	productIDStr := ctx.Param("product_id")

	productID, err := utils.ParseStringToUint32(productIDStr)
	if err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindParamFailMessage, err, nil)

		return nil
	}

	var body request.ProductItem

	if err := ctx.Bind(&body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, BindJsonFailMessage, err, nil)

		return nil
	}

	if err := ctx.Validate(body); err != nil {
		response.ErrorResponse(ctx, http.StatusBadRequest, "Invalid request data", err, nil)

		return nil
	}

	err = p.productUseCase.SaveProductItem(ctx, productID, body)

	if err != nil {
		var statusCode int

		switch {
		case errors.Is(err, usecase.ErrProductItemAlreadyExist):
			statusCode = http.StatusConflict
		case errors.Is(err, usecase.ErrNotEnoughVariations):
			statusCode = http.StatusBadRequest
		default:
			statusCode = http.StatusInternalServerError
		}

		response.ErrorResponse(ctx, statusCode, "Failed to add product item", err, nil)

		return nil
	}

	response.SuccessResponse(ctx, http.StatusCreated, "Successfully product item added", nil)

	return nil
}

func (p *ProductHandler) GetAllProductItemsAdmin() func(ctx echo.Context) error {
	return p.getAllProductItems()
}

func (p *ProductHandler) GetAllProductItemsUser() func(ctx echo.Context) error {
	return p.getAllProductItems()
}

// same functionality of get all product items for admin and user.
func (p *ProductHandler) getAllProductItems() func(ctx echo.Context) error {
	return func(ctx echo.Context) error {
		productIDStr := ctx.Param("product_id")

		productID, err := utils.ParseStringToUint32(productIDStr)
		if err != nil {
			response.ErrorResponse(ctx, http.StatusBadRequest, BindParamFailMessage, err, nil)

			return nil
		}

		productItems, err := p.productUseCase.FindAllProductItems(ctx, productID)
		if err != nil {
			response.ErrorResponse(ctx, http.StatusInternalServerError, "Failed to get all product items", err, nil)

			return nil
		}

		// check the product have productItem exist or not
		if len(productItems) == 0 {
			response.SuccessResponse(ctx, http.StatusOK, "No product items found", nil)

			return nil
		}

		response.SuccessResponse(ctx, http.StatusOK, "Successfully get all product items ", productItems)

		return nil
	}
}
