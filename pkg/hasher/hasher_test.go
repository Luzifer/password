package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAvailableHashs(t *testing.T) {
	hashs, err := GetHashMap("testpass")
	require.NoError(t, err)

	for impl := range implementations {
		assert.Contains(t, hashs, impl)
	}
}

func TestHTPasswd(t *testing.T) {
	hashs, err := GetHashMap("testpass")
	require.NoError(t, err)

	assert.Equal(t, "$apr1$", hashs["htpasswd_apr1"][:6])
	assert.Len(t, hashs["htpasswd_apr1"], 37)

	assert.Equal(t, "$5$", hashs["htpasswd_sha256"][:3])
	assert.Len(t, hashs["htpasswd_sha256"], 63)

	assert.Equal(t, "$6$", hashs["htpasswd_sha512"][:3])
	assert.Len(t, hashs["htpasswd_sha512"], 106)
}

func TestStandardHashs(t *testing.T) {
	hashs, err := GetHashMap("testpass")
	require.NoError(t, err)

	assert.Equal(t,
		"13d249f2cb4127b40cfa757866850278793f814ded3c587fe5889e889a7a9f6c",
		hashs["sha256"])

	assert.Equal(t,
		"78ddc8555bb1677ff5af75ba5fc02cb30bb592b0610277ae15055e189b77fe3fda496e5027a3d99ec85d54941adee1cc174b50438fdc21d82d0a79f85b58cf44",
		hashs["sha512"])
}
