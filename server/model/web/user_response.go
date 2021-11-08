package web

type UserResponse struct {
	ID           int32  `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	ProfileImage string `json:"profileImage"`
}
