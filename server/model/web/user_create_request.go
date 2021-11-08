package web

type UserCreateRequest struct {
	Username     string `validate:"required,min=1,max=30" json:"username"`
	Email        string `validate:"required,min=1,max=100" json:"email"`
	FirstName    string `validate:"required,min=1,max=100" json:"firstname"`
	LastName     string `validate:"max=100" json:"lastname"`
	Password     string `validate:"required,min=5,max=100" json:"password"`
	ProfileImage string `validate:"max=100" json:"profileImage"`
}
