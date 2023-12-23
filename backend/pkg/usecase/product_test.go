package usecase_test

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/domain"
	repoInterfaces "github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	productMock "github.com/shion0625/FYP/backend/pkg/repository/mock"
	cloudMock "github.com/shion0625/FYP/backend/pkg/service/mock/cloud"
	"github.com/shion0625/FYP/backend/pkg/usecase"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestProductUseCase_FindAllCategories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := productMock.NewMockProductRepository(ctrl)

	product := usecase.NewProductUseCase(mockProductRepo, nil)

	ctx := echo.New().NewContext(nil, nil)
	pagination := request.Pagination{PageNumber: 1, Count: 10}

	type input struct {
		pagination request.Pagination
	}

	type output struct {
		wantCategories []response.Category
		wantErr        error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *productMock.MockProductRepository)
		want          output
	}{
		"Normal Case: FindAllCategories": {
			input{pagination},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().FindAllMainCategories(gomock.Any(), pagination).Return([]response.Category{}, nil)
			},
			output{[]response.Category{}, nil},
		},
		"Abnormal Case: FindAllCategories": {
			input{pagination},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().FindAllMainCategories(gomock.Any(), pagination).Return(nil, errors.New("error"))
			},
			output{nil, fmt.Errorf("failed find all main categories: %w", errors.New("error"))},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductRepo)
			categories, err := product.FindAllCategories(ctx, tt.input.pagination)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantCategories, categories)
			}
		})
	}
}

