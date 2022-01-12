package internal

import (
	"github.com/bradfitz/gomemcache/memcache"
	"log"
	"math/big"
	"os"
	"strconv"
)

type Memcached struct {
	mc *memcache.Client
}

var mc = Memcached{
	memcache.New("memcached:" + os.Getenv("MC_PORT")),
}

func (m *Memcached) CheckMemcached(from int, to int) []*big.Int {
	var cachTable []*big.Int
	cachTable = make([]*big.Int, to+1)

	for i := from - 1; i <= to; i += 1 {
		cach, err := m.mc.Get(strconv.Itoa(i))
		if err != nil {
			return GenerateNum(from, to)
		}
		cachTable[i] = new(big.Int).SetBytes(cach.Value)
	}
	return cachTable[from-1 : to]
}

func (m *Memcached) AddMemcached(table []*big.Int) {
	for key, value := range table {
		err := m.mc.Set(&memcache.Item{Key: strconv.Itoa(key), Value: value.Bytes()})
		if err != nil {
			log.Println(err)
		}
	}
}
