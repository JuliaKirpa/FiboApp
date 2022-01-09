package pkg

import (
	"math/big"
	"strconv"
	"strings"
)

func Validate(from, to int) string {
	switch {
	case from > to:
		return "The end value of the interval is bigger than the start value"
	case to == 0:
		return "The end value should be bigger than 0"
	case from > 2147483647:
		return "The biggest value for from is 2 147 483 647"
	case to > 2147483647:
		return "The biggest value for to is 2 147 483 647"
	}
	return ""
}

func Prepare(table []*big.Int, from int) string {
	var response strings.Builder
	for key, value := range table {
		response.WriteString(strconv.Itoa(from + 1 + key))
		response.WriteString(" ")
		response.WriteString(value.String())
		response.WriteString("\n")
	}
	return response.String()
}
