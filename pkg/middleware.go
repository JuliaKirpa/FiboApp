package pkg

import "math/big"

func Validate(from, to int) string {
	switch {
	case from > to:
		return "The end value of the interval is bigger than the start value"
	case to == 0:
		return "The end value should be bigger than 0"
	case from < 0:
		return "Values coudn't be negative"
	case to < 0:
		return "Values coudn't be negative"
	}
	return ""
}

func Prepare(table []big.Int) string {
	return ""
}
