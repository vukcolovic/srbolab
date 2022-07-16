package model

import (
	"database/sql"
	"srbolabApp/util"
	"time"
)

type FuelConsumption struct {
	Id              int       `json:"id"`
	DateConsumption util.Date `json:"date_consumption"`
	FuelType        string    `json:"fuel_type"`
	Liter           float64   `json:"liter"`
	Price           float64   `json:"price"`
	CarRegistration string    `json:"car_registration"`
	PouredBy        User      `json:"poured_by"`
	CreatedBy       User      `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type FuelConsumptionDb struct {
	Id              int           `db:"id"`
	DateConsumption time.Time     `db:"date_consumption"`
	FuelType        string        `db:"fuel_type"`
	Liter           float64       `db:"liter"`
	Price           float64       `db:"price"`
	CarRegistration string        `db:"car_registration"`
	PouredBy        sql.NullInt64 `db:"poured_by"`
	CreatedBy       sql.NullInt64 `db:"created_by"`
	CreatedAt       time.Time     `db:"created_at"`
	UpdatedAt       time.Time     `db:"updated_at"`
}
