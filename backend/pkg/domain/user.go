package domain

import (
	"time"
)

type User struct {
	ID          uint      `gorm:"primaryKey;unique"         json:"id"`
	Age         uint      `binding:"required,numeric"       json:"age"`
	GoogleImage string    `json:"google_profile_image"`
	FirstName   string    `binding:"required,min=2,max=50"  gorm:"not null"        json:"first_name"`
	LastName    string    `binding:"required,min=1,max=50"  gorm:"not null"        json:"last_name"`
	UserName    string    `binding:"required,min=3,max=15"  gorm:"not null;unique" json:"user_name"`
	Email       string    `binding:"required,email"         gorm:"unique;not null" json:"email"`
	Phone       string    `binding:"required,min=10,max=10" gorm:"unique"          json:"phone"`
	Password    string    `binding:"required"               json:"password"`
	Verified    bool      `gorm:"default:false"             json:"verified"`
	BlockStatus bool      `gorm:"not null;default:false"    json:"block_status"`
	CreatedAt   time.Time `gorm:"not null"                  json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// many to many join.
type UserAddress struct {
	ID        uint `gorm:"primaryKey;unique" json:"id"`
	UserID    uint `gorm:"not null"          json:"user_id"`
	User      User
	AddressID uint `gorm:"not null" json:"address_id"`
	Address   Address
	IsDefault bool `json:"is_default"`
}

type Address struct {
	ID          uint   `gorm:"primaryKey;unique"               json:"id"`
	Name        string `binding:"required,min=2,max=50"        gorm:"not null" json:"name"`
	PhoneNumber string `binding:"required,min=10,max=10"       gorm:"not null" json:"phone_number"`
	House       string `binding:"required"                     gorm:"not null" json:"house"`
	Area        string `gorm:"not null"                        json:"area"`
	LandMark    string `binding:"required"                     gorm:"not null" json:"land_mark"`
	City        string `gorm:"not null"                        json:"city"`
	Pincode     uint   `binding:"required,numeric,min=6,max=6" gorm:"not null" json:"pincode"`
	CountryID   uint   `binding:"required"                     gorm:"not null" json:"country_id"`
	Country     Country
	CreatedAt   time.Time `gorm:"not null"   json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Country struct {
	ID          uint   `gorm:"primaryKey;unique;" json:"id"`
	CountryName string `gorm:"unique;not null"    json:"country_name"`
}