func TestProductUseCase_SaveCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := productMock.NewMockProductRepository(ctrl)

	product := usecase.NewProductUseCase(mockProductRepo, nil)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	categoryName := "Test Category"

	type input struct {
		categoryName string
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *productMock.MockProductRepository)
		want          output
	}{
		"Normal Case: SaveCategory": {
			input{categoryName},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().IsCategoryNameExist(gomock.Any(), categoryName).Return(false, nil).Times(1)
				mr.EXPECT().SaveCategory(gomock.Any(), categoryName).Return(nil).Times(1)
			},
			output{nil},
		},
		"Abnormal Case: SaveCategory": {
			input{categoryName},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().IsCategoryNameExist(gomock.Any(), categoryName).Return(true, nil).Times(1)
				mr.EXPECT().SaveCategory(gomock.Any(), categoryName).Return(nil).Times(0)
			},
			output{fmt.Errorf("category already exist")},
		},
		"Abnormal Case: SaveCategory1": {
			input{categoryName},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().IsCategoryNameExist(gomock.Any(), categoryName).Return(false, errors.New("error")).Times(1)
				mr.EXPECT().SaveCategory(gomock.Any(), categoryName).Return(nil).Times(0)
			},
			output{fmt.Errorf("failed to check category already exist: error")},
		},
		"Abnormal Case: SaveCategory2": {
			input{categoryName},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().IsCategoryNameExist(gomock.Any(), categoryName).Return(false, nil).Times(1)
				mr.EXPECT().SaveCategory(gomock.Any(), categoryName).Return(errors.New("error")).Times(1)
			},
			output{fmt.Errorf("failed to save category: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductRepo)
			err := product.SaveCategory(ctx, tt.input.categoryName)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}

func TestProductUseCase_SaveVariation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := productMock.NewMockProductRepository(ctrl)

	product := usecase.NewProductUseCase(mockProductRepo, nil)

	ctx := echo.New().NewContext(nil, nil)
	categoryID := uint(1)
	variationNames := []string{"Variation 1", "Variation 2"}

	type input struct {
		categoryID     uint
		variationNames []string
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *productMock.MockProductRepository)
		want          output
	}{
		"Normal Case: SaveVariation": {
			input{categoryID, variationNames},
			func(mr *productMock.MockProductRepository) {
				for _, variationName := range variationNames {
					mr.EXPECT().IsVariationNameExistForCategory(gomock.Any(), variationName, categoryID).Return(false, nil)
					mr.EXPECT().SaveVariation(gomock.Any(), categoryID, variationName).Return(nil)
				}
				mr.EXPECT().Transactions(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx echo.Context, f func(repo repoInterfaces.ProductRepository) error) error {
						return f(mr)
					},
				)
			},
			output{nil},
		},
		"Abnormal Case: SaveVariation1": {
			input{categoryID, variationNames},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().IsVariationNameExistForCategory(gomock.Any(), variationNames[0], categoryID).Return(false, errors.New("error"))
				mr.EXPECT().Transactions(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx echo.Context, f func(repo repoInterfaces.ProductRepository) error) error {
						return f(mr)
					},
				)
			},
			output{fmt.Errorf("failed to save variation: failed to check variation already exist: error")},
		},
		"Abnormal Case: SaveVariation": {
			input{categoryID, variationNames},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().IsVariationNameExistForCategory(gomock.Any(), variationNames[0], categoryID).Return(true, nil)
				mr.EXPECT().Transactions(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx echo.Context, f func(repo repoInterfaces.ProductRepository) error) error {
						return f(mr)
					},
				)
			},
			output{fmt.Errorf("failed to save variation: variation name %s: %w", variationNames[0], usecase.ErrVariationAlreadyExist)},
		},
		"Abnormal Case: SaveVariation2": {
			input{categoryID, variationNames},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().IsVariationNameExistForCategory(gomock.Any(), variationNames[0], categoryID).Return(false, nil)
				mr.EXPECT().SaveVariation(gomock.Any(), categoryID, variationNames[0]).Return(errors.New("error"))
				mr.EXPECT().Transactions(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx echo.Context, f func(repo repoInterfaces.ProductRepository) error) error {
						return f(mr)
					},
				)
			},
			output{fmt.Errorf("failed to save variation: failed to save variation: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductRepo)
			err := product.SaveVariation(ctx, tt.input.categoryID, tt.input.variationNames)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}

func TestProductUseCase_SaveVariationOption(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := productMock.NewMockProductRepository(ctrl)

	product := usecase.NewProductUseCase(mockProductRepo, nil)

	ctx := echo.New().NewContext(nil, nil)
	variationID := uint(1)
	variationOptionValues := []string{"Option 1", "Option 2"}

	type input struct {
		variationID           uint
		variationOptionValues []string
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *productMock.MockProductRepository)
		want          output
	}{
		"Normal Case: SaveVariationOption": {
			input{variationID, variationOptionValues},
			func(mr *productMock.MockProductRepository) {
				for _, variationValue := range variationOptionValues {
					mr.EXPECT().IsVariationValueExistForVariation(gomock.Any(), variationValue, variationID).Return(false, nil)
					mr.EXPECT().SaveVariationOption(gomock.Any(), variationID, variationValue).Return(nil)
				}
				mr.EXPECT().Transactions(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx echo.Context, f func(repo repoInterfaces.ProductRepository) error) error {
						return f(mr)
					},
				)
			},
			output{nil},
		},
		"Abnormal Case: SaveVariationOption1": {
			input{variationID, variationOptionValues},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().IsVariationValueExistForVariation(gomock.Any(), variationOptionValues[0], variationID).Return(true, errors.New("error"))
				mr.EXPECT().Transactions(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx echo.Context, f func(repo repoInterfaces.ProductRepository) error) error {
						return f(mr)
					},
				)
			},
			output{fmt.Errorf("failed to save variation option: failed to check variation already exist: error")},
		},
		"Abnormal Case: SaveVariationOption": {
			input{variationID, variationOptionValues},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().IsVariationValueExistForVariation(gomock.Any(), variationOptionValues[0], variationID).Return(true, nil)
				mr.EXPECT().Transactions(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx echo.Context, f func(repo repoInterfaces.ProductRepository) error) error {
						return f(mr)
					},
				)
			},
			output{fmt.Errorf("failed to save variation option: variation option value %s: %w", variationOptionValues[0], usecase.ErrVariationOptionAlreadyExist)},
		},
		"Abnormal Case: SaveVariationOption2": {
			input{variationID, variationOptionValues},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().IsVariationValueExistForVariation(gomock.Any(), variationOptionValues[0], variationID).Return(false, nil)
				mr.EXPECT().SaveVariationOption(gomock.Any(), variationID, variationOptionValues[0]).Return(errors.New("error"))

				mr.EXPECT().Transactions(gomock.Any(), gomock.Any()).DoAndReturn(
					func(ctx echo.Context, f func(repo repoInterfaces.ProductRepository) error) error {
						return f(mr)
					},
				)
			},
			output{fmt.Errorf("failed to save variation option: failed to save variation option: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductRepo)
			err := product.SaveVariationOption(ctx, tt.input.variationID, tt.input.variationOptionValues)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}

func TestProductUseCase_FindAllVariationsAndItsValues(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := productMock.NewMockProductRepository(ctrl)

	product := usecase.NewProductUseCase(mockProductRepo, nil)

	ctx := echo.New().NewContext(nil, nil)
	categoryID := uint(1)

	variations := []response.Variation{
		{ID: 1, Name: "Variation 1"},
		{ID: 2, Name: "Variation 2"},
	}
	type input struct {
		categoryID uint
	}

	type output struct {
		wantVariations []response.Variation
		wantErr        error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *productMock.MockProductRepository)
		want          output
	}{
		"Normal Case: FindAllVariationsAndItsValues": {
			input{categoryID},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().FindAllVariationsByCategoryID(gomock.Any(), categoryID).Return([]response.Variation{}, nil)
				mr.EXPECT().FindAllVariationOptionsByVariationID(gomock.Any(), gomock.Any()).Return([]response.VariationOption{}, nil).Times(0)
			},
			output{[]response.Variation{}, nil},
		},
		"Normal Case: FindAllVariationsAndItsValues1": {
			input{categoryID},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().FindAllVariationsByCategoryID(gomock.Any(), categoryID).Return(variations, nil)
				for _, variation := range variations {
					mr.EXPECT().FindAllVariationOptionsByVariationID(gomock.Any(), variation.ID).Return([]response.VariationOption{}, nil)
				}
			},
			output{[]response.Variation{
				{ID: 1, Name: "Variation 1", VariationOptions: []response.VariationOption{}},
				{ID: 2, Name: "Variation 2", VariationOptions: []response.VariationOption{}},
			}, nil},
		},
		"Abnormal Case: FindAllVariationsAndItsValues": {
			input{categoryID},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().FindAllVariationsByCategoryID(gomock.Any(), categoryID).Return(nil, errors.New("error"))
			},
			output{nil, fmt.Errorf("failed to find all variations of category: error")},
		},
		"Abnormal Case: FindAllVariationsAndItsValues1": {
			input{categoryID},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().FindAllVariationsByCategoryID(gomock.Any(), categoryID).Return(variations, nil)
				mr.EXPECT().FindAllVariationOptionsByVariationID(gomock.Any(), variations[0].ID).Return(nil, errors.New("error")).Times(1)
			},
			output{nil, fmt.Errorf("failed to get variation option: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductRepo)
			variations, err := product.FindAllVariationsAndItsValues(ctx, tt.input.categoryID)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantVariations, variations)
			}
		})
	}
}

