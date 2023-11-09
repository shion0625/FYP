package request

type Login struct {
	UserName string `binding:"omitempty,min=3,max=15"  json:"userName"`
	Phone    string `binding:"omitempty,min=10,max=10" json:"phone"`
	Email    string `binding:"omitempty,email"         json:"email"`
	Password string `binding:"required,min=5,max=30"   json:"password"`
}

type RefreshToken struct {
	RefreshToken string `binding:"min=10" json:"refreshToken"`
}

type SignUp struct {
	UserName        string `binding:"required,min=3,max=15"            json:"userName"`
	FirstName       string `binding:"required,min=2,max=50"            json:"firstName"`
	LastName        string `binding:"required,min=1,max=50"            json:"lastName"`
	Age             uint   `binding:"required,numeric"                 json:"age"`
	Email           string `binding:"required,email"                   json:"email"`
	Phone           string `binding:"required,min=10,max=10"           json:"phone"`
	Password        string `binding:"required,eqfield=ConfirmPassword" json:"password"`
	ConfirmPassword string `binding:"required"                         json:"confirmPassword"`
}
