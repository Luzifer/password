package securepassword

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestXKCDWordList(t *testing.T) {
	assert.GreaterOrEqual(t, len(xkcdWordList), 1000)
}

func TestXKCDGeneratePassword(t *testing.T) {
	for i := 4; i < 20; i++ {
		pwd, err := DefaultXKCD.GeneratePassword(i, false)
		require.NoError(t, err)

		assert.True(t, regexp.MustCompile(fmt.Sprintf("^([A-Z][a-z]+){%d}$", i)).MatchString(pwd))
	}
}

func TestXKCDDatePrepend(t *testing.T) {
	pwd, err := DefaultXKCD.GeneratePassword(4, true)
	require.NoError(t, err)

	assert.True(t, regexp.MustCompile(`^[0-9]{8}\.([A-Z][a-z]+){4}$`).MatchString(pwd))
}

func TestXKCDSeparator(t *testing.T) {
	gen := NewXKCDGenerator()
	gen.Separator = "-"

	pwd, err := gen.GeneratePassword(4, false)
	require.NoError(t, err)

	assert.True(t, regexp.MustCompile(`^(?:[A-Z][a-z]+-){3}(?:[A-Z][a-z]+)$`).MatchString(pwd))
}

func BenchmarkGeneratePasswords4Words(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		_, err = DefaultXKCD.GeneratePassword(4, false)
		require.NoError(b, err)
	}
}

func BenchmarkGeneratePasswords20Words(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		_, err = DefaultXKCD.GeneratePassword(20, false)
		require.NoError(b, err)
	}
}
