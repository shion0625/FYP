package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserRepository_FindUserByUserName(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	userRepo := repository.NewUserRepository(gormDB)

	tests := map[string]struct {
		userName      string
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: FindUserByUserName": {
			userName: "testuser",
			prepareMockFn: func() {
				rows := sqlmock.NewRows([]string{"id", "user_name", "email"}).
					AddRow("1", "testuser", "test@example.com")
				mock.ExpectQuery(`^SELECT \* FROM "users" WHERE user_name = (.+) ORDER BY "users"."id" LIMIT 1$`).WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err := userRepo.FindUserByUserName(echo.New().AcquireContext(), tt.userName)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepository_FindUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	userRepo := repository.NewUserRepository(gormDB)

	tests := map[string]struct {
		email         string
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: FindUserByEmail": {
			email: "test@example.com",
			prepareMockFn: func() {
				rows := sqlmock.NewRows([]string{"id", "user_name", "email"}).
					AddRow("1", "testuser", "test@example.com")
				mock.ExpectQuery(`^SELECT \* FROM "users" WHERE email = (.+) ORDER BY "users"."id" LIMIT 1$`).WithArgs("test@example.com").WillReturnRows(rows)
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err := userRepo.FindUserByEmail(echo.New().AcquireContext(), tt.email)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepository_FindUserByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	userRepo := repository.NewUserRepository(gormDB)

	tests := map[string]struct {
		userID        string
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: FindUserByUserID": {
			userID: "1",
			prepareMockFn: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "email"}).
					AddRow("1", "Test User", "test@example.com")
				mock.ExpectQuery(`SELECT \* FROM "users" WHERE id = \$1 ORDER BY "users"\."id" LIMIT 1`).WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err := userRepo.FindUserByUserID(echo.New().AcquireContext(), tt.userID)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepository_FindUserByPhoneNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	userRepo := repository.NewUserRepository(gormDB)

	tests := map[string]struct {
		phoneNumber   string
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: FindUserByPhoneNumber": {
			phoneNumber: "1234567890",
			prepareMockFn: func() {
				rows := sqlmock.NewRows([]string{"id", "user_name", "phone"}).
					AddRow("1", "testuser", "1234567890")
				mock.ExpectQuery(`^SELECT \* FROM "users" WHERE phone =(.+)$`).WithArgs("1234567890").WillReturnRows(rows)
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err := userRepo.FindUserByPhoneNumber(echo.New().AcquireContext(), tt.phoneNumber)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepository_FindUserByUserNameEmailOrPhone(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	userRepo := repository.NewUserRepository(gormDB)

	tests := map[string]struct {
		userDetails   domain.User
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: FindUserByUserNameEmailOrPhone": {
			userDetails: domain.User{
				UserName: "testuser",
				Email:    "test@example.com",
				Phone:    "1234567890",
			},
			prepareMockFn: func() {
				rows := sqlmock.NewRows([]string{"id", "user_name", "email", "phone"}).
					AddRow("1", "testuser", "test@example.com", "1234567890")
				mock.ExpectQuery(`^SELECT \* FROM "users" WHERE user_name = (.+) OR email = (.+) OR phone = (.+)`).
					WithArgs("testuser", "test@example.com", "1234567890").WillReturnRows(rows)
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err := userRepo.FindUserByUserNameEmailOrPhone(echo.New().AcquireContext(), tt.userDetails)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepository_SaveUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	userRepo := repository.NewUserRepository(gormDB)

	tests := map[string]struct {
		user          domain.User
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: SaveUser": {
			user: domain.User{
				ID:          "test-id",
				Age:         30,
				GoogleImage: "test-google-image",
				FirstName:   "Test",
				LastName:    "User",
				UserName:    "testuser",
				Email:       "testuser@example.com",
				Phone:       "1234567890",
				Password:    "password",
				Verified:    false,
				BlockStatus: false,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			prepareMockFn: func() {
				addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
				expectedSQL := `^INSERT INTO "users" \(.+\) VALUES \(.+\) RETURNING "id"$`
				mock.ExpectBegin()
				mock.ExpectQuery(expectedSQL).
					WithArgs(
						sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(addRow)
				mock.ExpectCommit()
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err := userRepo.SaveUser(echo.New().AcquireContext(), tt.user)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepository_FindAllAddressByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	userRepo := repository.NewUserRepository(gormDB)

	tests := map[string]struct {
		userID        string
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: FindAllAddressByUserID": {
			userID: "1",
			prepareMockFn: func() {
				rows := sqlmock.NewRows([]string{"id", "house", "name", "phone_number", "area", "land_mark", "city", "pincode", "country_name", "is_default"}).
					AddRow("1", "testhouse", "testname", "1234567890", "testarea", "testlandmark", "testcity", "123456", "testcountry", true)
				mock.ExpectQuery(`^SELECT (.+) FROM "user_addresses" JOIN addresses ON user_addresses.address_id = addresses.id WHERE user_addresses.user_id = (.+)$`).WithArgs("1").WillReturnRows(rows)
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err := userRepo.FindAllAddressByUserID(echo.New().AcquireContext(), tt.userID)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepository_FindAddressByUserIDAndAddressID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	userRepo := repository.NewUserRepository(gormDB)

	tests := map[string]struct {
		userID        string
		addressID     uint
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: FindAddressByUserIDAndAddressID": {
			userID:    "1",
			addressID: 1,
			prepareMockFn: func() {
				rows := sqlmock.NewRows([]string{"id", "house", "name", "phone_number", "area", "land_mark", "city", "pincode", "country_name", "is_default"}).
					AddRow("1", "testhouse", "testname", "1234567890", "testarea", "testlandmark", "testcity", "123456", "testcountry", true)
				mock.ExpectQuery(`^SELECT (.+) FROM "user_addresses" JOIN addresses ON user_addresses.address_id = addresses.id WHERE user_addresses.user_id = (.+) AND addresses.id = (.+)`).WithArgs("1", 1).WillReturnRows(rows)
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err := userRepo.FindAddressByUserIDAndAddressID(echo.New().AcquireContext(), tt.userID, tt.addressID)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepository_IsAddressAlreadyExistForUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	userRepo := repository.NewUserRepository(gormDB)

	tests := map[string]struct {
		address       domain.Address
		userID        string
		prepareMockFn func()
		wantErr       error
		wantExist     bool
	}{
		"Normal Case: IsAddressAlreadyExistForUser": {
			address: domain.Address{
				Name:     "testname",
				House:    "testhouse",
				LandMark: "testlandmark",
				Pincode:  "123456",
			},
			userID: "1",
			prepareMockFn: func() {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow("1")
				mock.ExpectQuery(`^SELECT addresses.id FROM "user_addresses" INNER JOIN addresses ON user_addresses.address_id = addresses.id WHERE addresses.name = (.+) AND addresses.house = (.+) AND addresses.land_mark = (.+) AND addresses.pincode = (.+) AND user_addresses.user_id = (.+)$`).WithArgs("testname", "testhouse", "testlandmark", "123456", "1").WillReturnRows(rows)
			},
			wantErr:   nil,
			wantExist: true,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			exist, err := userRepo.IsAddressAlreadyExistForUser(echo.New().AcquireContext(), tt.address, tt.userID)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if exist != tt.wantExist {
				t.Errorf("expected %v, but got %v", tt.wantExist, exist)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepository_IsAddressIDExist(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	userRepo := repository.NewUserRepository(gormDB)

	tests := map[string]struct {
		addressID     uint
		prepareMockFn func()
		wantErr       error
		wantExist     bool
	}{
		"Normal Case: IsAddressIDExist": {
			addressID: 1,
			prepareMockFn: func() {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow("1")
				mock.ExpectQuery(`^SELECT addresses.id FROM "addresses" WHERE id = (.+)$`).WithArgs(1).WillReturnRows(rows)
			},
			wantErr:   nil,
			wantExist: true,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			exist, err := userRepo.IsAddressIDExist(echo.New().AcquireContext(), tt.addressID)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if exist != tt.wantExist {
				t.Errorf("expected %v, but got %v", tt.wantExist, exist)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepository_SaveAddress(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	userRepo := repository.NewUserRepository(gormDB)

	tests := map[string]struct {
		address       domain.Address
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: SaveAddress": {
			address: domain.Address{
				Name:        "Test Name",
				PhoneNumber: "1234567890",
				House:       "Test House",
				Area:        "Test Area",
				LandMark:    "Test LandMark",
				City:        "Test City",
				Pincode:     "123456",
				CountryName: "Test Country",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			prepareMockFn: func() {
				addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
				expectedSQL := `^INSERT INTO "addresses" (.+) VALUES (.+) RETURNING "id"$`
				mock.ExpectBegin()
				mock.ExpectQuery(expectedSQL).
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(addRow)
				mock.ExpectCommit()
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err := userRepo.SaveAddress(echo.New().AcquireContext(), tt.address)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepository_UpdateAddress(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	userRepo := repository.NewUserRepository(gormDB)

	tests := map[string]struct {
		address       domain.Address
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: UpdateAddress": {
			address: domain.Address{
				ID:          1,
				Name:        "Updated Name",
				PhoneNumber: "Updated PhoneNumber",
				House:       "Updated House",
				Area:        "Updated Area",
				LandMark:    "Updated LandMark",
				City:        "Updated City",
				Pincode:     "Updated Pincode",
				CountryName: "Updated CountryName",
			},
			prepareMockFn: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`UPDATE "addresses" SET "area"=\$1,"city"=\$2,"country_name"=\$3,"house"=\$4,"land_mark"=\$5,"name"=\$6,"phone_number"=\$7,"pincode"=\$8,"updated_at"=\$9 WHERE id = \$10 AND "id" = \$11`).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			err := userRepo.UpdateAddress(echo.New().AcquireContext(), tt.address)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepository_UpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	userRepo := repository.NewUserRepository(gormDB)

	tests := map[string]struct {
		user          domain.User
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: UpdateUser": {
			user: domain.User{
				ID:        "testUser",
				Age:       30,
				FirstName: "Updated FirstName",
				LastName:  "Updated LastName",
				UserName:  "Updated UserName",
				Email:     "Updated Email",
				Phone:     "Updated Phone",
				Password:  "Updated Password",
			},
			prepareMockFn: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`UPDATE "users" SET "age"=\$1,"email"=\$2,"first_name"=\$3,"last_name"=\$4,"password"=\$5,"phone"=\$6,"updated_at"=\$7,"user_name"=\$8 WHERE id = \$9 AND "id" = \$10`).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			err := userRepo.UpdateUser(echo.New().AcquireContext(), tt.user)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepository_UpdateUserAddress(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	userRepo := repository.NewUserRepository(gormDB)

	tests := map[string]struct {
		userAddress   domain.UserAddress
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: UpdateUserAddress": {
			userAddress: domain.UserAddress{
				UserID:    "testUser",
				AddressID: 1,
				IsDefault: true,
			},
			prepareMockFn: func() {
				mock.ExpectBegin()
				mock.ExpectExec(`UPDATE "user_addresses" SET "is_default"=(.+) WHERE user_id = (.+)`).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
				mock.ExpectBegin()
				mock.ExpectExec(`UPDATE "user_addresses" SET "is_default"=\$1 WHERE address_id = \$2 AND user_id = \$3`).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(sqlmock.NewResult(1, 1))

				mock.ExpectCommit()
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			err := userRepo.UpdateUserAddress(echo.New().AcquireContext(), tt.userAddress)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepository_FindAllPaymentMethodsByUserID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	userRepo := repository.NewUserRepository(gormDB)

	tests := map[string]struct {
		userID        string
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: FindAllPaymentMethodsByUserID": {
			userID: "testUser",
			prepareMockFn: func() {
				rows := sqlmock.NewRows([]string{"id", "number", "card_company"}).
					AddRow(1, "1234567890123456", "Test Company")
				mock.ExpectQuery(`SELECT id, number,card_company FROM "payment_methods" WHERE user_id = \$1`).WithArgs("testUser").WillReturnRows(rows)
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err := userRepo.FindAllPaymentMethodsByUserID(echo.New().AcquireContext(), tt.userID)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestUserRepository_SavePaymentMethod(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	userRepo := repository.NewUserRepository(gormDB)

	tests := map[string]struct {
		paymentMethod domain.PaymentMethod
		prepareMockFn func()
		wantErr       error
	}{
		"Normal Case: SavePaymentMethod": {
			paymentMethod: domain.PaymentMethod{
				ID:          1,
				Number:      "1234567890123456",
				Expiry:      "12/24",
				Cvc:         "123",
				CardCompany: "Test Company",
				UserId:      "testUser",
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			},
			prepareMockFn: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(`^INSERT INTO "payment_methods" (.+) VALUES (.+) RETURNING "id"$`).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectCommit()
			},
			wantErr: nil,
		},
	}

	for testName, tt := range tests {
		t.Run(testName, func(t *testing.T) {
			tt.prepareMockFn()

			_, err := userRepo.SavePaymentMethod(echo.New().AcquireContext(), tt.paymentMethod)
			if err != tt.wantErr {
				t.Errorf("expected %v, but got %v", tt.wantErr, err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
