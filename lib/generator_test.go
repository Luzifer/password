package securepassword

import (
	"strings"
	"testing"
)

func TestInsecurePasswords(t *testing.T) {
	passwords := map[string]string{
		`8452028337962356`: "Password with only numeric characters was accepted.",
		`adfgjadrgdagasdf`: "Password with only lowercase characters was accepted.",
		`ASEFSTDHQAEGFADF`: "Password with only uppercase characters was accepted.",
		`135fach74nc94bd6`: "Password without uppercase characters was accepted.",
		`235JGOA0YTVKS46S`: "Password without lowercase characters was accepted.",
		`sdgAFfgADTSgafoa`: "Password without numeric characters was accepted",

		`cKTn5mQXfasdS6qy`: "Password with pattern asd was accepted.",
		`cKTn5mQXfdsaS6qy`: "Password with pattern dsa was accepted.",
		`cKTn5mQXf345S6qy`: "Password with pattern 345 was accepted.",
		`cKTn5mQXf987S6qy`: "Password with pattern 987 was accepted.",
		`cKTn5mQXfabcS6qy`: "Password with pattern abc was accepted.",
		`cKTn5mQXfcbaS6qy`: "Password with pattern cba was accepted.",
		`cKTn5mQXfABCS6qy`: "Password with pattern ABC was accepted.",
		`cKTn5mQXfONMS6qy`: "Password with pattern ONM was accepted.",

		`Gncj5zzK29Dvx92h`: "Password with character repetition was accepted",
		`Gncj5%%K29Dvx92h`: "Password with character repetition was accepted",
		`Gncj55%K29Dvx92h`: "Password with character repetition was accepted",
	}

	for password, errorMessage := range passwords {
		if NewSecurePassword().CheckPasswordSecurity(password, false) {
			t.Error(errorMessage)
		}
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

	for _, password := range passwords {
		if !NewSecurePassword().CheckPasswordSecurity(password, false) {
			t.Errorf("Password was rejected: %s", password)
		}
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

	for _, password := range passwords {
		if NewSecurePassword().CheckPasswordSecurity(password, true) {
			t.Errorf("Password was accepted: %s", password)
		}
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

	for _, password := range passwords {
		if !NewSecurePassword().CheckPasswordSecurity(password, true) {
			t.Errorf("Password was rejected: %s", password)
		}
	}
}

func TestPasswordGeneration(t *testing.T) {
	password, _ := NewSecurePassword().GeneratePassword(20, false)

	if len(password) != 20 {
		t.Error("Password did not match requested length")
	}

	if !NewSecurePassword().CheckPasswordSecurity(password, false) {
		t.Error("Password did not pass security test")
	}

	if NewSecurePassword().CheckPasswordSecurity(password, true) {
		t.Error("Password without special characters passed security test for passwords with special characters")
	}

	password, _ = NewSecurePassword().GeneratePassword(32, true)

	if len(password) != 32 {
		t.Error("Password did not match requested length")
	}

	if !NewSecurePassword().CheckPasswordSecurity(password, false) {
		t.Error("Password did not pass security test for passwords without special characters")
	}

	if !NewSecurePassword().CheckPasswordSecurity(password, true) {
		t.Error("Password without special characters did not security test for passwords with special characters")
	}
}

func TestImpossiblePasswords(t *testing.T) {
	for i := 0; i < 4; i++ {
		_, err := NewSecurePassword().GeneratePassword(i, false)
		if err != ErrLengthTooLow {
			t.Errorf("Password with a length of %d did not throw as ErrLengthTooLow error", i)
		}
	}
}

func TestBadCharacters(t *testing.T) {
	badCharacters := []string{"I", "l", "0", "O", "B", "8"}

	for i := 0; i < 500; i++ {
		pwd, err := NewSecurePassword().GeneratePassword(20, false)
		if err != nil {
			t.Errorf("An error occured: %s", err)
		}
		for _, char := range badCharacters {
			if strings.Contains(pwd, char) {
				t.Errorf("Password '%s' contained blacklisted character: '%s'", pwd, char)
				return
			}
		}
	}
}

func BenchmarkGeneratePasswords8Char(b *testing.B) {
	pwd := NewSecurePassword()
	for i := 0; i < b.N; i++ {
		pwd.GeneratePassword(8, false)
	}
}

func BenchmarkGeneratePasswords8CharSpecial(b *testing.B) {
	pwd := NewSecurePassword()
	for i := 0; i < b.N; i++ {
		pwd.GeneratePassword(8, true)
	}
}

func BenchmarkGeneratePasswords16Char(b *testing.B) {
	pwd := NewSecurePassword()
	for i := 0; i < b.N; i++ {
		pwd.GeneratePassword(16, false)
	}
}

func BenchmarkGeneratePasswords16CharSpecial(b *testing.B) {
	pwd := NewSecurePassword()
	for i := 0; i < b.N; i++ {
		pwd.GeneratePassword(16, true)
	}
}

func BenchmarkGeneratePasswords32Char(b *testing.B) {
	pwd := NewSecurePassword()
	for i := 0; i < b.N; i++ {
		pwd.GeneratePassword(32, false)
	}
}

func BenchmarkGeneratePasswords32CharSpecial(b *testing.B) {
	pwd := NewSecurePassword()
	for i := 0; i < b.N; i++ {
		pwd.GeneratePassword(32, true)
	}
}

func BenchmarkGeneratePasswords128Char(b *testing.B) {
	pwd := NewSecurePassword()
	for i := 0; i < b.N; i++ {
		pwd.GeneratePassword(128, false)
	}
}

func BenchmarkGeneratePasswords128CharSpecial(b *testing.B) {
	pwd := NewSecurePassword()
	for i := 0; i < b.N; i++ {
		pwd.GeneratePassword(128, true)
	}
}
