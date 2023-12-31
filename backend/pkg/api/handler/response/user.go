package response

import "time"

// user details response.
type User struct {
	ID          string    `copier:"must"      json:"id"`
	GoogleImage string    `json:"googleImage"`
	FirstName   string    `copier:"must"      json:"firstName"`
	LastName    string    `copier:"must"      json:"lastName"`
	Age         uint      `copier:"must"      json:"age"`
	Email       string    `copier:"must"      json:"email"`
	UserName    string    `copire:"must"      json:"userName"`
	Phone       string    `copier:"must"      json:"phone"`
	Verified    bool      `json:"verified"`
	BlockStatus bool      `copier:"must"      json:"blockStatus"`
	CreatedAt   time.Time `gorm:"not null"    json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// address.
type Address struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	House       string `json:"house"`
	Area        string `json:"area"`
	LandMark    string `json:"landMark"`
	City        string `json:"city"`
	Pincode     string `json:"pincode"`
	CountryName string `json:"countryName"`

	IsDefault *bool `json:"isDefault"`
}

type PaymentMethod struct {
	ID          uint   `gorm:"primaryKey;not null" json:"id"`
	Number      string `gorm:"unique;not null"     json:"number"`
	CardCompany string `gorm:"not null"            json:"cardCompany"`
}
