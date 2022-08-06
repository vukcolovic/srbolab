package model

import (
	"database/sql"
	"srbolabApp/util"
	"time"
)

type User struct {
	Id              int       `json:"id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	PhoneNumber     string    `json:"phone_number"`
	ContractNumber  string    `json:"contract_number"`
	ContractType    string    `json:"contract_type"`
	JMBG            string    `json:"jmbg"`
	Adress          string    `json:"adress"`
	StartedWork     util.Date `json:"started_work"`
	Password        string    `json:"password"`
	CurrentPassword string    `json:"current_password"`
	Deleted         bool      `json:"deleted"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type UserDb struct {
	Id              int            `db:"id"`
	FirstName       string         `db:"first_name"`
	LastName        string         `db:"last_name"`
	Email           string         `db:"email"`
	PhoneNumber     sql.NullString `db:"phone_number"`
	ContractNumber  sql.NullString `db:"contract_number"`
	ContractType    sql.NullString `db:"contract_type"`
	JMBG            sql.NullString `db:"jmbg"`
	Adress          sql.NullString `db:"adress"`
	StartedWork     sql.NullTime   `db:"started_work"`
	Password        string         `db:"password"`
	CurrentPassword string         `db:"current_password"`
	Deleted         bool           `db:"deleted"`
	CreatedAt       time.Time      `db:"created_at"`
	UpdatedAt       time.Time      `db:"updated_at"`
}
