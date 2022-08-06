package util

import (
	"encoding/json"
	"fmt"
	"time"
)

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
		t, err = time.Parse(time.RFC3339Nano, s)
		if err != nil {
			return fmt.Errorf("invalid date: %v", err)
		}
	}
	d.Time = t
	return nil
}
