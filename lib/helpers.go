package securepassword

import (
	"crypto/rand"
	"math/big"

	"github.com/pkg/errors"
)

func randIntn(max int) (int, error) {
	cidx, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, errors.Wrap(err, "generating random number")
	}

	return int(cidx.Int64()), nil
}
