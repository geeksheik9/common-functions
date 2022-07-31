package models

type User struct {
	ID        int64    `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	CreatedAt int64    `json:"createdAt"`
	UpdatedAt int64    `json:"updatedAt"`
	Token     string   `json:"token"`
	Roles     []string `json:"roles"`
}