func TestProductUseCase_FindAllProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := productMock.NewMockProductRepository(ctrl)
	mockCloudService := cloudMock.NewMockCloudService(ctrl)

	product := usecase.NewProductUseCase(mockProductRepo, mockCloudService)

	ctx := echo.New().NewContext(nil, nil)
	pagination := request.Pagination{PageNumber: 1, Count: 10}
	categoryID := uint(1)
	brandID := uint(1)

	products := []response.Product{
		{ID: 1, Name: "Product 1", Image: "image1.jpg"},
		{ID: 2, Name: "Product 2", Image: "image2.jpg"},
	}

	type input struct {
		pagination request.Pagination
		categoryID *uint
		brandID    *uint
	}

	type output struct {
		wantProducts []response.Product
		wantErr      error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService)
		want          output
	}{
		"Normal Case: FindAllProducts": {
			input{pagination, &categoryID, &brandID},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().FindAllProducts(gomock.Any(), pagination, &categoryID, &brandID).Return(products, nil)
				cs.EXPECT().GetFileUrl(gomock.Any(), gomock.Any()).Return("tom.jpg", nil).Times(2)
			},
			output{[]response.Product{
				{ID: 1, Name: "Product 1", Image: "tom.jpg"},
				{ID: 2, Name: "Product 2", Image: "tom.jpg"},
			}, nil},
		},
		"Normal Case: FindAllProducts1": {
			input{pagination, &categoryID, &brandID},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().FindAllProducts(gomock.Any(), pagination, &categoryID, &brandID).Return(products, nil)
				cs.EXPECT().GetFileUrl(gomock.Any(), gomock.Any()).Return("", errors.New("error")).Times(2)
			},
			output{products, nil},
		},
		"Abnormal Case: FindAllProducts": {
			input{pagination, &categoryID, &brandID},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().FindAllProducts(gomock.Any(), pagination, &categoryID, &brandID).Return(nil, errors.New("error"))
			},
			output{nil, fmt.Errorf("failed to get product details from database: %w", errors.New("error"))},
		},
	}
	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductRepo, mockCloudService)
			products, err := product.FindAllProducts(ctx, tt.input.pagination, tt.input.categoryID, tt.input.brandID)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantProducts, products)
			}
		})
	}
}

