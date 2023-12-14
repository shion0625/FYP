package domain

import (
	"time"
)

type UserAddress struct {
	ID        uint    `gorm:"primaryKey;unique" json:"id"`
	UserID    string  `gorm:"not null"          json:"userId"`
	User      User    `json:"-"`
	AddressID uint    `gorm:"not null"          json:"addressId"`
	Address   Address `json:"-"`
	IsDefault bool    `json:"isDefault"`
}

type User struct {
	ID          string    `gorm:"primaryKey;size:255;default:gen_random_uuid()" json:"id"`
	Age         uint      `json:"age"`
	GoogleImage string    `json:"googleImage"`
	FirstName   string    `gorm:"not null"                                      json:"firstName"`
	LastName    string    `gorm:"not null"                                      json:"lastName"`
	UserName    string    `gorm:"not null;unique"                               json:"userName"`
	Email       string    `gorm:"unique;not null"                               json:"email"`
	Phone       string    `gorm:"unique"                                        json:"phone"`
	Password    string    `json:"password"`
	Verified    bool      `gorm:"default:false"                                 json:"verified"`
	BlockStatus bool      `gorm:"not null;default:false"                        json:"blockStatus"`
	CreatedAt   time.Time `gorm:"not null"                                      json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// many to many join.
type Address struct {
	ID          uint      `gorm:"primaryKey;unique" json:"id"`
	Name        string    `gorm:"not null"          json:"name"`
	PhoneNumber string    `gorm:"not null"          json:"phoneNumber"`
	House       string    `gorm:"not null"          json:"house"` // address line
	Area        string    `gorm:"not null"          json:"area"`  // state
	LandMark    string    `gorm:"not null"          json:"landMark"`
	City        string    `gorm:"not null"          json:"city"`
	Pincode     string    `gorm:"not null"          json:"pincode"`
	CountryName string    `gorm:"not null"          json:"countryName"`
	CreatedAt   time.Time `gorm:"not null"          json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type PaymentMethod struct {
	ID          uint      `gorm:"primaryKey;not null" json:"id"`
	Number      string    `gorm:"unique;not null"     json:"number"`
	Expiry      string    `gorm:"unique;not null"     json:"expiry"`
	Cvc         string    `gorm:"not null"            json:"cvc"`
	CardCompany string    `gorm:"not null"     json:"cardCompany"`
	UserId      string    `gorm:"not null"            json:"userId"`
	CreatedAt   time.Time `gorm:"not null"            json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
