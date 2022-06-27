package securepassword

import (
	"bufio"
	"crypto/sha1" //#nosec: G505 // HIBP uses shortened SHA1 to query hashes of vulnerable passwordss
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// ErrPasswordInBreach signals the password passed was found in any
// breach at least once. The password should not be used if this
// error is returned.
var ErrPasswordInBreach = errors.New("given password is known to HaveIBeenPwned")

// CheckHIBPPasswordHash accesses the HaveIBeenPwned API with the
// first 5 characters of the SHA1 hash of the password and scans the
// result for the password hash. If the hash is found the
// ErrPasswordInBreach error is thrown. In case of an HTTP error
// another error is thrown. The result will be nil when the password
// hash was not returned in the API output.
//
// See more details at https://haveibeenpwned.com/API/v2#PwnedPasswords
func CheckHIBPPasswordHash(password string) error {
	fullHash := fmt.Sprintf("%x", sha1.Sum([]byte(password))) //#nosec: G401 // See crypto/sha1 import
	checkHash := fullHash[0:5]

	resp, err := http.Get("https://api.pwnedpasswords.com/range/" + checkHash)
	if err != nil {
		return errors.Wrap(err, "HTTP request failed")
	}
	defer resp.Body.Close()

	// Response format:
	// 0018A45C4D1DEF81644B54AB7F969B88D65:1
	// 00D4F6E8FA6EECAD2A3AA415EEC418D38EC:2

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), fullHash) {
			// We don't care for the exact number but only for a match
			return ErrPasswordInBreach
		}
	}

	return nil
}