func TestProductUseCase_GetProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := productMock.NewMockProductRepository(ctrl)
	mockCloudService := cloudMock.NewMockCloudService(ctrl)

	product := usecase.NewProductUseCase(mockProductRepo, mockCloudService)

	ctx := echo.New().NewContext(nil, nil)
	productID := uint(1)

	type input struct {
		productID uint
	}

	type output struct {
		wantProduct response.Product
		wantErr     error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService)
		want          output
	}{
		"Normal Case: GetProduct": {
			input{productID},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().FindProductByID(gomock.Any(), productID).Return(response.Product{}, nil)
				cs.EXPECT().GetFileUrl(gomock.Any(), gomock.Any()).Return("", nil)
			},
			output{response.Product{}, nil},
		},
		"Abnormal Case: GetProduct": {
			input{productID},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().FindProductByID(gomock.Any(), productID).Return(response.Product{}, errors.New("error"))
			},
			output{response.Product{}, fmt.Errorf("failed to get product from database: %w", errors.New("error"))},
		},
		"Abnormal Case: GetProduct1": {
			input{productID},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().FindProductByID(gomock.Any(), productID).Return(response.Product{}, nil)
				cs.EXPECT().GetFileUrl(gomock.Any(), gomock.Any()).Return("", errors.New("error"))
			},
			output{response.Product{}, fmt.Errorf("failed to get image url from could service: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductRepo, mockCloudService)
			productRes, err := product.GetProduct(ctx, tt.input.productID)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantProduct, productRes)
			}
		})
	}
}

