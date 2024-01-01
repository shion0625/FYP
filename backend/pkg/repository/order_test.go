package repository_test

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/request"
	"github.com/shion0625/FYP/backend/pkg/repository"
	"github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestOrderRepository_Transactions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	orderRepo := repository.NewOrderRepository(gormDB)

	tests := map[string]struct {
		trxFn         func(repo interfaces.OrderRepository) error
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: Transactions": {
			trxFn: func(repo interfaces.OrderRepository) error {
				return nil
			},
			prepareMockFn: func() {
				mock.ExpectBegin()
				mock.ExpectCommit()
			},
			wantErr: nil,
		},
		"Error Case: Transactions": {
			trxFn: func(repo interfaces.OrderRepository) error {
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

			err := orderRepo.Transactions(echo.New().AcquireContext(), tt.trxFn)

			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestOrderRepository_SaveOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	orderRepo := repository.NewOrderRepository(gormDB)

	tests := map[string]struct {
		inputUserID   string
		inputPayOrder request.PayOrder
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: SaveOrder": {
			inputUserID: "test_user",
			inputPayOrder: request.PayOrder{
				TotalFee:        1000,
				AddressID:       1,
				PaymentMethodID: 1,
				ProductItemInfo: []request.ProductItemInfo{
					{
						ProductItemID: 1,
						Count:         2,
						VariationValues: &[]request.VariationValues{
							{
								VariationID:       1,
								Name:              "Size",
								VariationOptionID: 1,
								Value:             "M",
							},
						},
					},
				},
			},
			prepareMockFn: func() {
				addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
				expectedSQL := `^INSERT INTO "shop_orders" \(.+\) VALUES \(.+\) RETURNING "id"$`
				expectedSQL2 := `^INSERT INTO "shop_order_variations" \(.+\) VALUES \(.+\) RETURNING "id"$`
				expectedSQL3 := `^INSERT INTO "shop_order_product_items" \(.+\) VALUES \(.+\) RETURNING "id"$`

				mock.ExpectBegin()
				mock.ExpectQuery(expectedSQL).
					WithArgs(
						sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
						sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(addRow)
				mock.ExpectCommit()
				mock.ExpectBegin()

				mock.ExpectQuery(expectedSQL2).
					WithArgs(
						sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(),
						sqlmock.AnyArg()).
					WillReturnRows(addRow)
				mock.ExpectCommit()
				mock.ExpectBegin()
				mock.ExpectQuery(expectedSQL3).
					WithArgs(
						sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(addRow)
				mock.ExpectCommit()
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			err := orderRepo.SaveOrder(echo.New().AcquireContext(), tt.inputUserID, tt.inputPayOrder)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestOrderRepository_UpdateProductItemStock(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	orderRepo := repository.NewOrderRepository(gormDB)

	tests := map[string]struct {
		inputProductItemID    uint
		inputPurchaseQuantity uint
		prepareMockFn         func()
		wantNewStock          uint
		wantErr               error
	}{
		"Normal Case: UpdateProductItemStock": {
			inputProductItemID:    1,
			inputPurchaseQuantity: 5,
			prepareMockFn: func() {
				mock.ExpectQuery(`SELECT (.+) FROM "product_items"`).WithArgs(sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"qty_in_stock"}).AddRow(10))
				expectedSQL := `UPDATE "product_items" SET "qty_in_stock"=.+ WHERE id =.+`
				mock.ExpectBegin()
				mock.ExpectExec(expectedSQL).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantNewStock: 5,
			wantErr:      nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			newStock, err := orderRepo.UpdateProductItemStock(nil, tt.inputProductItemID, tt.inputPurchaseQuantity)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if newStock != tt.wantNewStock {
				t.Errorf("expected %v, but got %v", tt.wantNewStock, newStock)
			}
		})
	}
}

func TestOrderRepository_PayOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	orderRepo := repository.NewOrderRepository(gormDB)

	tests := map[string]struct {
		inputPaymentMethodID uint
		prepareMockFn        func()
		wantErr              error
	}{
		"Normal Case: PayOrder": {
			inputPaymentMethodID: 1,
			prepareMockFn: func() {
				expectedSQL := `^SELECT \* FROM "payment_methods" WHERE "payment_methods"."id" = \$1 ORDER BY "payment_methods"."id" LIMIT 1$`
				mock.ExpectQuery(expectedSQL).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			err := orderRepo.PayOrder(echo.New().AcquireContext(), tt.inputPaymentMethodID)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestOrderRepository_GetShopOrders(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	orderRepo := repository.NewOrderRepository(gormDB)

	tests := map[string]struct {
		inputUserID     string
		inputPagination request.Pagination
		prepareMockFn   func()
		wantErr         error
	}{
		"Normal Case: GetShopOrders": {
			inputUserID: "test_user",
			inputPagination: request.Pagination{
				PageNumber: 20,
				Count:      10,
			},
			prepareMockFn: func() {
				expectedSQL := `^SELECT \* FROM "shop_orders" WHERE user_id= (.+) LIMIT (.+) OFFSET (.+)`
				expectedSQL1 := `^SELECT \* FROM "shop_order_product_items" WHERE shop_order_id = (.+)`
				expectedSQL2 := `^SELECT shop_order_variations.id, variations.name, shop_order_variations.variation_option_id, variation_options.value FROM "shop_order_variations" INNER JOIN variations ON shop_order_variations.variation_id = variations.id INNER JOIN variation_options ON shop_order_variations.variation_option_id = variation_options.id WHERE shop_order_id = (.+) AND product_item_id = (.+)`
				expectedSQL3 := `^SELECT \* FROM "product_items" WHERE id = (.+)`
				mock.ExpectQuery(expectedSQL).
					WithArgs("test_user").
					WillReturnRows(sqlmock.NewRows([]string{"id", "user_id"}).AddRow(1, "test_user"))
				mock.ExpectQuery(expectedSQL1).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectQuery(expectedSQL2).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"VariationID", "Name", "VariationOptionID", "Value"}).AddRow(1, "Size", 1, "M"))
				mock.ExpectQuery(expectedSQL3).
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err := orderRepo.GetShopOrders(echo.New().AcquireContext(), tt.inputUserID, tt.inputPagination)

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
