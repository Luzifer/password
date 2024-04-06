// Package securepassword implements a password generator and check.
package securepassword

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const minPasswordLength = 4

// SecurePassword provides methods for generating secure passwords and
// checking the security requirements of passwords
type SecurePassword struct {
	characterTables map[string]string
	insecurePattern []string
	badCharacters   []string
}

// ErrLengthTooLow represents an error thrown if the password will
// never be able match the security considerations in this package
var ErrLengthTooLow = errors.New("passwords with a length lower than 4 will never meet the security requirements")

// NewSecurePassword initializes a new SecurePassword generator
func NewSecurePassword() *SecurePassword {
	return &SecurePassword{
		characterTables: map[string]string{
			"numeric": "0123456789",
			"simple":  "abcdefghijklmnopqrstuvwxyz",
			"special": "!#$%&()*+,-_./:;=?@[]^{}~|",
		},
		insecurePattern: []string{
			"abcdefghijklmnopqrstuvwxyz",      // Alphabet
			"zyxwvutsrqponmlkjihgfedcba",      // Alphabet reversed
			"01234567890",                     // Numeric increasing
			"09876543210",                     // Numeric decreasing
			"qwertzuiopasdfghjklyxcvbnm",      // German keyboard layout
			"mnbvcxylkjhgfdsapoiuztrewq",      // German keyboard layout reversed
			"qwertyuiopasdfghjklzxcvbnm",      // US keyboard layout
			"mnbvcxzlkjhgfdsapoiuytrewq",      // US keyboard layour reversed
			"789_456_123_147_258_369_159_753", // Numpad patterns
		},
		badCharacters: []string{"I", "l", "0", "O", "B", "8"}, // Characters that could lead to confusion due to font
	}
}

// GeneratePassword generates a new password with a given length and
// optional special characters in it. The password is automatically
// checked against CheckPasswordSecurity in order to only deliver secure
// passwords.
//
//revive:disable-next-line:flag-parameter
func (s *SecurePassword) GeneratePassword(length int, special bool) (string, error) {
	// Sanity check
	if length < minPasswordLength {
		return "", ErrLengthTooLow
	}

	characterTable := strings.Join([]string{
		s.characterTables["simple"],
		strings.ToUpper(s.characterTables["simple"]),
		s.characterTables["numeric"],
	}, "")
	if special {
		characterTable = strings.Join([]string{characterTable, s.characterTables["special"]}, "")
	}

	password := ""
	for {
		cidx, err := randIntn(len(characterTable))
		if err != nil {
			return "", fmt.Errorf("generating random number: %w", err)
		}

		char := string(characterTable[cidx])
		if strings.Contains(strings.Join(s.badCharacters, ""), char) {
			continue
		}

		password = fmt.Sprintf("%s%s",
			password,
			char,
		)
		if len(password) == length {
			if s.CheckPasswordSecurity(password, special) {
				break
			}
			password = ""
		}
	}
	return password, nil
}

// CheckPasswordSecurity executes three checks to ensure the passwords
// meet the security considerations in this package:
//
// 1. The password may not contain pattern found on the keyboard or in alphabet
// 2. The password must have 3 or 4 different character groups in it
// 3. The password may not have repeating characters
func (s *SecurePassword) CheckPasswordSecurity(password string, needsSpecialCharacters bool) bool {
	return !s.hasInsecurePattern(password) &&
		s.matchesBasicSecurity(password, needsSpecialCharacters) &&
		!s.hasCharacterRepetition(password)
}

func (s *SecurePassword) hasInsecurePattern(password string) bool {
	for i := 0; i < len(password)-3; i++ {
		slice := password[i : i+3] // Extract an 3 char slice to check
		for _, pattern := range s.insecurePattern {
			if strings.Contains(pattern, slice) {
				return true
			}
			if strings.Contains(strings.ToUpper(pattern), slice) {
				return true
			}
		}
	}

	return false
}

//revive:disable-next-line:flag-parameter
func (*SecurePassword) matchesBasicSecurity(password string, needsSpecialCharacters bool) bool {
	bytePassword := []byte(password)

	// Passwords does require numeric characters
	if !regexp.MustCompile(`[0-9]`).Match(bytePassword) {
		return false
	}

	// Passwords does require lowercase alphabetical characters
	if !regexp.MustCompile(`[a-z]`).Match(bytePassword) {
		return false
	}

	// Passwords does require uppercase alphabetical characters
	if !regexp.MustCompile(`[A-Z]`).Match(bytePassword) {
		return false
	}

	// If request was to require special characters check for their existence
	if needsSpecialCharacters && !regexp.MustCompile(`[^a-zA-Z0-9]`).Match(bytePassword) {
		return false
	}

	return true
}

func (*SecurePassword) hasCharacterRepetition(password string) bool {
	for i := 1; i < len(password); i++ {
		if password[i-1] == password[i] {
			return true
		}
	}
	return false
}
