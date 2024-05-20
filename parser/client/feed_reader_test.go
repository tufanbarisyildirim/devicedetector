package client

import (
	"path/filepath"
	"testing"

	"github.com/gamebtc/devicedetector/parser"
	"github.com/stretchr/testify/require"
)

func TestFeedReaderParse(t *testing.T) {
	ps := NewFeedReader(filepath.Join(dir, FixtureFileFeedReader))
	var list []*ClientFixture
	err := parser.ReadYamlFile(`fixtures/feed_reader.yml`, &list)
	if err != nil {
		t.Error(err)
	}

	for _, item := range list {
		ua := item.UserAgent
		r := ps.Parse(ua)
		require.EqualValues(t, item.ClientMatchResult, r)
	}
}