func TestProductUseCase_SaveProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := productMock.NewMockProductRepository(ctrl)
	mockCloudService := cloudMock.NewMockCloudService(ctrl)

	product := usecase.NewProductUseCase(mockProductRepo, mockCloudService)

	ctx := echo.New().NewContext(nil, nil)
	productReq := request.Product{
		Name:            "Test Product",
		Description:     "Test Description",
		CategoryID:      1,
		BrandID:         1,
		Price:           100,
		ImageFileHeader: &multipart.FileHeader{},
	}

	type input struct {
		product request.Product
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService)
		want          output
	}{
		"Normal Case: SaveProduct": {
			input{productReq},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().IsProductNameExist(gomock.Any(), productReq.Name).Return(false, nil)
				cs.EXPECT().SaveFile(gomock.Any(), productReq.ImageFileHeader).Return("test", nil)
				mr.EXPECT().SaveProduct(gomock.Any(), gomock.Any()).Return(nil)
			},
			output{nil},
		},
		"Abnormal Case: SaveProduct1": {
			input{productReq},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().IsProductNameExist(gomock.Any(), productReq.Name).Return(false, errors.New("error"))
			},
			output{fmt.Errorf("failed to check product name already exist: error")},
		},
		"Abnormal Case: SaveProduct": {
			input{productReq},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().IsProductNameExist(gomock.Any(), productReq.Name).Return(true, nil)
			},
			output{fmt.Errorf("product name %s: %w", productReq.Name, usecase.ErrProductAlreadyExist)},
		},
		"Abnormal Case: SaveProduct2": {
			input{productReq},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().IsProductNameExist(gomock.Any(), productReq.Name).Return(false, nil).Times(1)
				cs.EXPECT().SaveFile(gomock.Any(), productReq.ImageFileHeader).Return("", errors.New("error")).Times(1)
			},
			output{fmt.Errorf("failed to save image on cloud storage: error")},
		},
		"Abnormal Case: SaveProduct3": {
			input{productReq},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().IsProductNameExist(gomock.Any(), productReq.Name).Return(false, nil).Times(1)
				cs.EXPECT().SaveFile(gomock.Any(), productReq.ImageFileHeader).Return("", nil).Times(1)
				mr.EXPECT().SaveProduct(gomock.Any(), gomock.Any()).Return(errors.New("error")).Times(1)
			},
			output{fmt.Errorf("failed to save product: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductRepo, mockCloudService)
			err := product.SaveProduct(ctx, tt.input.product)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}

func TestProductUseCase_SaveProductItem(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := productMock.NewMockProductRepository(ctrl)
	mockCloudService := cloudMock.NewMockCloudService(ctrl)

	product := usecase.NewProductUseCase(mockProductRepo, mockCloudService)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	productID := uint(1)
	productItem := request.ProductItem{
		QtyInStock:         10,
		Price:              100,
		VariationOptionIDs: []uint{1, 2},
		ImageFileHeaders:   []*multipart.FileHeader{{}},
	}

	type input struct {
		productID   uint
		productItem request.ProductItem
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService)
		want          output
	}{
		"Normal Case: SaveProductItem": {
			input{productID, productItem},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().FindVariationCountForProduct(gomock.Any(), productID).Return(uint(len(productItem.VariationOptionIDs)), nil)
				mr.EXPECT().FindAllProductItemIDsByProductIDAndVariationOptionID(gomock.Any(), productID, gomock.Any()).Return([]uint{}, nil).AnyTimes()
				mr.EXPECT().Transactions(gomock.Any(), gomock.Any()).DoAndReturn(
					func(_ echo.Context, f func(mr repoInterfaces.ProductRepository) error) error {
						return f(mr)
					},
				).AnyTimes()
				cs.EXPECT().SaveFile(gomock.Any(), gomock.Any()).Return("test", nil).AnyTimes()
				mr.EXPECT().SaveProductItem(gomock.Any(), gomock.Any()).Return(uint(1), nil).AnyTimes()
				mr.EXPECT().SaveProductConfiguration(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
				mr.EXPECT().SaveProductItemImage(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			},
			output{nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductRepo, mockCloudService)
			err := product.SaveProductItem(ctx, tt.input.productID, tt.input.productItem)
			fmt.Print(err)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}

func TestProductUseCase_SaveProductItemAbnormal(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := productMock.NewMockProductRepository(ctrl)
	mockCloudService := cloudMock.NewMockCloudService(ctrl)

	product := usecase.NewProductUseCase(mockProductRepo, mockCloudService)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)

	productID := uint(1)
	productItem := request.ProductItem{
		QtyInStock:         10,
		Price:              100,
		VariationOptionIDs: []uint{1, 2},
		ImageFileHeaders:   []*multipart.FileHeader{{}},
	}

	type input struct {
		productID   uint
		productItem request.ProductItem
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService)
		want          output
	}{
		"Abnormal Case: SaveProductItem": {
			input{productID, productItem},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().FindVariationCountForProduct(gomock.Any(), productID).Return(uint(len(productItem.VariationOptionIDs)+1), nil)
			},
			output{usecase.ErrNotEnoughVariations},
		},
		"Abnormal Case: SaveProductItem2": {
			input{productID, productItem},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().FindVariationCountForProduct(gomock.Any(), productID).Return(uint(len(productItem.VariationOptionIDs)), nil).Times(1)
				mr.EXPECT().FindAllProductItemIDsByProductIDAndVariationOptionID(gomock.Any(), productID, gomock.Any()).Return([]uint{}, errors.New("error")).Times(1)
			},
			output{fmt.Errorf("failed to find product item ids from database using product id and variation option id: error")},
		},
		"Abnormal Case: SaveProductItem1": {
			input{productID, productItem},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().FindVariationCountForProduct(gomock.Any(), productID).Return(uint(0), errors.New("error"))
			},
			output{fmt.Errorf("failed to get variation count of product from database: error")},
		},

		"Abnormal Case: SaveProductItem3": {
			input{productID, productItem},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().FindVariationCountForProduct(gomock.Any(), productID).Return(uint(len(productItem.VariationOptionIDs)), nil).Times(1)
				mr.EXPECT().FindAllProductItemIDsByProductIDAndVariationOptionID(gomock.Any(), productID, gomock.Any()).Return([]uint{1, 1}, nil).Times(1)
			},
			output{usecase.ErrProductItemAlreadyExist},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductRepo, mockCloudService)
			err := product.SaveProductItem(ctx, tt.input.productID, tt.input.productItem)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}

