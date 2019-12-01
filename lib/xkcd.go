package securepassword

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/Luzifer/go_helpers/v2/str"
)

type XKCD struct {
	// Separator to be used between words
	Separator string
}

var (
	// ErrTooFewWords represents an error thrown if the password will
	// have fewer than four words and are not considered to be safe
	ErrTooFewWords = errors.New("XKCD passwords with less than 4 words makes no sense")
	// DefaultXKCD contains an default instance of the XKCD password
	// generator
	DefaultXKCD = NewXKCDGenerator()
)

// NewXKCDGenerator initializes a new XKCD password generator
// https://xkcd.com/936/
func NewXKCDGenerator() *XKCD { return &XKCD{} }

// GeneratePassword generates a password with the number of words
// given and optionally the current date prepended
func (x XKCD) GeneratePassword(length int, addDate bool) (string, error) {
	if length < 4 {
		return "", ErrTooFewWords
	}

	var (
		password  string
		usedWords []string
	)

	if addDate {
		password = time.Now().Format("20060102.")
	}

	rand.Seed(time.Now().UnixNano())
	for len(usedWords) < length {
		word := strings.Title(xkcdWordList[rand.Intn(len(xkcdWordList))])
		if str.StringInSlice(word, usedWords) {
			// Don't use a word twice
			continue
		}

		usedWords = append(usedWords, word)
	}

	return password + strings.Join(usedWords, x.Separator), nil
}
