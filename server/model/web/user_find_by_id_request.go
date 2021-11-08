package web

type UserFindByID struct {
	ID int `validate:"required" json:"id"`
}
