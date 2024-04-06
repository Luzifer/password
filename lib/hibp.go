package securepassword

import (
	"bufio"
	"context"
	"crypto/sha1" //#nosec: G505 // HIBP uses shortened SHA1 to query hashes of vulnerable passwordss
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const hibpTimeout = 2 * time.Second

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

	ctx, cancel := context.WithTimeout(context.TODO(), hibpTimeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://api.pwnedpasswords.com/range/%s", checkHash), nil)
	if err != nil {
		return fmt.Errorf("creating HTTP request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("executing HTTP request: %w", err)
	}
	defer resp.Body.Close() //nolint:errcheck

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