func TestProductUseCase_FindAllProductItems(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := productMock.NewMockProductRepository(ctrl)
	mockCloudService := cloudMock.NewMockCloudService(ctrl)

	product := usecase.NewProductUseCase(mockProductRepo, mockCloudService)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	productID := uint(1)

	productItemsDB := []response.ProductItemsDB{
		{
			ID:               1,
			Name:             "Test Name",
			ItemName:         "Test ItemName",
			Price:            100,
			DiscountPrice:    80,
			SKU:              "Test SKU",
			QtyInStock:       10,
			CategoryName:     "Test CategoryName",
			MainCategoryName: "Test MainCategoryName",
			BrandID:          1,
			BrandName:        "Test BrandName",
			Images:           []string{"image1", "image2"},
		},
	}

	type input struct {
		productID uint
	}

	type output struct {
		wantProductItems []response.ProductItems
		wantErr          error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService)
		want          output
	}{
		"Normal Case: FindAllProductItems with complete data": {
			input{productID},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().FindAllProductItems(gomock.Any(), productID).Return(productItemsDB, nil).Times(1)
				mr.EXPECT().FindAllVariationValuesOfProductItem(gomock.Any(), gomock.Any()).Return([]response.ProductVariationValue{}, nil).Times(1)
				mr.EXPECT().FindAllProductItemImages(gomock.Any(), gomock.Any()).Return([]string{"image1dec", "image2dec"}, nil).AnyTimes()
				cs.EXPECT().GetFileUrl(gomock.Any(), gomock.Any()).Return("changeImage", nil).AnyTimes()
			},
			output{
				wantProductItems: []response.ProductItems{
					{
						ID:               1,
						Name:             "Test Name",
						ItemName:         "Test ItemName",
						Price:            100,
						DiscountPrice:    80,
						SKU:              "Test SKU",
						QtyInStock:       10,
						CategoryName:     "Test CategoryName",
						MainCategoryName: "Test MainCategoryName",
						BrandID:          1,
						BrandName:        "Test BrandName",
						Images:           []string{"changeImage", "changeImage"},
						VariationValues:  []response.ProductVariationValue{},
					},
				},
				wantErr: nil,
			},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductRepo, mockCloudService)
			productItems, err := product.FindAllProductItems(ctx, tt.input.productID)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantProductItems, productItems)
			}
		})
	}
}

