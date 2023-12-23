package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	usecaseMock "github.com/shion0625/FYP/backend/pkg/usecase/mock"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func createMultipartFileHeader(fileName string, fileContent string) (*multipart.FileHeader, error) {
	// Create a buffer to store our message
	body := &bytes.Buffer{}

	// Create a multipart writer
	writer := multipart.NewWriter(body)

	// Create a new form-data header with the provided file name
	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}

	// Write the file content to the part
	_, err = part.Write([]byte(fileContent))
	if err != nil {
		return nil, err
	}

	// Populate the file header
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// Create a new request to parse the form data
	req, err := http.NewRequest("POST", "/upload", body)
	if err != nil {
		return nil, err
	}

	// Set the content type of the request
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Parse the form data of the request
	err = req.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		return nil, err
	}

	// Get the file header
	fileHeader := req.MultipartForm.File["file"][0]

	return fileHeader, nil
}

func TestProductHandler_GetAllCategories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := usecaseMock.NewMockProductUseCase(ctrl)

	productHandler := handler.NewProductHandler(mockProductUseCase)

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		prepareMockFn func(m *usecaseMock.MockProductUseCase)
		want          output
	}{
		"Normal Case: GetAllCategories": {
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().FindAllCategories(gomock.Any(), gomock.Any()).Return([]response.Category{
					{
						ID:   1,
						Name: "Test Category",
					},
				}, nil)
			},
			output{http.StatusOK, "Successfully retrieved all categories", nil},
		},
		"Abnormal Case: GetAllCategories": {
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().FindAllCategories(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to retrieve categories", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductUseCase)

			req := httptest.NewRequest(echo.GET, "/getAllCategories", nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)

			err := productHandler.GetAllCategories(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestProductHandler_SaveCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := usecaseMock.NewMockProductUseCase(ctrl)

	productHandler := handler.NewProductHandler(mockProductUseCase)

	type input struct {
		categoryRequest request.Category
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockProductUseCase)
		want          output
	}{
		"Normal Case: SaveCategory": {
			input{
				request.Category{
					Name: "Test Category",
				},
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().SaveCategory(gomock.Any(), gomock.Any()).Return(nil)
			},
			output{http.StatusCreated, "Successfully added category", nil},
		},
		"Abnormal Case: SaveCategory": {
			input{
				request.Category{
					Name: "Test Category",
				},
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().SaveCategory(gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to add category", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductUseCase)

			body, _ := json.Marshal(tt.input.categoryRequest)
			req := httptest.NewRequest(echo.POST, "/saveCategory", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)

			err := productHandler.SaveCategory(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestProductHandler_UpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := usecaseMock.NewMockProductUseCase(ctrl)

	productHandler := handler.NewProductHandler(mockProductUseCase)

	type input struct {
		updateProductRequest request.UpdateProduct
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockProductUseCase)
		want          output
	}{
		"Normal Case: UpdateProduct": {
			input{
				request.UpdateProduct{
					ID:          1,
					Name:        "Test Product",
					Description: "This is a test product",
					CategoryID:  1,
					Price:       1000,
					Image:       "test_image.jpg",
				},
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().UpdateProduct(gomock.Any(), gomock.Any()).Return(nil)
			},
			output{http.StatusOK, "Successfully product updated", nil},
		},
		"Abnormal Case: UpdateProduct": {
			input{
				request.UpdateProduct{
					ID:          1,
					Name:        "Test Product",
					Description: "This is a test product",
					CategoryID:  1,
					Price:       1000,
					Image:       "test_image.jpg",
				},
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().UpdateProduct(gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to update product", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductUseCase)

			body, _ := json.Marshal(tt.input.updateProductRequest)
			req := httptest.NewRequest(echo.PUT, "/updateProduct", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)

			err := productHandler.UpdateProduct(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestProductHandler_SaveVariation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := usecaseMock.NewMockProductUseCase(ctrl)

	productHandler := handler.NewProductHandler(mockProductUseCase)

	type input struct {
		categoryIDStr string
		variation     request.Variation
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockProductUseCase)
		want          output
	}{
		"Normal Case: SaveVariation": {
			input{
				categoryIDStr: "1",
				variation: request.Variation{
					Names: []string{"Size", "Color"},
				},
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().SaveVariation(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			output{http.StatusCreated, "Successfully added variations", nil},
		},
		"Abnormal Case: SaveVariation": {
			input{
				categoryIDStr: "1",
				variation: request.Variation{
					Names: []string{"Size", "Color"},
				},
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().SaveVariation(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to add variation", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductUseCase)

			body, _ := json.Marshal(tt.input.variation)
			req := httptest.NewRequest(echo.POST, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/saveVariation/:category_id")
			ctx.SetParamNames("category_id")
			ctx.SetParamValues(tt.input.categoryIDStr)

			err := productHandler.SaveVariation(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestProductHandler_SaveVariationOption(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := usecaseMock.NewMockProductUseCase(ctrl)

	productHandler := handler.NewProductHandler(mockProductUseCase)

	type input struct {
		variationIDStr  string
		variationOption request.VariationOption
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockProductUseCase)
		want          output
	}{
		"Normal Case: SaveVariationOption": {
			input{
				variationIDStr: "1",
				variationOption: request.VariationOption{
					Values: []string{"S", "M", "L"},
				},
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().SaveVariationOption(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			output{http.StatusCreated, "Successfully added variation options", nil},
		},
		"Abnormal Case: SaveVariationOption": {
			input{
				variationIDStr: "1",
				variationOption: request.VariationOption{
					Values: []string{"S", "M", "L"},
				},
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().SaveVariationOption(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to add variation options", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductUseCase)

			body, _ := json.Marshal(tt.input.variationOption)
			req := httptest.NewRequest(echo.POST, "/", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/saveVariation/:variation_id")
			ctx.SetParamNames("variation_id")
			ctx.SetParamValues(tt.input.variationIDStr)

			err := productHandler.SaveVariationOption(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestProductHandler_GetAllVariations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := usecaseMock.NewMockProductUseCase(ctrl)

	productHandler := handler.NewProductHandler(mockProductUseCase)

	type input struct {
		categoryIDStr string
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockProductUseCase)
		want          output
	}{
		"Normal Case: GetAllVariations": {
			input{
				categoryIDStr: "1",
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().FindAllVariationsAndItsValues(gomock.Any(), gomock.Any()).Return([]response.Variation{
					{
						ID:   1,
						Name: "Test Variation",
					},
				}, nil)
			},
			output{http.StatusOK, "Successfully retrieved all variations and its values", nil},
		},
		"Abnormal Case: GetAllVariations": {
			input{
				categoryIDStr: "1",
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().FindAllVariationsAndItsValues(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to Get variations and its values", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductUseCase)

			req := httptest.NewRequest(echo.GET, "/getAllVariations/"+tt.input.categoryIDStr, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/getAllVariations/:category_id")
			ctx.SetParamNames("category_id")
			ctx.SetParamValues(tt.input.categoryIDStr)

			err := productHandler.GetAllVariations(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestProductHandler_SaveProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := usecaseMock.NewMockProductUseCase(ctrl)

	productHandler := handler.NewProductHandler(mockProductUseCase)
	imageFileHeader, _ := createMultipartFileHeader("test_image.jpg", "file content")

	type input struct {
		productRequest request.Product
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockProductUseCase)
		want          output
	}{
		"Normal Case: SaveProduct": {
			input{
				request.Product{
					Name:            "Test Product",
					Description:     "This is a test product",
					CategoryID:      1,
					Price:           1000,
					BrandID:         1,
					ImageFileHeader: imageFileHeader,
				},
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().SaveProduct(gomock.Any(), gomock.Any()).Return(nil)
			},
			output{http.StatusCreated, "Successfully product added", nil},
		},
		"Abnormal Case: SaveProduct": {
			input{
				request.Product{
					Name:            "Test Product",
					Description:     "This is a test product",
					CategoryID:      1,
					Price:           1000,
					BrandID:         1,
					ImageFileHeader: imageFileHeader,
				},
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().SaveProduct(gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to add product", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductUseCase)

			body, _ := json.Marshal(tt.input.productRequest)
			req := httptest.NewRequest(echo.POST, "/saveProduct", bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)

			err := productHandler.SaveProduct(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestProductHandler_GetAllProductsAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := usecaseMock.NewMockProductUseCase(ctrl)

	productHandler := handler.NewProductHandler(mockProductUseCase)

	type input struct {
		categoryIDStr string
		brandIDStr    string
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockProductUseCase)
		want          output
	}{
		"Normal Case: GetAllProductsAdmin": {
			input{
				categoryIDStr: "1",
				brandIDStr:    "1",
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().FindAllProducts(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]response.Product{
					{
						ID:          1,
						Name:        "Test Product",
						Description: "This is a test product",
						CategoryID:  1,
						Price:       1000,
						Image:       "test_image.jpg",
					},
				}, nil)
			},
			output{http.StatusOK, "Successfully found all products", nil},
		},
		"Abnormal Case: GetAllProductsAdmin": {
			input{
				categoryIDStr: "1",
				brandIDStr:    "1",
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().FindAllProducts(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to Get all products", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductUseCase)

			req := httptest.NewRequest(echo.GET, "/getAllProductsAdmin?category_id="+tt.input.categoryIDStr+"&brand_id="+tt.input.brandIDStr, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)

			err := productHandler.GetAllProductsAdmin()(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestProductHandler_GetAllProductsUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := usecaseMock.NewMockProductUseCase(ctrl)

	productHandler := handler.NewProductHandler(mockProductUseCase)

	type input struct {
		categoryIDStr string
		brandIDStr    string
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockProductUseCase)
		want          output
	}{
		"Normal Case: GetAllProductsUser": {
			input{
				categoryIDStr: "1",
				brandIDStr:    "1",
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().FindAllProducts(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return([]response.Product{
					{
						ID:          1,
						Name:        "Test Product",
						Description: "This is a test product",
						CategoryID:  1,
						Price:       1000,
						Image:       "test_image.jpg",
					},
				}, nil)
			},
			output{http.StatusOK, "Successfully found all products", nil},
		},
		"Abnormal Case: GetAllProductsUser": {
			input{
				categoryIDStr: "1",
				brandIDStr:    "1",
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().FindAllProducts(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to Get all products", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductUseCase)

			req := httptest.NewRequest(echo.GET, "/getAllProductsUser?category_id="+tt.input.categoryIDStr+"&brand_id="+tt.input.brandIDStr, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)

			err := productHandler.GetAllProductsUser()(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestProductHandler_GetProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := usecaseMock.NewMockProductUseCase(ctrl)

	productHandler := handler.NewProductHandler(mockProductUseCase)

	type input struct {
		productIDStr string
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockProductUseCase)
		want          output
	}{
		"Normal Case: GetProduct": {
			input{
				productIDStr: "1",
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().GetProduct(gomock.Any(), gomock.Any()).Return(response.Product{
					ID:          1,
					Name:        "Test Product",
					Description: "This is a test product",
					CategoryID:  1,
					Price:       1000,
					Image:       "test_image.jpg",
				}, nil)
			},
			output{http.StatusOK, "Successfully found all products", nil},
		},
		"Abnormal Case: GetProduct": {
			input{
				productIDStr: "1",
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().GetProduct(gomock.Any(), gomock.Any()).Return(response.Product{}, errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to Get all products", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductUseCase)

			req := httptest.NewRequest(echo.GET, "/getProduct?product_id="+tt.input.productIDStr, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/getProduct/:product_id")
			ctx.SetParamNames("product_id")
			ctx.SetParamValues(tt.input.productIDStr)

			err := productHandler.GetProduct(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestProductHandler_SaveProductItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := usecaseMock.NewMockProductUseCase(ctrl)

	productHandler := handler.NewProductHandler(mockProductUseCase)

	type input struct {
		productIDStr string
		productItem  request.ProductItem
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockProductUseCase)
		want          output
	}{
		"Normal Case: SaveProductItem": {
			input{
				productIDStr: "1",
				productItem: request.ProductItem{
					Name:               "Test Product Item",
					Price:              1000,
					VariationOptionIDs: []uint{1, 2},
					QtyInStock:         10,
					ImageFileHeaders:   []*multipart.FileHeader{},
				},
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().SaveProductItem(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			output{http.StatusCreated, "Successfully product item added", nil},
		},
		"Abnormal Case: SaveProductItem": {
			input{
				productIDStr: "1",
				productItem: request.ProductItem{
					Name:               "Test Product Item",
					Price:              1000,
					VariationOptionIDs: []uint{1, 2},
					QtyInStock:         10,
					ImageFileHeaders:   []*multipart.FileHeader{},
				},
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().SaveProductItem(gomock.Any(), gomock.Any(), gomock.Any()).Return(errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to add product item", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductUseCase)

			body, _ := json.Marshal(tt.input.productItem)
			req := httptest.NewRequest(echo.POST, "/saveProductItem?product_id="+tt.input.productIDStr, bytes.NewBuffer(body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/saveVariation/:product_id")
			ctx.SetParamNames("product_id")
			ctx.SetParamValues(tt.input.productIDStr)

			err := productHandler.SaveProductItem(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestProductHandler_GetAllProductItemsAdmin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := usecaseMock.NewMockProductUseCase(ctrl)

	productHandler := handler.NewProductHandler(mockProductUseCase)

	type input struct {
		productIDStr string
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockProductUseCase)
		want          output
	}{
		"Normal Case: GetAllProductItemsAdmin": {
			input{
				productIDStr: "1",
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().FindAllProductItems(gomock.Any(), gomock.Any()).Return([]response.ProductItems{
					{
						ID:               1,
						Name:             "Test Product Item",
						ProductID:        1,
						ItemName:         "Test Item",
						Price:            1000,
						DiscountPrice:    900,
						SKU:              "SKU123",
						QtyInStock:       10,
						CategoryName:     "Test Category",
						MainCategoryName: "Test Main Category",
						BrandID:          1,
						BrandName:        "Test Brand",
						VariationValues: []response.ProductVariationValue{
							{
								VariationID:       1,
								Name:              "Color",
								VariationOptionID: 1,
								Value:             "Red",
							},
						},
						Images: []string{
							"image1.jpg",
							"image2.jpg",
						},
					},
				}, nil)
			},
			output{http.StatusOK, "Successfully get all product items", nil},
		},
		"Abnormal Case: GetAllProductItemsAdmin": {
			input{
				productIDStr: "1",
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().FindAllProductItems(gomock.Any(), gomock.Any()).Return([]response.ProductItems{}, errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to get all product items", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductUseCase)

			req := httptest.NewRequest(echo.GET, "/getAllProductItemsAdmin/"+tt.input.productIDStr, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/saveVariation/:product_id")
			ctx.SetParamNames("product_id")
			ctx.SetParamValues(tt.input.productIDStr)

			err := productHandler.GetAllProductItemsAdmin()(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}

func TestProductHandler_GetAllProductItemsUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductUseCase := usecaseMock.NewMockProductUseCase(ctrl)

	productHandler := handler.NewProductHandler(mockProductUseCase)

	type input struct {
		productIDStr string
	}

	type output struct {
		wantCode int
		wantStr  string
		wantErr  error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(m *usecaseMock.MockProductUseCase)
		want          output
	}{
		"Normal Case: GetAllProductItemsUser": {
			input{
				productIDStr: "1",
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().FindAllProductItems(gomock.Any(), gomock.Any()).Return([]response.ProductItems{
					{
						ID:          1,
						Name:        "Test Product Item",
						ProductID:   1,
						ItemName:    "Test Item",
						Price:       1000,
						QtyInStock:  10,
						CategoryName: "Test Category",
						BrandID:     1,
						BrandName:   "Test Brand",
						VariationValues: []response.ProductVariationValue{
							{
								VariationID:       1,
								Name:              "Color",
								VariationOptionID: 1,
								Value:             "Red",
							},
						},
						Images: []string{
							"image1.jpg",
							"image2.jpg",
						},
					},
				}, nil)
			},
			output{http.StatusOK, "Successfully get all product items", nil},
		},
		"Abnormal Case: GetAllProductItemsUser": {
			input{
				productIDStr: "1",
			},
			func(m *usecaseMock.MockProductUseCase) {
				m.EXPECT().FindAllProductItems(gomock.Any(), gomock.Any()).Return([]response.ProductItems{}, errors.New("error"))
			},
			output{http.StatusInternalServerError, "Failed to get all product items", nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductUseCase)

			req := httptest.NewRequest(echo.GET, "/getAllProductItemsUser/"+tt.input.productIDStr, nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e := echo.New()
			e.Validator = &CustomValidator{validator: validator.New(validator.WithRequiredStructEnabled())}
			ctx := e.NewContext(req, rec)
			ctx.SetPath("/saveVariation/:product_id")
			ctx.SetParamNames("product_id")
			ctx.SetParamValues(tt.input.productIDStr)

			err := productHandler.GetAllProductItemsUser()(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.want.wantCode, rec.Code)
			assert.Contains(t, rec.Body.String(), tt.want.wantStr)
		})
	}
}
