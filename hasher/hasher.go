// Package hasher contains methods to generate hashes for a given password
package hasher

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"

	"github.com/GehirnInc/crypt"
	"github.com/GehirnInc/crypt/apr1_crypt"
	"github.com/GehirnInc/crypt/sha256_crypt"
	"github.com/GehirnInc/crypt/sha512_crypt"
	"golang.org/x/crypto/bcrypt"
)

type hasherFunc func(password string) (string, error)

var implementations = map[string]hasherFunc{
	"htpasswd_apr1":   implHTAPR1,
	"htpasswd_bcrypt": implBcrypt,
	"htpasswd_sha256": implHTSHA256,
	"htpasswd_sha512": implHTSHA512,
	"sha256":          implSHA256,
	"sha512":          implSHA512,
}

// GetHashMap generates a map of hashes of the given password
func GetHashMap(password string) (map[string]string, error) {
	result := map[string]string{}

	for name, hf := range implementations {
		h, err := hf(password)
		if err != nil {
			return nil, fmt.Errorf("hashing %s: %w", name, err)
		}

		result[name] = h
	}

	return result, nil
}

func generic(h hash.Hash, password string) (_ string, err error) {
	if _, err = h.Write([]byte(password)); err != nil {
		return "", fmt.Errorf("writing password to hash: %w", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func genericHT(c crypt.Crypter, password string) (string, error) {
	h, err := c.Generate([]byte(password), nil) // Salt is auto-generated
	if err != nil {
		return "", fmt.Errorf("generating hash: %w", err)
	}

	return h, nil
}

func implBcrypt(password string) (string, error) {
	bc, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("generating bcrypt hash: %w", err)
	}
	return string(bc), nil
}

func implHTAPR1(password string) (string, error)   { return genericHT(apr1_crypt.New(), password) }
func implHTSHA256(password string) (string, error) { return genericHT(sha256_crypt.New(), password) }
func implHTSHA512(password string) (string, error) { return genericHT(sha512_crypt.New(), password) }

func implSHA256(password string) (string, error) { return generic(sha256.New(), password) }
func implSHA512(password string) (string, error) { return generic(sha512.New(), password) }
