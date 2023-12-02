package request

type Login struct {
	UserName string `json:"userName" validate:"omitempty,min=3,max=15"`
	Phone    string `json:"phone"    validate:"omitempty,min=10,max=10"`
	Email    string `json:"email"    validate:"omitempty,email"`
	Password string `json:"password" validate:"required,min=5,max=30"`
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken" validate:"min=10"`
}

type SignUp struct {
	UserName        string `json:"userName"        validate:"required,min=3,max=15"`
	FirstName       string `json:"firstName"       validate:"required,min=2,max=50"`
	LastName        string `json:"lastName"        validate:"required,min=1,max=50"`
	Age             uint   `json:"age"             validate:"required,numeric"`
	Email           string `json:"email"           validate:"required,email"`
	Phone           string `json:"phone"           validate:"required,min=10,max=10"`
	Password        string `json:"password"        validate:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
}
