package model

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

//filter irregularity
type IrregularityFilter struct {
	Subject    string             `json:"subject"`
	Level      *IrregularityLevel `json:"irregularity_level"`
	Controller *User              `json:"controller"`
	Checked    string             `json:"checked"`
	DateFrom   Date               `json:"date_from"`
	DateTo     Date               `json:"date_to"`
}

type Date struct{ time.Time }

func (d *Date) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("date should be a string, got %s", data)
	}
	if s == "" {
		return nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("invalid date: %v", err)
	}
	d.Time = t
	return nil
}
