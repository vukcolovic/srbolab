package model

import (
	"database/sql"
	"time"
)

//json struct
type Irregularity struct {
	Id            int               `json:"id"`
	Subject       string            `json:"subject"`
	Level         IrregularityLevel `json:"irregularity_level"`
	Controller    User              `json:"controller"`
	CreatedBy     User              `json:"created_by"`
	Description   string            `json:"description"`
	Notice        string            `json:"notice"`
	Corrected     bool              `json:"corrected"`
	CorrectedBy   User              `json:"corrected_by"`
	CorrectedDate time.Time         `json:"corrected_date"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
}

//db struct

type IrregularityDb struct {
	Id            int           `db:"id"`
	Subject       string        `db:"subject"`
	Level         int           `db:"level_id"`
	Controller    sql.NullInt64 `db:"controller_id"`
	CreatedBy     sql.NullInt64 `db:"created_by"`
	Description   string        `db:"description"`
	Notice        string        `db:"notice"`
	Corrected     bool          `db:"corrected"`
	CorrectedBy   sql.NullInt64 `db:"corrected_by"`
	CorrectedDate time.Time     `db:"corrected_date"`
	CreatedAt     time.Time     `db:"created_at"`
	UpdatedAt     time.Time     `db:"updated_at"`
}
