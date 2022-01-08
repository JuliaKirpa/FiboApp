package internal

import (
	"github.com/bradfitz/gomemcache/memcache"
	"math/big"
	"strconv"
)

func GenerateNum(from, to int) []*big.Int {
	var table []*big.Int
	mc := memcache.New(":11211")

	table = make([]*big.Int, to+1)

	table[0] = new(big.Int).SetInt64(0)
	mc.Set(&memcache.Item{Key: "0", Value: []byte("0")})
	table[1] = new(big.Int).SetInt64(1)
	//	mc.Set(&memcache.Item{Key: "1", Value: []byte("1")})

	for i := 2; i <= to; i += 1 {
		table[i] = new(big.Int).Add(table[i-1], table[i-2])
		mc.Set(&memcache.Item{Key: strconv.Itoa(i), Value: table[i].Bytes()})
	}

	return table[from : to+1]
}
