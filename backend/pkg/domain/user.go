package domain

import (
	"time"
)

type UserAddress struct {
	ID        uint   `gorm:"primaryKey;unique" json:"id"`
	UserID    string `gorm:"not null"          json:"userId"`
	User      User
	AddressID uint `gorm:"not null" json:"addressId"`
	Address   Address
	IsDefault bool `json:"isDefault"`
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
	ID          uint   `gorm:"primaryKey;unique" json:"id"`
	Name        string `gorm:"not null"          json:"name"`
	PhoneNumber string `gorm:"not null"          json:"phoneNumber"`
	House       string `gorm:"not null"          json:"house"`
	Area        string `gorm:"not null"          json:"area"`
	LandMark    string `gorm:"not null"          json:"landMark"`
	City        string `gorm:"not null"          json:"city"`
	Pincode     uint   `gorm:"not null"          json:"pincode"`
	CountryID   uint   `gorm:"not null"          json:"countryId"`
	Country     Country
	CreatedAt   time.Time `gorm:"not null"  json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Country struct {
	ID          uint   `gorm:"primaryKey;unique;" json:"id"`
	CountryName string `gorm:"unique;not null"    json:"countryName"`
}

type PaymentMethod struct {
	ID           uint      `gorm:"primaryKey;not null" json:"id"`
	CreditNumber string    `gorm:"unique;not null"     json:"creditNumber"`
	Cvv          string    `gorm:"not null"            json:"cvv"`
	UserId       string    `gorm:"not null"            json:"userId"`
	CreatedAt    time.Time `gorm:"not null"            json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
