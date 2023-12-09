package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/shion0625/FYP/backend/pkg/api/handler/response"
	"github.com/shion0625/FYP/backend/pkg/domain"
	"github.com/shion0625/FYP/backend/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB: db}
}

func (c *userDatabase) FindUserByUserID(ctx echo.Context, userID string) (user domain.User, err error) {
	err = c.DB.Where("id = ?", userID).First(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByUserName(ctx echo.Context, userName string) (user domain.User, err error) {
	err = c.DB.Where("user_name = ?", userName).First(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByEmail(ctx echo.Context, email string) (user domain.User, err error) {
	err = c.DB.Where("email = ?", email).First(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByPhoneNumber(ctx echo.Context, phoneNumber string) (user domain.User, err error) {
	err = c.DB.Where("phone = ?", phoneNumber).First(&user).Error

	return user, err
}

func (c *userDatabase) FindUserByUserNameEmailOrPhone(ctx echo.Context,
	userDetails domain.User,
) (user domain.User, err error) {
	err = c.DB.Where("user_name = ? OR email = ? OR phone = ?",
		userDetails.UserName, userDetails.Email, userDetails.Phone).Find(&user).Error

	return user, err
}

func (c *userDatabase) SaveUser(ctx echo.Context, user domain.User) (userID string, err error) {
	// save the user details
	user = domain.User{
		Age:         user.Age,
		GoogleImage: user.GoogleImage,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		UserName:    user.UserName,
		Email:       user.Email,
		Phone:       user.Phone,
		Password:    user.Password,
		CreatedAt:   time.Now(),
	}
	result := c.DB.Create(&user)

	return user.ID, result.Error
}

func (c *userDatabase) FindAllAddressByUserID(ctx echo.Context, userID string) (addresses []response.Address, err error) {
	err = c.DB.
		Table("user_addresses").
		Select("addresses.id, addresses.house, addresses.name, addresses.phone_number, addresses.area, addresses.land_mark, addresses.city, addresses.pincode, addresses.country_id, countries.country_name, user_addresses.is_default").
		Joins("JOIN addresses ON user_addresses.address_id = addresses.id").
		Joins("JOIN countries ON addresses.country_id = countries.id").
		Where("user_addresses.user_id = ?", userID).
		Scan(&addresses).Error

	return addresses, err
}

func (c *userDatabase) IsAddressAlreadyExistForUser(ctx echo.Context, address domain.Address, userID string) (exist bool, err error) {
	address.CountryID = 1 // hardcoded !!!! should change

	err = c.DB.
		Select("CASE WHEN addresses.id != 0 THEN 'T' ELSE 'F' END AS exist").
		Joins("INNER JOIN user_addresses ON addresses.id = user_addresses.address_id").
		Where("addresses.name = ? AND addresses.house = ? AND addresses.land_mark = ? AND addresses.pincode = ? AND addresses.country_id = ? AND user_addresses.user_id = ?", address.Name, address.House, address.LandMark, address.Pincode, address.CountryID, userID).
		First(&exist).Error

	return exist, err
}

func (c *userDatabase) IsAddressIDExist(ctx echo.Context, addressID uint) (exist bool, err error) {
	err = c.DB.
		Select("EXISTS(SELECT 1 FROM addresses WHERE id = ?) AS exist", addressID).
		Where("id = ?", addressID).
		First(&exist).Error

	return exist, err
}

func (c *userDatabase) SaveAddress(ctx echo.Context, address domain.Address) (addressID uint, err error) {
	address.CreatedAt = time.Now()
	result := c.DB.Create(&address)

	if result.Error != nil {
		return addressID, errors.New("failed to insert address on database")
	}

	return address.ID, nil
}

func (c *userDatabase) SaveUserAddress(ctx echo.Context, userAddress domain.UserAddress) error {
	// first check user's first address is this or not
	var userID uint
	if err := c.DB.Table("user_addresses").Select("address_id").Where("user_id = ?", userAddress.UserID).Scan(&userID).Error; err != nil {
		return fmt.Errorf("failed to check if user already has an address with user_id %v", userAddress.UserID)
	}

	// if the given address needs to be set as default, then remove all others from default
	if userID == 0 { // it means user has no other addresses
		userAddress.IsDefault = true
	} else if userAddress.IsDefault {
		query := `UPDATE user_addresses SET is_default = 'f' WHERE user_id = ?`
		if c.DB.Exec(query, userAddress.UserID).Scan(&userAddress).Error != nil {
			return errors.New("failed to remove default status of address")
		}
	}

	// insert the user address
	if err := c.DB.Table("user_addresses").Create(&userAddress).Error; err != nil {
		return errors.New("failed to insert userAddress into database")
	}

	return nil
}

func (c *userDatabase) UpdateAddress(ctx echo.Context, address domain.Address) error {
	address.CountryID = 1 // hardcoded !!!! should change
	if err := c.DB.Model(&address).Where("id = ?", address.ID).Updates(map[string]interface{}{
		"name":         address.Name,
		"phone_number": address.PhoneNumber,
		"house":        address.House,
		"area":         address.Area,
		"land_mark":    address.LandMark,
		"city":         address.City,
		"pincode":      address.Pincode,
		"country_id":   address.CountryID,
		"updated_at":   time.Now(),
	}).Error; err != nil {
		return errors.New("failed to update the address for edit address")
	}

	return nil
}

func (c *userDatabase) UpdateUser(ctx echo.Context, user domain.User) (err error) {
	updatedAt := time.Now()
	// check password need to update or not
	if user.Password != "" {
		err = c.DB.Model(&user).Where("id = ?", user.ID).Updates(map[string]interface{}{
			"user_name":  user.UserName,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"age":        user.Age,
			"email":      user.Email,
			"phone":      user.Phone,
			"password":   user.Password,
			"updated_at": updatedAt,
		}).Error
	} else {
		err = c.DB.Model(&user).Where("id = ?", user.ID).Updates(map[string]interface{}{
			"user_name":  user.UserName,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"age":        user.Age,
			"email":      user.Email,
			"phone":      user.Phone,
			"updated_at": updatedAt,
		}).Error
	}

	if err != nil {
		return fmt.Errorf("failed to update user detail of user with user_id %s", user.ID)
	}

	return nil
}

func (c *userDatabase) UpdateUserAddress(ctx echo.Context, userAddress domain.UserAddress) error {
	// if it needs to be set as default, then change the old default
	if userAddress.IsDefault {
		if err := c.DB.Model(&domain.UserAddress{}).Where("user_id = ?", userAddress.UserID).Update("is_default", false).Error; err != nil {
			return errors.New("failed to remove default status of address")
		}
	}

	// update the user address
	if err := c.DB.Model(&userAddress).Where("address_id = ? AND user_id = ?", userAddress.AddressID, userAddress.UserID).Update("is_default", userAddress.IsDefault).Error; err != nil {
		return errors.New("failed to update user address")
	}

	return nil
}
