package internal

import (
	"github.com/bradfitz/gomemcache/memcache"
	"log"
	"math/big"
	"strconv"
)

func CheckMemcached(from, to int, mc *memcache.Client) []*big.Int {
	var cachTable []*big.Int
	cachTable = make([]*big.Int, to+1)

	for i := from - 1; i <= to; i += 1 {
		cach, err := mc.Get(strconv.Itoa(i))
		if err != nil {
			return GenerateNum(from, to, mc)
		}
		cachTable[i] = new(big.Int).SetBytes(cach.Value)
	}
	return cachTable[from-1 : to]
}

func AddMemcached(mc *memcache.Client, table []*big.Int) {
	for key, value := range table {
		err := mc.Set(&memcache.Item{Key: strconv.Itoa(key), Value: value.Bytes()})
		if err != nil {
			log.Println(err)
		}
	}
}
