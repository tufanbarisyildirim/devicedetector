package device

import (
	"path/filepath"
	"testing"

	"github.com/gamebtc/devicedetector/parser"
	"github.com/stretchr/testify/require"
)

func TestConsoleParse(t *testing.T) {
	ps := NewConsole(filepath.Join(dir, FixtureFileConsole))
	var list []*DeviceFixture
	err := parser.ReadYamlFile(`fixtures/console.yml`, &list)
	if err != nil {
		t.Error(err)
	}

	for _, item := range list {
		ua := item.UserAgent
		r := ps.Parse(ua)
		test := item.GetDeviceMatchResult()
		require.EqualValues(t, test, r)
	}
}
