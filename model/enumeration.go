package model

import "time"

type IrregularityLevel struct {
	Id        int       `json:"id" db:"id"`
	Code      string    `json:"code" db:"code"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
