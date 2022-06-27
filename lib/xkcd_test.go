package securepassword

import (
	"fmt"
	"regexp"
	"testing"
)

func TestXKCDWordList(t *testing.T) {
	if w := len(xkcdWordList); w < 1000 {
		t.Fatalf("Word list is expected to contain at least 1000 words, has %d", w)
	}
}

func TestXKCDGeneratePassword(t *testing.T) {
	for i := 4; i < 20; i++ {
		pwd, err := DefaultXKCD.GeneratePassword(i, false)
		if err != nil {
			t.Fatalf("Generated had an error: %s", err)
		}

		if !regexp.MustCompile(fmt.Sprintf("^([A-Z][a-z]+){%d}$", i)).MatchString(pwd) {
			t.Errorf("Password %q is expected to contain %d words, did not match expected RegEx", pwd, i)
		}
	}
}

func TestXKCDDatePrepend(t *testing.T) {
	pwd, err := DefaultXKCD.GeneratePassword(4, true)
	if err != nil {
		t.Fatalf("Generated had an error: %s", err)
	}

	if !regexp.MustCompile(`^[0-9]{8}\.([A-Z][a-z]+){4}$`).MatchString(pwd) {
		t.Errorf("Password %q did not match expected RegEx", pwd)
	}
}

func TestXKCDSeparator(t *testing.T) {
	gen := NewXKCDGenerator()
	gen.Separator = "-"

	pwd, err := gen.GeneratePassword(4, false)
	if err != nil {
		t.Fatalf("Generated had an error: %s", err)
	}

	if !regexp.MustCompile(`^(?:[A-Z][a-z]+-){3}(?:[A-Z][a-z]+)$`).MatchString(pwd) {
		t.Errorf("Password %q did not match expected RegEx", pwd)
	}
}

func BenchmarkGeneratePasswords4Words(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DefaultXKCD.GeneratePassword(4, false)
	}
}

func BenchmarkGeneratePasswords20Words(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DefaultXKCD.GeneratePassword(20, false)
	}
}