// func TestProductUseCase_FindAllProductItemsAbnormal(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockProductRepo := productMock.NewMockProductRepository(ctrl)
// 	mockCloudService := cloudMock.NewMockCloudService(ctrl)

// 	product := usecase.NewProductUseCase(mockProductRepo, mockCloudService)

// 	req := httptest.NewRequest(echo.GET, "/", nil)
// 	rec := httptest.NewRecorder()
// 	ctx := echo.New().NewContext(req, rec)
// 	productID := uint(1)

// 	productItemsDB := []response.ProductItemsDB{
// 		{
// 			ID:               1,
// 			Name:             "Test Name",
// 			ItemName:         "Test ItemName",
// 			Price:            100,
// 			DiscountPrice:    80,
// 			SKU:              "Test SKU",
// 			QtyInStock:       10,
// 			CategoryName:     "Test CategoryName",
// 			MainCategoryName: "Test MainCategoryName",
// 			BrandID:          1,
// 			BrandName:        "Test BrandName",
// 			Images:           []string{"image1", "image2"},
// 		},
// 	}
// 	type input struct {
// 		productID uint
// 	}

// 	type output struct {
// 		wantProductItems []response.ProductItems
// 		wantErr          error
// 	}

// 	tests := map[string]struct {
// 		input         input
// 		prepareMockFn func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService)
// 		want          output
// 	}{
// 		"Abnormal Case: FindAllProductItems": {
// 			input{productID},
// 			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
// 				mr.EXPECT().FindAllProductItems(gomock.Any(), productID).Return(nil, errors.New("error")).Times(1)
// 			},
// 			output{nil, fmt.Errorf("failed to find all product items: %w", errors.New("error"))},
// 		},
// 		"Abnormal Case: FindAllProductItems1": {
// 			input{productID},
// 			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
// 				mr.EXPECT().FindAllProductItems(gomock.Any(), productID).Return(productItemsDB, nil).Times(1)
// 				mr.EXPECT().FindAllVariationValuesOfProductItem(gomock.Any(), gomock.Any()).Return(nil, errors.New("error")).Times(1)
// 			},
// 			output{nil, fmt.Errorf("failed to find variation values product item: error")},
// 		},
// 		"Abnormal Case: FindAllProductItems2": {
// 			input{productID},
// 			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
// 				mr.EXPECT().FindAllProductItems(gomock.Any(), productID).Return(productItemsDB, nil).Times(1)
// 				mr.EXPECT().FindAllVariationValuesOfProductItem(gomock.Any(), gomock.Any()).Return([]response.ProductVariationValue{}, nil).AnyTimes()
// 				mr.EXPECT().FindAllProductItemImages(gomock.Any(), gomock.Any()).Return([]string{}, errors.New("error")).Times(1)
// 			},
// 			output{nil, fmt.Errorf("failed to find images of product item: error")},
// 		},
// 	}

// 	for testName, tt := range tests {
// 		t.Run(testName, func(t *testing.T) {
// 			tt.prepareMockFn(mockProductRepo, mockCloudService)
// 			productItems, err := product.FindAllProductItems(ctx, tt.input.productID)
// 			if err != nil {
// 				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
// 			} else {
// 				assert.Equal(t, tt.want.wantProductItems, productItems)
// 			}
// 		})
// 	}
// }

