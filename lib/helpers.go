package securepassword

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

// ErrInvalidRandBound is returned when randIntn is called with a non-positive bound.
var ErrInvalidRandBound = errors.New("securepassword: rand bound must be positive")

func randIntn(maxN int) (int, error) {
	if maxN <= 0 {
		return 0, ErrInvalidRandBound
	}

	cidx, err := rand.Int(rand.Reader, big.NewInt(int64(maxN)))
	if err != nil {
		return 0, fmt.Errorf("generating random number: %w", err)
	}

	return int(cidx.Int64()), nil
}
