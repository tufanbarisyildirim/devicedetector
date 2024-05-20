package client

import (
	"path/filepath"
	"testing"

	"github.com/gianluca-marchini/devicedetector/parser"
	"github.com/stretchr/testify/require"
)

func TestLibraryParse(t *testing.T) {
	var ps = NewLibrary(filepath.Join(dir, FixtureFileLibrary))
	var list []*ClientFixture
	err := parser.ReadYamlFile(`fixtures/library.yml`, &list)
	if err != nil {
		t.Error(err)
	}

	for _, item := range list {
		ua := item.UserAgent
		r := ps.Parse(ua)
		require.EqualValues(t, item.ClientMatchResult, r)
	}
}
