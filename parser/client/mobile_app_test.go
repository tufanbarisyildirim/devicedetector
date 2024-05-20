package client

import (
	"path/filepath"
	"testing"

	"github.com/gianluca-marchini/devicedetector/parser"
	"github.com/stretchr/testify/require"
)

func TestMediaAppParse(t *testing.T) {
	ps := NewMobileApp(filepath.Join(dir, FixtureFileMobileApp))
	var list []*ClientFixture
	err := parser.ReadYamlFile(`fixtures/mobile_app.yml`, &list)
	if err != nil {
		t.Error(err)
	}

	for _, item := range list {
		ua := item.UserAgent
		r := ps.Parse(ua)
		require.EqualValues(t, item.ClientMatchResult, r)
	}
}
