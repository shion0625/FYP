package request

// for address add address.
type Address struct {
	Name        string `json:"name"        validate:"required,min=2,max=100"`
	PhoneNumber string `json:"phoneNumber" validate:"required,e164"`
	House       string `json:"house"       validate:"required"`
	Area        string `json:"area"`
	LandMark    string `json:"landMark"    validate:"required"`
	City        string `json:"city"`
	Pincode     string `json:"pincode"     validate:"required"`
	CountryName string `json:"countryName" validate:"required"`
	IsDefault   *bool  `json:"isDefault"`
}

type EditAddress struct {
	ID          uint   `json:"id"          validate:"required"`
	Name        string `json:"name"        validate:"required,min=2,max=100"`
	PhoneNumber string `json:"phoneNumber" validate:"required,e164"`
	House       string `json:"house"       validate:"required"`
	Area        string `json:"area"`
	LandMark    string `json:"landMark"    validate:"required"`
	City        string `json:"city"`
	Pincode     string `json:"pincode"     validate:"required"`
	CountryName string `json:"countryName" validate:"required"`
	IsDefault   *bool  `json:"isDefault"`
}

type EditUser struct {
	UserName        string `json:"userName"        validate:"required,alphanum,min=3,max=15"`
	FirstName       string `json:"firstName"       validate:"required,alpha,min=2,max=50"`
	LastName        string `json:"lastName"        validate:"required,alpha,min=1,max=50"`
	Age             uint   `json:"age"             validate:"required,number,gte=0,lte=120"`
	Email           string `json:"email"           validate:"required,email"`
	Phone           string `json:"phone"           validate:"required,e164"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword" validate:"eqfield=Password"`
}

type PaymentMethod struct {
	Number string `gorm:"unique;not null" json:"number"`
	Name   string `gorm:"not null"        json:"name"`
	Expiry string `gorm:"not null"        json:"expiry"`
	Cvc    string `gorm:"not null"        json:"cvc"`
}

type UpdatePaymentMethod struct {
	ID     uint   `json:"id"              validate:"required"`
	Number string `gorm:"unique;not null" json:"number"       validate:"required"`
	Name   string `gorm:"not null"        json:"name"         validate:"required"`
	Expiry string `gorm:"not null"        json:"expiry"       validate:"required"`
	Cvc    string `gorm:"not null"        json:"cvc"          validate:"required"`
}
