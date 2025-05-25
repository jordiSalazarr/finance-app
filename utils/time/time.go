package time_utils

import (
	"fmt"
	"time"
)

type DateOnly struct {
	time.Time
}

const layout = "2006-01-02"

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	str := string(b)
	str = str[1 : len(str)-1] // eliminar comillas del JSON

	t, err := time.Parse(layout, str)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", d.Format(layout))), nil
}

func (d *DateOnly) UnmarshalText(text []byte) error {
	t, err := time.Parse(layout, string(text))
	if err != nil {
		return fmt.Errorf("invalid date format: %w", err)
	}
	d.Time = t
	return nil
}
