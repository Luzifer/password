package hasher

import (
	"crypto"
	"crypto/rand"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"

	"fmt"
	"math/big"

	"golang.org/x/crypto/bcrypt"

	"github.com/tredoe/osutil/user/crypt"
	_ "github.com/tredoe/osutil/user/crypt/apr1_crypt"
	_ "github.com/tredoe/osutil/user/crypt/sha256_crypt"
	_ "github.com/tredoe/osutil/user/crypt/sha512_crypt"
)

type hasherFunc func(password string) (string, error)

var (
	implementations = map[string]hasherFunc{
		"htpasswd_apr1":   implHTAPR1,
		"htpasswd_bcrypt": implBcrypt,
		"htpasswd_sha256": implHTSHA256,
		"htpasswd_sha512": implHTSHA512,
		"sha1":            implSHA1,
		"sha256":          implSHA256,
		"sha512":          implSHA512,
	}

	saltSet  = []byte(`abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789./`)
	saltSize = 12
)

func GetHashMap(password string) (map[string]string, error) {
	result := map[string]string{}

	for name, hf := range implementations {
		h, err := hf(password)
		if err != nil {
			return nil, err
		}

		result[name] = h
	}

	return result, nil
}

func getSalt() ([]byte, error) {
	salt := make([]byte, saltSize)
	saltSetLength := big.NewInt(int64(len(saltSet)))

	for i := 0; i < saltSize; i++ {
		pos, err := rand.Int(rand.Reader, saltSetLength)
		if err != nil {
			return nil, err
		}
		salt[i] = saltSet[pos.Int64()]
	}

	return salt, nil
}

func implBcrypt(password string) (string, error) {
	bc, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bc), err
}

func genericHT(password, prefix string) (string, error) {
	salt, err := getSalt()
	if err != nil {
		return "", err
	}

	return crypt.NewFromHash(prefix+string(salt)).Generate([]byte(password), append([]byte(prefix), salt...))
}

func implHTAPR1(password string) (string, error)   { return genericHT(password, "$apr1$") }
func implHTSHA256(password string) (string, error) { return genericHT(password, "$5$") }
func implHTSHA512(password string) (string, error) { return genericHT(password, "$6$") }

func generic(password string, h crypto.Hash) (string, error) {
	w := h.New()
	w.Write([]byte(password))
	return fmt.Sprintf("%x", w.Sum(nil)), nil
}

func implSHA1(password string) (string, error)   { return generic(password, crypto.SHA1) }
func implSHA256(password string) (string, error) { return generic(password, crypto.SHA256) }
func implSHA512(password string) (string, error) { return generic(password, crypto.SHA512) }