func TestProductUseCase_FindAllProductItemsAbnormal3(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := productMock.NewMockProductRepository(ctrl)
	mockCloudService := cloudMock.NewMockCloudService(ctrl)

	product := usecase.NewProductUseCase(mockProductRepo, mockCloudService)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, rec)
	productID := uint(1)

	productItemsDB := []response.ProductItemsDB{
		{
			ID:               1,
			Name:             "Test Name",
			ItemName:         "Test ItemName",
			Price:            100,
			DiscountPrice:    80,
			SKU:              "Test SKU",
			QtyInStock:       10,
			CategoryName:     "Test CategoryName",
			MainCategoryName: "Test MainCategoryName",
			BrandID:          1,
			BrandName:        "Test BrandName",
			Images:           []string{"image1", "image2"},
		},
	}
	type input struct {
		productID uint
	}

	type output struct {
		wantProductItems []response.ProductItems
		wantErr          error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService)
		want          output
	}{
		"Abnormal Case: FindAllProductItems3": {
			input{productID},
			func(mr *productMock.MockProductRepository, cs *cloudMock.MockCloudService) {
				mr.EXPECT().FindAllProductItems(gomock.Any(), productID).Return(productItemsDB, nil).Times(1)
				mr.EXPECT().FindAllVariationValuesOfProductItem(gomock.Any(), gomock.Any()).Return([]response.ProductVariationValue{}, nil).AnyTimes()
				mr.EXPECT().FindAllProductItemImages(gomock.Any(), gomock.Any()).Return([]string{"image1dec", "image2dec"}, nil).Times(1)
				cs.EXPECT().GetFileUrl(gomock.Any(), gomock.Any()).Return("", errors.New("error")).AnyTimes()
			},
			output{nil, fmt.Errorf("failed to get image url from could service: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductRepo, mockCloudService)
			productItems, err := product.FindAllProductItems(ctx, tt.input.productID)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantProductItems, productItems)
			}
		})
	}
}

func TestProductUseCase_UpdateProduct(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepo := productMock.NewMockProductRepository(ctrl)

	product := usecase.NewProductUseCase(mockProductRepo, nil)

	ctx := echo.New().NewContext(nil, nil)
	updateDetails := domain.Product{
		ID:   1,
		Name: "Test Product",
	}

	type input struct {
		updateDetails domain.Product
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func(mr *productMock.MockProductRepository)
		want          output
	}{
		"Normal Case: UpdateProduct": {
			input{updateDetails},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().IsProductNameExistForOtherProduct(gomock.Any(), updateDetails.Name, updateDetails.ID).Return(false, nil)
				mr.EXPECT().UpdateProduct(gomock.Any(), updateDetails).Return(nil)
			},
			output{nil},
		},
		"Abnormal Case: UpdateProduct1": {
			input{updateDetails},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().IsProductNameExistForOtherProduct(gomock.Any(), updateDetails.Name, updateDetails.ID).Return(true, errors.New("error"))
			},
			output{fmt.Errorf("failed to check product name already exist for other product: error")},
		},
		"Abnormal Case: UpdateProduct": {
			input{updateDetails},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().IsProductNameExistForOtherProduct(gomock.Any(), updateDetails.Name, updateDetails.ID).Return(true, nil)
			},
			output{fmt.Errorf("product name %s: %w", updateDetails.Name, usecase.ErrProductAlreadyExist)},
		},
		"Abnormal Case: UpdateProduct2": {
			input{updateDetails},
			func(mr *productMock.MockProductRepository) {
				mr.EXPECT().IsProductNameExistForOtherProduct(gomock.Any(), updateDetails.Name, updateDetails.ID).Return(false, nil)
				mr.EXPECT().UpdateProduct(gomock.Any(), updateDetails).Return(errors.New("error"))
			},
			output{fmt.Errorf("failed to update product: error")},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn(mockProductRepo)
			err := product.UpdateProduct(ctx, tt.input.updateDetails)
			if err != nil {
				assert.Equal(t, tt.want.wantErr.Error(), err.Error())
			} else {
				assert.Equal(t, tt.want.wantErr, err)
			}
		})
	}
}
