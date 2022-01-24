package pkg

import (
	"strconv"
)

func Validate(start, end string) (from int, to int, err string) {
	from, er := strconv.Atoi(start)
	if er != nil {
		return 0, 0, "Use positive integer value for FROM\n"
	}
	to, er1 := strconv.Atoi(end)
	if er1 != nil {
		return 0, 0, "Use positive integer value for TO\n"
	}
	switch {
	case from == 0:
		return 0, 0, "The start value should be greater than 0\n"
	case to == 0:
		return 0, 0, "The end value should be greater than 0\n"
	case from < 0:
		return 0, 0, "The start value should be greater than 0\n"
	case to < 0:
		return 0, 0, "The end value should be greater than 0\n"
	case from > to:
		return 0, 0, "The end value of the interval is greater than the start value\n"
	case from > 2147483647:
		return 0, 0, "The biggest value for from is 2 147 483 647\n"
	case to > 2147483647:
		return 0, 0, "The biggest value for to is 2 147 483 647\n"
	}
	return from, to, ""
}
