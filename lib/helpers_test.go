package securepassword

import (
	"errors"
	"testing"
)

func TestRandIntn(t *testing.T) {
	var (
		bound      = 16
		sampleSize = 200000
	)

	for range sampleSize {
		v, err := randIntn(bound)
		if err != nil {
			t.Fatalf("error in rng: %s", err)
		}

		if v < 0 || v >= bound {
			t.Errorf("rng yielded number out-of-range 0-%d: %d", bound, v)
		}
	}
}

func TestRandIntnRejectsNonPositive(t *testing.T) {
	if _, err := randIntn(0); !errors.Is(err, ErrInvalidRandBound) {
		t.Fatalf("randIntn(0): got %v, want ErrInvalidRandBound", err)
	}
	if _, err := randIntn(-5); !errors.Is(err, ErrInvalidRandBound) {
		t.Fatalf("randIntn(-5): got %v, want ErrInvalidRandBound", err)
	}
}
