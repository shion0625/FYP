package request

type Login struct {
	UserName string `binding:"omitempty,min=3,max=15"  json:"user_name"`
	Phone    string `binding:"omitempty,min=10,max=10" json:"phone"`
	Email    string `binding:"omitempty,email"         json:"email"`
	Password string `binding:"required,min=5,max=30"   json:"password"`
}

type RefreshToken struct {
	RefreshToken string `binding:"min=10" json:"refresh_token"`
}

type SignUp struct {
	UserName        string `binding:"required,min=3,max=15"            json:"user_name"`
	FirstName       string `binding:"required,min=2,max=50"            json:"first_name"`
	LastName        string `binding:"required,min=1,max=50"            json:"last_name"`
	Age             uint   `binding:"required,numeric"                 json:"age"`
	Email           string `binding:"required,email"                   json:"email"`
	Phone           string `binding:"required,min=10,max=10"           json:"phone"`
	Password        string `binding:"required,eqfield=ConfirmPassword" json:"password"`
	ConfirmPassword string `binding:"required"                         json:"confirm_password"`
}
