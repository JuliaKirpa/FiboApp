package internal

import (
	"math/big"
)

func GenerateNum(from, to int) []*big.Int {
	var table []*big.Int
	//var cachTable []*big.Int
	//
	//cachTable = make([]*big.Int, to+1)
	//
	//for i := from - 1; i <= to; i += 1 {
	//	cach, err := mc.Get(strconv.Itoa(i))
	//	if err != nil {
	//		break
	//	}
	//	cachTable[i] = new(big.Int).SetBytes(cach.Value)
	//}
	//errCach := mc.Touch(strconv.Itoa(to), 0)
	//if errCach == nil {
	//	return cachTable[from-1 : to]
	//}

	table = make([]*big.Int, to+1)

	table[0] = new(big.Int).SetInt64(0)
	//err1 := mc.Set(&memcache.Item{Key: "0", Value: table[0].Bytes()})
	//if err1 != nil {
	//	log.Println(err1)
	//}
	table[1] = new(big.Int).SetInt64(1)
	//err2 := mc.Set(&memcache.Item{Key: "1", Value: table[1].Bytes()})
	//if err2 != nil {
	//	log.Println(err2)
	//}

	for i := 2; i <= to; i += 1 {
		table[i] = new(big.Int).Add(table[i-1], table[i-2])
		//err := mc.Set(&memcache.Item{Key: strconv.Itoa(i), Value: table[i].Bytes()})
		//if err != nil {
		//	log.Println(err)
		//}
	}
	return table[from-1 : to]
}
