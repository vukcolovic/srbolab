package model

type IrregularityLevel struct {
	Id        int    `json:"id" db:"id"`
	Code      string `json:"code" db:"code"`
	CreatedAt string `json:"created_at" db:"created_at"`
}
