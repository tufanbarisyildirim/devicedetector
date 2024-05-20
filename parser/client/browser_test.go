package client

import (
	"path/filepath"
	"testing"

	"github.com/gamebtc/devicedetector/parser"
	"github.com/stretchr/testify/require"
)

func TestBrowserParse(t *testing.T) {
	ps := NewBrowser(filepath.Join(dir, FixtureFileBrowser))
	var list []*ClientFixture
	err := parser.ReadYamlFile(`fixtures/browser.yml`, &list)
	if err != nil {
		t.Error(err)
	}

	for _, item := range list {
		ua := item.UserAgent
		r := ps.Parse(ua)
		require.EqualValues(t, item.ClientMatchResult, r)
	}
}
