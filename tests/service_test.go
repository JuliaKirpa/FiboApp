package tests

import (
	"FiboApp/internal"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestGenerateNum(t *testing.T) {
	req := require.New(t)

	table := internal.GenerateNum(5, 10)
	exTable := []*big.Int{big.NewInt(3), big.NewInt(5), big.NewInt(8), big.NewInt(13), big.NewInt(21), big.NewInt(34)}

	req.Equal(exTable, table)
}
