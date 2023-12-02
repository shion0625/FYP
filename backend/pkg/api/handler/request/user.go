package request

// for address add address.
type Address struct {
	Name        string `json:"name"        validate:"required,min=2,max=50"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10,max=11"`
	House       string `json:"house"       validate:"required"`
	Area        string `json:"area"`
	LandMark    string `json:"landMark"    validate:"required"`
	City        string `json:"city"`
	Pincode     uint   `json:"pincode"     validate:"required"`
	// CountryID   uint   `json:"country_id" validate:"required"`

	IsDefault *bool `json:"isDefault"`
}
type EditAddress struct {
	ID          uint   `json:"id"          validate:"required"`
	Name        string `json:"name"        validate:"required,min=2,max=50"`
	PhoneNumber string `json:"phoneNumber" validate:"required,len=10"`
	House       string `json:"house"       validate:"required"`
	Area        string `json:"area"`
	LandMark    string `json:"landMark"    validate:"required"`
	City        string `json:"city"`
	Pincode     uint   `json:"pincode"     validate:"required"`
	// CountryID   uint   `json:"country_id" validate:"required"`

	IsDefault *bool `json:"isDefault"`
}

type EditUser struct {
	UserName        string `json:"userName"        validate:"required,min=3,max=15"`
	FirstName       string `json:"firstName"       validate:"required,min=2,max=50"`
	LastName        string `json:"lastName"        validate:"required,min=1,max=50"`
	Age             uint   `json:"age"             validate:"required,numeric"`
	Email           string `json:"email"           validate:"required,email"`
	Phone           string `json:"phone"           validate:"required,len=10"`
	Password        string `json:"password"        validate:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
}
