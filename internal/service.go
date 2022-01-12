package internal

import (
	"math/big"
)

func GenerateNum(from, to int) []*big.Int {
	var table []*big.Int
	table = make([]*big.Int, to+1)

	table[0] = new(big.Int).SetInt64(0)
	table[1] = new(big.Int).SetInt64(1)

	for i := 2; i <= to; i += 1 {
		table[i] = new(big.Int).Add(table[i-1], table[i-2])
	}
	go mc.AddMemcached(table)
	return table[from-1 : to]
}
