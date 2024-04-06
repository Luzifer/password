package securepassword

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func randIntn(max int) (int, error) {
	cidx, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, fmt.Errorf("generating random number: %w", err)
	}

	return int(cidx.Int64()), nil
}
