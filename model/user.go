package model

type User struct {
	Id        int    `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
	Password  string `json:"-" db:"password"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
}
