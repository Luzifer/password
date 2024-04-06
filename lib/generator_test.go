package securepassword

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInsecurePasswords(t *testing.T) {
	passwords := map[string]string{
		`8452028337962356`: "password with only numeric characters was accepted",
		`adfgjadrgdagasdf`: "password with only lowercase characters was accepted",
		`ASEFSTDHQAEGFADF`: "password with only uppercase characters was accepted",
		`135fach74nc94bd6`: "password without uppercase characters was accepted",
		`235JGOA0YTVKS46S`: "password without lowercase characters was accepted",
		`sdgAFfgADTSgafoa`: "password without numeric characters was accepte",

		`cKTn5mQXfasdS6qy`: "password with pattern asd was accepted",
		`cKTn5mQXfdsaS6qy`: "password with pattern dsa was accepted",
		`cKTn5mQXf345S6qy`: "password with pattern 345 was accepted",
		`cKTn5mQXf987S6qy`: "password with pattern 987 was accepted",
		`cKTn5mQXfabcS6qy`: "password with pattern abc was accepted",
		`cKTn5mQXfcbaS6qy`: "password with pattern cba was accepted",
		`cKTn5mQXfABCS6qy`: "password with pattern ABC was accepted",
		`cKTn5mQXfONMS6qy`: "password with pattern ONM was accepted",

		`Gncj5zzK29Dvx92h`: "password with character repetition was accepted",
		`Gncj5%%K29Dvx92h`: "password with character repetition was accepted",
		`Gncj55%K29Dvx92h`: "password with character repetition was accepted",
	}

	sp := NewSecurePassword()
	for password, errorMessage := range passwords {
		assert.False(t, sp.CheckPasswordSecurity(password, false), errorMessage)
	}
}

func TestSecurePasswords(t *testing.T) {
	passwords := []string{
		`6e1GZ6V2empWAky5Z13a`,
		`DLHZA2zfWor1XUoJYvFR`,
		`sMf3uNf2E1pxPFMymah5`,
		`prb4tX1vtyVL7R316dKU`,
		`7bWc9C1ciL62h5u26Z9g`,
	}

	sp := NewSecurePassword()
	for _, password := range passwords {
		assert.True(t, sp.CheckPasswordSecurity(password, false), "password was rejected: %s", password)
	}
}

func TestPasswordWithoutSpecialCharaterFail(t *testing.T) {
	passwords := []string{
		`6e1GZ6V2empWAky5Z13a`,
		`DLHZA2zfWor1XUoJYvFR`,
		`sMf3uNf2E1pxPFMymah5`,
		`prb4tX1vtyVL7R316dKU`,
		`7bWc9C1ciL62h5u26Z9g`,
	}

	sp := NewSecurePassword()
	for _, password := range passwords {
		assert.False(t, sp.CheckPasswordSecurity(password, true), "password was accepted: %s", password)
	}
}

func TestSecurePasswordWithSpecialCharacter(t *testing.T) {
	passwords := []string{
		`a*5S(+zQ9=<J~bH}!F])`,
		`9?)k1Ge?[F}5w{#$Ho(6`,
		`LRrot)%qVH!>3%/1MiNr`,
		`/K}.]C4-a/{,39r$"(D+`,
		`9dk#:@xjPd_m$:F"}>Cj`,
	}

	sp := NewSecurePassword()
	for _, password := range passwords {
		assert.True(t, sp.CheckPasswordSecurity(password, true), "password was rejected: %s", password)
	}
}

func TestPasswordGeneration(t *testing.T) {
	sp := NewSecurePassword()
	password, err := sp.GeneratePassword(20, false)
	require.NoError(t, err)

	assert.Len(t, password, 20)
	assert.True(t, sp.CheckPasswordSecurity(password, false))
	assert.False(t, sp.CheckPasswordSecurity(password, true))

	password, err = NewSecurePassword().GeneratePassword(32, true)
	require.NoError(t, err)

	assert.Len(t, password, 32)
	assert.True(t, sp.CheckPasswordSecurity(password, false))
	assert.True(t, sp.CheckPasswordSecurity(password, true))
}

func TestImpossiblePasswords(t *testing.T) {
	sp := NewSecurePassword()
	for i := 0; i < 4; i++ {
		_, err := sp.GeneratePassword(i, false)
		assert.ErrorIs(t, err, ErrLengthTooLow)
	}
}

func TestBadCharacters(t *testing.T) {
	badCharacters := []string{"I", "l", "0", "O", "B", "8"}
	sp := NewSecurePassword()

	for i := 0; i < 500; i++ {
		pwd, err := sp.GeneratePassword(20, false)
		require.NoError(t, err)

		for _, char := range badCharacters {
			assert.NotContains(t, pwd, char)
		}
	}
}

func BenchmarkGeneratePasswords8Char(b *testing.B) {
	pwd := NewSecurePassword()
	var err error

	for i := 0; i < b.N; i++ {
		_, err = pwd.GeneratePassword(8, false)
		require.NoError(b, err)
	}
}

func BenchmarkGeneratePasswords8CharSpecial(b *testing.B) {
	pwd := NewSecurePassword()
	var err error

	for i := 0; i < b.N; i++ {
		_, err = pwd.GeneratePassword(8, true)
		require.NoError(b, err)
	}
}

func BenchmarkGeneratePasswords16Char(b *testing.B) {
	pwd := NewSecurePassword()
	var err error

	for i := 0; i < b.N; i++ {
		_, err = pwd.GeneratePassword(16, false)
		require.NoError(b, err)
	}
}

func BenchmarkGeneratePasswords16CharSpecial(b *testing.B) {
	pwd := NewSecurePassword()
	var err error

	for i := 0; i < b.N; i++ {
		_, err = pwd.GeneratePassword(16, true)
		require.NoError(b, err)
	}
}

func BenchmarkGeneratePasswords32Char(b *testing.B) {
	pwd := NewSecurePassword()
	var err error

	for i := 0; i < b.N; i++ {
		_, err = pwd.GeneratePassword(32, false)
		require.NoError(b, err)
	}
}

func BenchmarkGeneratePasswords32CharSpecial(b *testing.B) {
	pwd := NewSecurePassword()
	var err error

	for i := 0; i < b.N; i++ {
		_, err = pwd.GeneratePassword(32, true)
		require.NoError(b, err)
	}
}

func BenchmarkGeneratePasswords128Char(b *testing.B) {
	pwd := NewSecurePassword()
	var err error

	for i := 0; i < b.N; i++ {
		_, err = pwd.GeneratePassword(128, false)
		require.NoError(b, err)
	}
}

func BenchmarkGeneratePasswords128CharSpecial(b *testing.B) {
	pwd := NewSecurePassword()
	var err error

	for i := 0; i < b.N; i++ {
		_, err = pwd.GeneratePassword(128, true)
		require.NoError(b, err)
	}
}
