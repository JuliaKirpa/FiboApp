package tests

import (
	"FiboApp/pkg"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestJSON(t *testing.T) {
	req := require.New(t)

	table := []*big.Int{big.NewInt(3), big.NewInt(5), big.NewInt(8), big.NewInt(13), big.NewInt(21), big.NewInt(34)}
	jsonData := pkg.JSON(table, 5)

	exJsonData := []pkg.FiboElement{
		{
			5,
			big.NewInt(3),
		},
		{
			6,
			big.NewInt(5),
		},
		{
			7,
			big.NewInt(8),
		},
		{
			8,
			big.NewInt(13),
		},
		{
			9,
			big.NewInt(21),
		},
		{
			10,
			big.NewInt(34),
		},
	}
	req.Equal(exJsonData, jsonData)
}
