package web

type UserUpdateRequest struct {
	ID           int    `validate:"required" json:"id"`
	Username     string `validate:"required,min=1,max=30" json:"username"`
	Email        string `validate:"required,min=1,max=100" json:"email"`
	FirstName    string `validate:"required,min=1,max=100" json:"firstname"`
	LastName     string `validate:"required,min=1,max=100" json:"lastname"`
	ProfileImage string `validate:"min=1,max=100" json:"profileImage"`
}
