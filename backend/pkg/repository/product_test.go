package repository_test

import (
	// "errors".
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/repository"
	"github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestProductRepository_Transactions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		trxFn         func(repo interfaces.ProductRepository) error
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: Transactions": {
			trxFn: func(repo interfaces.ProductRepository) error {
				return nil
			},
			prepareMockFn: func() {
				mock.ExpectBegin()
				mock.ExpectCommit()
			},
			wantErr: nil,
		},
		"Error Case: Transactions": {
			trxFn: func(repo interfaces.ProductRepository) error {
				return errors.New("transaction error")
			},
			prepareMockFn: func() {
				mock.ExpectBegin()
				mock.ExpectRollback()
			},
			wantErr: errors.New("transaction error"),
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			err := productRepo.Transactions(echo.New().AcquireContext(), tt.trxFn)

			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductRepository_IsCategoryNameExist(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)
	tests := map[string]struct {
		input         string
		prepareMockFn func()
		want          bool
	}{
		"Normal Case: IsCategoryNameExist": {
			input: "Test Category",
			prepareMockFn: func() {
				expectedSQL := `^SELECT \* FROM "categories" WHERE name = \$1 AND category_id IS NULL ORDER BY "categories"."id" LIMIT 1`
				categories := sqlmock.NewRows([]string{"name"}).AddRow("Test Category")
				mock.ExpectQuery(expectedSQL).WithArgs("Test Category").WillReturnRows(categories)
			},
			want: true,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			exist, err := productRepo.IsCategoryNameExist(echo.New().AcquireContext(), tt.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if exist != tt.want {
				t.Errorf("expected %v, but got %v", tt.want, exist)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductRepository_SaveCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		input         string
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: SaveCategory": {
			input: "Test Category",
			prepareMockFn: func() {
				addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
				expectedSQL := `^INSERT INTO "categories" \(.+\) VALUES \(.+\) RETURNING "id"$`
				mock.ExpectBegin()
				mock.ExpectQuery(expectedSQL).
					WithArgs(
						sqlmock.AnyArg()).
					WillReturnRows(addRow)
				mock.ExpectCommit()
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			err := productRepo.SaveCategory(echo.New().AcquireContext(), tt.input)

			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductRepository_FindAllMainCategories(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)
	tests := map[string]struct {
		input         request.Pagination
		prepareMockFn func()
		want          []response.Category
	}{
		"Normal Case: FindAllMainCategories": {
			input: request.Pagination{PageNumber: 30, Count: 10},
			prepareMockFn: func() {
				expectedSQL := `^SELECT \* FROM "categories" LIMIT (.+) OFFSET (.+)$`
				mock.ExpectQuery(expectedSQL).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Test Category"))
			},
			want: []response.Category{{Name: "Test Category"}},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			categories, err := productRepo.FindAllMainCategories(echo.New().AcquireContext(), tt.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(categories, tt.want) {
				t.Errorf("expected %v, but got %v", tt.want, categories)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductRepository_FindAllVariationsByCategoryID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		input         uint
		prepareMockFn func()
		want          []response.Variation
	}{
		"Normal Case: FindAllVariationsByCategoryID": {
			input: 1,
			prepareMockFn: func() {
				expectedSQL := `^SELECT \* FROM "variations" WHERE category_id = (.+)$`
				mock.ExpectQuery(expectedSQL).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Test Variation"))
			},
			want: []response.Variation{{Name: "Test Variation"}},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			variations, err := productRepo.FindAllVariationsByCategoryID(echo.New().AcquireContext(), tt.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(variations, tt.want) {
				t.Errorf("expected %v, but got %v", tt.want, variations)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductRepository_FindAllVariationOptionsByVariationID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		input         uint
		prepareMockFn func()
		want          []response.VariationOption
	}{
		"Normal Case: FindAllVariationOptionsByVariationID": {
			input: 1,
			prepareMockFn: func() {
				expectedSQL := `^SELECT \* FROM "variation_options" WHERE variation_id = (.+)$`
				mock.ExpectQuery(expectedSQL).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"value"}).AddRow("Test VariationOption"))
			},
			want: []response.VariationOption{{Value: "Test VariationOption"}},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			variationOptions, err := productRepo.FindAllVariationOptionsByVariationID(echo.New().AcquireContext(), tt.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(variationOptions, tt.want) {
				t.Errorf("expected %v, but got %v", tt.want, variationOptions)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductRepository_IsVariationNameExistForCategory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputName     string
		inputCategory uint
		prepareMockFn func()
		want          bool
	}{
		"Normal Case: IsVariationNameExistForCategory": {
			inputName:     "Test Variation",
			inputCategory: 1,
			prepareMockFn: func() {
				expectedSQL := `^SELECT \* FROM "variations" WHERE name = (.+) AND category_id = (.+)$`
				mock.ExpectQuery(expectedSQL).WithArgs("Test Variation", 1).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Test Variation"))
			},
			want: true,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			exist, err := productRepo.IsVariationNameExistForCategory(echo.New().AcquireContext(), tt.inputName, tt.inputCategory)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if exist != tt.want {
				t.Errorf("expected %v, but got %v", tt.want, exist)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductRepository_IsVariationValueExistForVariation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputValue     string
		inputVariation uint
		prepareMockFn  func()
		want           bool
	}{
		"Normal Case: IsVariationValueExistForVariation": {
			inputValue:     "Test Value",
			inputVariation: 1,
			prepareMockFn: func() {
				expectedSQL := `^SELECT \* FROM "variation_options" WHERE value = (.+) AND variation_id = (.+)$`
				mock.ExpectQuery(expectedSQL).WithArgs("Test Value", 1).WillReturnRows(sqlmock.NewRows([]string{"value"}).AddRow("Test Value"))
			},
			want: true,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			exist, err := productRepo.IsVariationValueExistForVariation(echo.New().AcquireContext(), tt.inputValue, tt.inputVariation)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if exist != tt.want {
				t.Errorf("expected %v, but got %v", tt.want, exist)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductRepository_SaveVariation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputCategory  uint
		inputVariation string
		prepareMockFn  func()
		wantErr        error
	}{
		"Normal Case: SaveVariation": {
			inputCategory:  1,
			inputVariation: "Test Variation",
			prepareMockFn: func() {
				addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
				expectedSQL := `^INSERT INTO "variations" \(.+\) VALUES \(.+\) RETURNING "id"$`
				mock.ExpectBegin()
				mock.ExpectQuery(expectedSQL).
					WithArgs(
						sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(addRow)
				mock.ExpectCommit()
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			err := productRepo.SaveVariation(echo.New().AcquireContext(), tt.inputCategory, tt.inputVariation)

			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductRepository_SaveVariationOption(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputVariationID  uint
		inputVariationVal string
		prepareMockFn     func()
		wantErr           error
	}{
		"Normal Case: SaveVariationOption": {
			inputVariationID:  1,
			inputVariationVal: "Test Variation Option",
			prepareMockFn: func() {
				addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
				expectedSQL := `^INSERT INTO "variation_options" \(.+\) VALUES \(.+\) RETURNING "id"$`
				mock.ExpectBegin()
				mock.ExpectQuery(expectedSQL).
					WithArgs(
						sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(addRow)
				mock.ExpectCommit()
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			err := productRepo.SaveVariationOption(echo.New().AcquireContext(), tt.inputVariationID, tt.inputVariationVal)

			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductRepository_FindProductByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		input         uint
		prepareMockFn func()
		want          response.Product
	}{
		"Normal Case: FindProductByID": {
			input: 1,
			prepareMockFn: func() {
				expectedSQL := `^SELECT \* FROM "products" WHERE id = (.+)$`
				mock.ExpectQuery(expectedSQL).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Test Product"))
			},
			want: response.Product{Name: "Test Product"},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			product, err := productRepo.FindProductByID(echo.New().AcquireContext(), tt.input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(product, tt.want) {
				t.Errorf("expected %v, but got %v", tt.want, product)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductRepository_IsProductNameExistForOtherProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputName     string
		inputProduct  uint
		prepareMockFn func()
		want          bool
	}{
		"Normal Case: IsProductNameExistForOtherProduct": {
			inputName:    "Test Product",
			inputProduct: 1,
			prepareMockFn: func() {
				expectedSQL := `^SELECT \* FROM "products" WHERE name = (.+) AND id != (.+)$`
				mock.ExpectQuery(expectedSQL).WithArgs("Test Product", 1).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Test Product"))
			},
			want: true,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			exist, err := productRepo.IsProductNameExistForOtherProduct(echo.New().AcquireContext(), tt.inputName, tt.inputProduct)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if exist != tt.want {
				t.Errorf("expected %v, but got %v", tt.want, exist)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductRepository_IsProductNameExist(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputName     string
		prepareMockFn func()
		want          bool
	}{
		"Normal Case: IsProductNameExist": {
			inputName: "Test Product",
			prepareMockFn: func() {
				expectedSQL := `^SELECT \* FROM "products" WHERE name = (.+)$`
				mock.ExpectQuery(expectedSQL).WithArgs("Test Product").WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Test Product"))
			},
			want: true,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			exist, err := productRepo.IsProductNameExist(echo.New().AcquireContext(), tt.inputName)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if exist != tt.want {
				t.Errorf("expected %v, but got %v", tt.want, exist)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductRepository_SaveProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	type input struct {
		product domain.Product
	}

	type output struct {
		wantErr error
	}

	tests := map[string]struct {
		input         input
		prepareMockFn func()
		want          output
	}{
		"Normal Case: SaveProduct": {
			input{
				product: domain.Product{
					Name:        "Test Product",
					Description: "This is a test product",
					CategoryID:  1,
					BrandID:     1,
					Price:       1000,
					Image:       "test_image.jpg",
				},
			},
			func() {
				addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
				expectedSQL := `^INSERT INTO "products" \(.+\) VALUES \(.+\) RETURNING "id"$`
				mock.ExpectBegin()
				mock.ExpectQuery(expectedSQL).
					WithArgs(
						sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
						sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
						sqlmock.AnyArg()).
					WillReturnRows(addRow)
				mock.ExpectCommit()
			},
			output{wantErr: nil},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			err := productRepo.SaveProduct(echo.New().AcquireContext(), tt.input.product)
			fmt.Print("err")
			fmt.Print(err)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductRepository_UpdateProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputProduct  domain.Product
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: UpdateProduct": {
			inputProduct: domain.Product{ID: 1, Name: "Updated Product"},
			prepareMockFn: func() {
				// addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
				expectedSQL := `UPDATE "products" SET "name"=.+,"updated_at"=.+ WHERE id =.+ AND "id" =.+`
				mock.ExpectBegin()
				mock.ExpectExec(expectedSQL).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			err := productRepo.UpdateProduct(echo.New().AcquireContext(), tt.inputProduct)

			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductRepository_FindAllProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputPagination request.Pagination
		inputCategoryID *uint
		inputBrandID    *uint
		prepareMockFn   func()
		want            []response.Product
	}{
		"Normal Case: FindAllProducts": {
			inputPagination: request.Pagination{Count: 10, PageNumber: 30},
			inputCategoryID: nil,
			inputBrandID:    nil,
			prepareMockFn: func() {
				expectedSQL := `SELECT (.+) FROM products p INNER JOIN categories sc ON p\.category_id = sc\.id INNER JOIN brands b ON b\.id = p\.brand_id ORDER BY created_at DESC LIMIT (.+) OFFSET (.+)`
				mock.ExpectQuery(expectedSQL).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Test Product"))
			},
			want: []response.Product{{Name: "Test Product"}},
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			products, err := productRepo.FindAllProducts(echo.New().AcquireContext(), tt.inputPagination, tt.inputCategoryID, tt.inputBrandID)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(products, tt.want) {
				t.Errorf("expected %v, but got %v", tt.want, products)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductDatabase_FindProductItemByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputProductItemID uint
		prepareMockFn      func()
		wantErr            error
	}{
		"Normal Case: FindProductItemByID": {
			inputProductItemID: 1,
			prepareMockFn: func() {
				expectedSQL := `^SELECT \* FROM "product_items" WHERE id = (.+)$`
				mock.ExpectQuery(expectedSQL).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err = productRepo.FindProductItemByID(echo.New().AcquireContext(), tt.inputProductItemID)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductDatabase_FindVariationCountForProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputProductID uint
		prepareMockFn  func()
		wantErr        error
	}{
		"Normal Case: FindVariationCountForProduct": {
			inputProductID: 1,
			prepareMockFn: func() {
				expectedSQL := `SELECT count\(\*\) FROM variations v INNER JOIN categories c ON c\.id = v\.category_id INNER JOIN products p ON p\.category_id = v\.category_id WHERE p\.id = (.+)$`
				mock.ExpectQuery(expectedSQL).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err = productRepo.FindVariationCountForProduct(echo.New().AcquireContext(), tt.inputProductID)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductDatabase_FindAllProductItemIDsByProductIDAndVariationOptionID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputProductID         uint
		inputVariationOptionID uint
		prepareMockFn          func()
		wantErr                error
	}{
		"Normal Case: FindAllProductItemIDsByProductIDAndVariationOptionID": {
			inputProductID:         1,
			inputVariationOptionID: 1,
			prepareMockFn: func() {
				expectedSQL := `SELECT \* FROM product_items pi INNER JOIN product_configurations pc ON pi\.id = pc\.product_item_id WHERE pi\.product_id = \$1 AND variation_option_id = \$2$`
				mock.ExpectQuery(expectedSQL).WithArgs(1, 1).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err = productRepo.FindAllProductItemIDsByProductIDAndVariationOptionID(context.Background(), tt.inputProductID, tt.inputVariationOptionID)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductDatabase_SaveProductConfiguration(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputProductItemID     uint
		inputVariationOptionID uint
		prepareMockFn          func()
		wantErr                error
	}{
		"Normal Case: SaveProductConfiguration": {
			inputProductItemID:     1,
			inputVariationOptionID: 1,
			prepareMockFn: func() {
				expectedSQL := `^INSERT INTO "product_configurations" \("product_item_id","variation_option_id"\) VALUES \(\$1,\$2\)$`
				mock.ExpectBegin()
				mock.ExpectExec(expectedSQL).WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			err = productRepo.SaveProductConfiguration(echo.New().AcquireContext(), tt.inputProductItemID, tt.inputVariationOptionID)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductDatabase_SaveProductItem(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputProductItem domain.ProductItem
		prepareMockFn    func()
		wantErr          error
	}{
		"Normal Case: SaveProductItem": {
			inputProductItem: domain.ProductItem{Name: "Test Product Item"},
			prepareMockFn: func() {
				addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
				expectedSQL := `^INSERT INTO "product_items" \(.+\) VALUES \(.+\) RETURNING "id"$`
				expectedSQL1 := `SELECT \* FROM "product_items" WHERE "product_items"\."id" = \$1`
				mock.ExpectBegin()
				mock.ExpectQuery(expectedSQL).
					WithArgs(
						sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
						sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(addRow)
				mock.ExpectCommit()
				mock.ExpectQuery(expectedSQL1).
					WithArgs(
						sqlmock.AnyArg()).
					WillReturnRows(addRow)
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err = productRepo.SaveProductItem(echo.New().AcquireContext(), tt.inputProductItem)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductDatabase_FindAllProductItems(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputProductID uint
		prepareMockFn  func()
		wantErr        error
	}{
		"Normal Case: FindAllProductItems": {
			inputProductID: 1,
			prepareMockFn: func() {
				expectedSQL := `SELECT p\.name, pi\.id, pi\.product_id,pi\.name AS item_name, pi\.price, pi\.discount_price, pi\.qty_in_stock, pi\.sku, p\.category_id, sc\.name AS category_name, p\.brand_id, b\.name AS brand_name FROM product_items pi INNER JOIN products p ON p\.id = pi\.product_id INNER JOIN categories sc ON p\.category_id = sc\.id INNER JOIN brands b ON b\.id = p\.brand_id WHERE pi\.product_id = \$1`
				mock.ExpectQuery(expectedSQL).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"name"}).AddRow("Test Product Item"))
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err = productRepo.FindAllProductItems(echo.New().AcquireContext(), tt.inputProductID)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductDatabase_FindAllVariationValuesOfProductItem(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputProductItemID uint
		prepareMockFn      func()
		wantErr            error
	}{
		"Normal Case: FindAllVariationValuesOfProductItem": {
			inputProductItemID: 1,
			prepareMockFn: func() {
				expectedSQL := `SELECT v\.id AS variation_id, v\.name, vo\.id AS variation_option_id, vo\.value FROM product_configurations pc INNER JOIN variation_options vo ON vo\.id = pc\.variation_option_id INNER JOIN variations v ON v\.id = vo\.variation_id WHERE pc\.product_item_id = \$1`
				mock.ExpectQuery(expectedSQL).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"variation_id", "name", "variation_option_id", "value"}).AddRow(1, "Test Variation", 1, "Test Value"))
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err = productRepo.FindAllVariationValuesOfProductItem(echo.New().AcquireContext(), tt.inputProductItemID)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductDatabase_SaveProductItemImage(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputProductItemID uint
		inputImage         string
		prepareMockFn      func()
		wantErr            error
	}{
		"Normal Case: SaveProductItemImage": {
			inputProductItemID: 1,
			inputImage:         "test_image.jpg",
			prepareMockFn: func() {
				addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
				expectedSQL := `^INSERT INTO "product_images" \(.+\) VALUES \(.+\) RETURNING "id"$`
				mock.ExpectBegin()
				mock.ExpectQuery(expectedSQL).
					WithArgs(
						sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(addRow)
				mock.ExpectCommit()
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			err = productRepo.SaveProductItemImage(echo.New().AcquireContext(), tt.inputProductItemID, tt.inputImage)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestProductDatabase_FindAllProductItemImages(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	productRepo := repository.NewProductRepository(gormDB)

	tests := map[string]struct {
		inputProductItemID uint
		prepareMockFn      func()
		wantErr            error
	}{
		"Normal Case: FindAllProductItemImages": {
			inputProductItemID: 1,
			prepareMockFn: func() {
				expectedSQL := `SELECT image FROM "product_images" WHERE product_item_id = \$1`
				mock.ExpectQuery(expectedSQL).WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"image"}).AddRow("test_image.jpg"))
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err = productRepo.FindAllProductItemImages(echo.New().AcquireContext(), tt.inputProductItemID)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
