package domain

import (
	"time"
)

type User struct {
	ID          uint      `gorm:"primaryKey;unique"         json:"id"`
	Age         uint      `binding:"required,numeric"       json:"age"`
	GoogleImage string    `json:"googleImage"`
	FirstName   string    `binding:"required,min=2,max=50"  gorm:"not null"        json:"firstName"`
	LastName    string    `binding:"required,min=1,max=50"  gorm:"not null"        json:"lastName"`
	UserName    string    `binding:"required,min=3,max=15"  gorm:"not null;unique" json:"userName"`
	Email       string    `binding:"required,email"         gorm:"unique;not null" json:"email"`
	Phone       string    `binding:"required,min=10,max=10" gorm:"unique"          json:"phone"`
	Password    string    `binding:"required"               json:"password"`
	Verified    bool      `gorm:"default:false"             json:"verified"`
	BlockStatus bool      `gorm:"not null;default:false"    json:"blockStatus"`
	CreatedAt   time.Time `gorm:"not null"                  json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// many to many join.
type UserAddress struct {
	ID        uint `gorm:"primaryKey;unique" json:"id"`
	UserID    uint `gorm:"not null"          json:"userId"`
	User      User
	AddressID uint `gorm:"not null" json:"addressId"`
	Address   Address
	IsDefault bool `json:"isDefault"`
}

type Address struct {
	ID          uint   `gorm:"primaryKey;unique"               json:"id"`
	Name        string `binding:"required,min=2,max=50"        gorm:"not null" json:"name"`
	PhoneNumber string `binding:"required,min=10,max=10"       gorm:"not null" json:"phoneNumber"`
	House       string `binding:"required"                     gorm:"not null" json:"house"`
	Area        string `gorm:"not null"                        json:"area"`
	LandMark    string `binding:"required"                     gorm:"not null" json:"landMark"`
	City        string `gorm:"not null"                        json:"city"`
	Pincode     uint   `binding:"required,numeric,min=6,max=6" gorm:"not null" json:"pincode"`
	CountryID   uint   `binding:"required"                     gorm:"not null" json:"countryId"`
	Country     Country
	CreatedAt   time.Time `gorm:"not null"  json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type Country struct {
	ID          uint   `gorm:"primaryKey;unique;" json:"id"`
	CountryName string `gorm:"unique;not null"    json:"countryName"`
}
