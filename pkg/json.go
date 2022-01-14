package pkg

import (
	"math/big"
)

type FiboNum struct {
	Id    int      `json:"id"`
	Value *big.Int `json:"value"`
}

type ArrValues []FiboNum

func JSON(table []*big.Int, from int) ArrValues {
	var jsonData ArrValues

	for key, value := range table {
		num := FiboNum{key + from, value}
		jsonData = append(jsonData, num)
	}

	return jsonData
}
