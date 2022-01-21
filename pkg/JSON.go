package pkg

import "math/big"

type FiboElement struct {
	Number int
	Value  *big.Int
}

func JSON(table []*big.Int, from int) []FiboElement {
	var jsonData []FiboElement

	for key, value := range table {
		element := FiboElement{Number: key + from, Value: value}
		jsonData = append(jsonData, element)
	}

	return jsonData
}
