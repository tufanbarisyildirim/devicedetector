package parser

import (
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

const dir = "../regexes"

func TestVendors(t *testing.T) {
	v, err := NewVendor(filepath.Join(dir, FixtureFileVendor))
	require.NoError(t, err)
	str, _ := json.Marshal(v)
	require.Equal(t, "{}", string(str))
}

func TestReg(t *testing.T) {
	name := `Chrome(?:/(\d+[\.\d]+))?`
	ua := `Mozilla/5.0 (Linux; Android 4.2.2; ARCHOS 101 PLATINUM Build/JDQ39) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/34.0.1847.114 Safari/537.36`
	matches := MatchUserAgent(ua, name)

	require.Equal(t, " Chrome/34.0.1847.114", matches[0])
	require.Equal(t, "34.0.1847.114", matches[1])
}
