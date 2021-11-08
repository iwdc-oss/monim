package web

type UserFindByUsernameAndPasswordRequest struct {
	Username string `validate:"required,min=1,max=30" json:"username"`
	Password string `validate:"required,min=5,max=100" json:"password"`
}
