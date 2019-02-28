package hasher

import (
	"strings"
	"testing"

	"github.com/Luzifer/go_helpers/v2/str"
)

func TestAvailableHashs(t *testing.T) {
	hashs, err := GetHashMap("testpass")
	if err != nil {
		t.Fatalf("Hash map generation failed: %s", err)
	}

	for _, impl := range []string{
		"htpasswd_apr1",
		"htpasswd_bcrypt",
		"htpasswd_sha256",
		"htpasswd_sha512",
		"sha1",
		"sha256",
		"sha512",
	} {
		if _, ok := hashs[impl]; !ok {
			t.Errorf("Hash implementation %q is missing", impl)
		}
	}
}

func TestSalt(t *testing.T) {
	knownSalts := []string{}

	for i := 0; i < 100; i++ {
		salt, err := getSalt()
		if err != nil {
			t.Fatalf("Hash generation failed: %s", err)
		}

		if len(salt) != saltSize {
			t.Errorf("Salt did not have desired size of %d: %d", saltSize, len(salt))
		}

		ssalt := string(salt)
		if str.StringInSlice(ssalt, knownSalts) {
			t.Fatalf("Received collision of hashes: %q", ssalt)
		}

		knownSalts = append(knownSalts, ssalt)
	}
}

func TestHTPasswd(t *testing.T) {
	hashs, err := GetHashMap("testpass")
	if err != nil {
		t.Fatalf("Hash map generation failed: %s", err)
	}

	if len(hashs["htpasswd_sha512"]) != 102 || !strings.HasPrefix(hashs["htpasswd_sha512"], "$6$") {
		t.Errorf("Invalid htpasswd SHA512 hash: %q", hashs["htpasswd_sha512"])
	}

	if len(hashs["htpasswd_sha256"]) != 59 || !strings.HasPrefix(hashs["htpasswd_sha256"], "$5$") {
		t.Errorf("Invalid htpasswd SHA256 hash: %q", hashs["htpasswd_sha256"])
	}

	if len(hashs["htpasswd_apr1"]) != 37 || !strings.HasPrefix(hashs["htpasswd_apr1"], "$apr1$") {
		t.Errorf("Invalid htpasswd APR1 hash: %q", hashs["htpasswd_apr1"])
	}
}

func TestStandardHashs(t *testing.T) {
	hashs, err := GetHashMap("testpass")
	if err != nil {
		t.Fatalf("Hash map generation failed: %s", err)
	}

	if hashs["sha1"] != "206c80413b9a96c1312cc346b7d2517b84463edd" {
		t.Errorf("Invalid SHA1 hash: %q", hashs["sha1"])
	}

	if hashs["sha256"] != "13d249f2cb4127b40cfa757866850278793f814ded3c587fe5889e889a7a9f6c" {
		t.Errorf("Invalid SHA256 hash: %q", hashs["sha256"])
	}

	if hashs["sha512"] != "78ddc8555bb1677ff5af75ba5fc02cb30bb592b0610277ae15055e189b77fe3fda496e5027a3d99ec85d54941adee1cc174b50438fdc21d82d0a79f85b58cf44" {
		t.Errorf("Invalid SHA512 hash: %q", hashs["sha512"])
	}
}
