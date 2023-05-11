package user

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}
