package internal

import (
	"github.com/bradfitz/gomemcache/memcache"
	"log"
	"math/big"
	"strconv"
)

func GenerateNum(from, to int, mc *memcache.Client) []*big.Int {
	var table []*big.Int
	var cachTable []*big.Int

	cachTable = make([]*big.Int, to+1)

	for i := from; i <= to; i += 1 {
		cach, err := mc.Get(strconv.Itoa(i))
		if err != nil {
			break
		}
		cachTable[i] = new(big.Int).SetBytes(cach.Value)
	}
	errCach := mc.Touch(strconv.Itoa(to), 0)
	if errCach == nil {
		return cachTable[from : to+1]
	}

	table = make([]*big.Int, to+1)

	table[0] = new(big.Int).SetInt64(0)
	err := mc.Set(&memcache.Item{Key: "0", Value: []byte("0")})
	if err != nil {
		log.Fatal(err)
	}
	table[1] = new(big.Int).SetInt64(1)
	mc.Set(&memcache.Item{Key: "1", Value: []byte("1")})

	for i := 2; i <= to; i += 1 {
		table[i] = new(big.Int).Add(table[i-1], table[i-2])
		mc.Set(&memcache.Item{Key: strconv.Itoa(i), Value: table[i].Bytes()})
	}
	return table[from : to+1]
}
