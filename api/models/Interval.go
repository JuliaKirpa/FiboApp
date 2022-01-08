package models

import "errors"

type Interval struct {
	From int `json:"from"`
	To   int `json:"to"`
}

func (i Interval) Validate(from, to int) error {
	switch {
	case from > to:
		return errors.New("The end value of the interval is bigger than the start value")
	case to == 0:
		return errors.New("The end value should be bigger than 0")
	case from < 0:
		return errors.New("Values coudn't be negative")
	case to < 0:
		return errors.New("Values coudn't be negative")
	}
	return nil
}
