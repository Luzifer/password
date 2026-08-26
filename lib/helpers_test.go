package securepassword

import "testing"

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
