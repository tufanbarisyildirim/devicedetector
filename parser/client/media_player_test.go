package client

import (
	"path/filepath"
	"testing"

	"github.com/gianluca-marchini/devicedetector/parser"
	"github.com/stretchr/testify/require"
)

func TestMediaPlayerParse(t *testing.T) {
	ps := NewMediaPlayer(filepath.Join(dir, FixtureFileMediaPlayer))
	var list []*ClientFixture
	err := parser.ReadYamlFile(`fixtures/mediaplayer.yml`, &list)
	if err != nil {
		t.Error(err)
	}

	for _, item := range list {
		ua := item.UserAgent
		r := ps.Parse(ua)
		require.EqualValues(t, item.ClientMatchResult, r)
	}
}
