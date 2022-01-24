package tests

import (
	"FiboApp/pkg"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidate(t *testing.T) {
	req := require.New(t)
	from, to, err := pkg.Validate("17", "29")
	expFrom := 17
	expTo := 29

	req.Empty(err)
	req.NotEmpty(from)
	req.NotEmpty(to)
	req.Greater(to, from)
	req.Greater(from, 0)
	req.Greater(to, 0)
	req.Equal(expFrom, from)
	req.Equal(expTo, to)
}
