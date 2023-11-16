package request

// for address add address.
type Address struct {
	Name        string `binding:"required,min=2,max=50"  json:"name"`
	PhoneNumber string `binding:"required,min=10,max=10" json:"phoneNumber"`
	House       string `Binding:"required"               json:"house"`
	Area        string `json:"area"`
	LandMark    string `binding:"required"               json:"landMark"`
	City        string `json:"city"`
	Pincode     uint   `binding:"required"               json:"pincode"`
	// CountryID   uint   `json:"country_id" binding:"required"`

	IsDefault *bool `json:"isDefault"`
}
type EditAddress struct {
	ID          uint   `binding:"required"               json:"id"`
	Name        string `binding:"required,min=2,max=50"  json:"name"`
	PhoneNumber string `binding:"required,min=10,max=10" json:"phoneNumber"`
	House       string `binding:"required"               json:"house"`
	Area        string `json:"area"`
	LandMark    string `binding:"required"               json:"landMark"`
	City        string `json:"city"`
	Pincode     uint   `binding:"required"               json:"pincode"`
	// CountryID   uint   `json:"country_id" binding:"required"`

	IsDefault *bool `json:"isDefault"`
}

type EditUser struct {
	UserName        string `binding:"required,min=3,max=15"             json:"userName"`
	FirstName       string `binding:"required,min=2,max=50"             json:"firstName"`
	LastName        string `binding:"required,min=1,max=50"             json:"lastName"`
	Age             uint   `binding:"required,numeric"                  json:"age"`
	Email           string `binding:"required,email"                    json:"email"`
	Phone           string `binding:"required,min=10,max=10"            json:"phone"`
	Password        string `binding:"omitempty,eqfield=ConfirmPassword" json:"password"`
	ConfirmPassword string `binding:"omitempty"                         json:"confirmPassword"`
}
