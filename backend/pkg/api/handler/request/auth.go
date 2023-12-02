package request

type Login struct {
	UserName string `json:"userName" validate:"alphanum,min=3,max=15"`
	Phone    string `json:"phone"    validate:"e164"`
	Email    string `json:"email"    validate:"email"`
	Password string `json:"password" validate:"required,min=5,max=30"`
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken" validate:"min=10"`
}

type SignUp struct {
	UserName        string `json:"userName"        validate:"required,alphanum,min=3,max=15"`
	FirstName       string `json:"firstName"       validate:"required,alpha,min=2,max=50"`
	LastName        string `json:"lastName"        validate:"required,alpha,min=1,max=50"`
	Age             uint   `json:"age"             validate:"required,number,gte=0,lte=120"`
	Email           string `json:"email"           validate:"required,email"`
	Phone           string `json:"phone"           validate:"required,e164"`
	Password        string `json:"password"        validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}
