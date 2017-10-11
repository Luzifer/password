// Package securepassword implements a password generator and check.
package securepassword // import "github.com/Luzifer/password/lib"

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

// SecurePassword provides methods for generating secure passwords and
// checking the security requirements of passwords
type SecurePassword struct {
	CharacterTables map[string]string
	InsecurePattern []string
	BadCharacters   []string
}

var (
	// ErrLengthTooLow represents an error thrown if the password will
	// never be able match the security considerations in this package
	ErrLengthTooLow = errors.New("Passwords with a length lower than 4 will never meet the security requirements")
	// ErrGeneratorMissingFeatures represents an error thrown if the generator
	// is missing a character set required for generating passwords
	ErrGeneratorMissingFeatures = errors.New("Password generator was not initialized with all required features")
	// DefaultGenerator contains a pre-filled generator using the libraries
	// default values for character tables, insecure patterns and bad
	// characters. Mostly this is the implementation to use except you do
	// have special requirements such as a limit of special characters.
	DefaultGenerator = NewSecurePassword()
)

// NewSecurePassword initializes a new SecurePassword generator
func NewSecurePassword() *SecurePassword {
	return &SecurePassword{
		CharacterTables: map[string]string{
			"numeric": "0123456789",
			"simple":  "abcdefghijklmnopqrstuvwxyz",
			"special": "!#$%&()*+,-_./:;=?@[]^{}~|",
		},
		InsecurePattern: []string{
			"abcdefghijklmnopqrstuvwxyz",      // Alphabet
			"zyxwvutsrqponmlkjihgfedcba",      // Alphabet reversed
			"01234567890",                     // Numeric increasing
			"09876543210",                     // Numeric decreasing
			"qwertzuiopasdfghjklyxcvbnm",      // German keyboard layout
			"mnbvcxylkjhgfdsapoiuztrewq",      // German keyboard layout reversed
			"qwertyuiopasdfghjklzxcvbnm",      // US keyboard layout
			"mnbvcxzlkjhgfdsapoiuytrewq",      // US keyboard layout reversed
			"789_456_123_147_258_369_159_753", // Numpad patterns
		},
		BadCharacters: []string{"I", "l", "0", "O", "B", "8"}, // Characters that could lead to confusion due to font
	}
}

// GeneratePassword generates a new password with a given length and
// optional special characters in it. The password is automatically
// checked against CheckPasswordSecurity in order to only deliver secure
// passwords.
func (s *SecurePassword) GeneratePassword(length int, special bool) (string, error) {
	// Sanity check
	if length < 4 {
		return "", ErrLengthTooLow
	}

	if err := s.validateGenerator(); err != nil {
		return "", err
	}

	characterTable := strings.Join([]string{
		s.CharacterTables["simple"],
		strings.ToUpper(s.CharacterTables["simple"]),
		s.CharacterTables["numeric"],
	}, "")
	if special {
		characterTable = strings.Join([]string{characterTable, s.CharacterTables["special"]}, "")
	}

	password := ""
	rand.Seed(time.Now().UnixNano())
	for {
		char := string(characterTable[rand.Intn(len(characterTable))])
		if strings.Contains(strings.Join(s.BadCharacters, ""), char) {
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
		for _, pattern := range s.InsecurePattern {
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

func (s *SecurePassword) matchesBasicSecurity(password string, needsSpecialCharacters bool) bool {
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

	// If request was to require special characters check for their existance
	if needsSpecialCharacters && !regexp.MustCompile(`[^a-zA-Z0-9]`).Match(bytePassword) {
		return false
	}

	return true
}

func (s *SecurePassword) hasCharacterRepetition(password string) bool {
	for i := 1; i < len(password); i++ {
		if password[i-1] == password[i] {
			return true
		}
	}
	return false
}

func (s *SecurePassword) validateGenerator() error {
	if _, ok := s.CharacterTables["numeric"]; !ok {
		return ErrGeneratorMissingFeatures
	}

	if _, ok := s.CharacterTables["simple"]; !ok {
		return ErrGeneratorMissingFeatures
	}

	if _, ok := s.CharacterTables["special"]; !ok {
		return ErrGeneratorMissingFeatures
	}

	return nil
}
