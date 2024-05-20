package client

import (
	"path/filepath"
	"testing"

	"github.com/gianluca-marchini/devicedetector/parser"
	"github.com/stretchr/testify/require"
)

func TestPimParse(t *testing.T) {
	ps := NewPim(filepath.Join(dir, FixtureFilePim))
	var list []*ClientFixture
	err := parser.ReadYamlFile(`fixtures/pim.yml`, &list)
	if err != nil {
		t.Error(err)
	}

	for _, item := range list {
		ua := item.UserAgent
		r := ps.Parse(ua)
		require.EqualValues(t, item.ClientMatchResult, r)
	}
}
