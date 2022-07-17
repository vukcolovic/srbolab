package model

import "time"

type User struct {
	Id              int       `json:"id" db:"id"`
	FirstName       string    `json:"first_name" db:"first_name"`
	LastName        string    `json:"last_name" db:"last_name"`
	Email           string    `json:"email" db:"email"`
	PhoneNumber     string    `json:"phone_number" db:"phone_number"`
	ContractNumber  string    `json:"contract_number" db:"contract_number"`
	ContractType    string    `json:"contract_type" db:"contract_type"`
	JMBG            string    `json:"jmbg" db:"jmbg"`
	Adress          string    `json:"adress" db:"adress"`
	StartedWork     time.Time `json:"started_work" db:"started_work"`
	Password        string    `json:"password" db:"password"`
	CurrentPassword string    `json:"current_password"`
	Deleted         bool      `json:"deleted" db:"deleted"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
