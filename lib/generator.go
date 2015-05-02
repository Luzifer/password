package securepassword // import "github.com/Luzifer/password/lib"

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type SecurePassword struct {
	characterTables map[string]string
	insecurePattern []string
	badCharacters   []string
}

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

func (s *SecurePassword) GeneratePassword(length int, special bool) string {
	characterTable := strings.Join([]string{
		s.characterTables["simple"],
		strings.ToUpper(s.characterTables["simple"]),
		s.characterTables["numeric"],
	}, "")
	if special {
		characterTable = strings.Join([]string{characterTable, s.characterTables["special"]}, "")
	}

	password := ""
	rand.Seed(time.Now().UnixNano())
	for {
		password = fmt.Sprintf("%s%s",
			password,
			string(characterTable[rand.Intn(len(characterTable))]),
		)
		if len(password) == length {
			if s.CheckPasswordSecurity(password, special) {
				break
			}
			password = ""
		}
	}
	return password
}

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
