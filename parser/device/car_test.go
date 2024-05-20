package device

import (
	"path/filepath"
	"testing"

	"github.com/gamebtc/devicedetector/parser"
	"github.com/stretchr/testify/require"
)

func TestCarParse(t *testing.T) {
	ps := NewCar(filepath.Join(dir, FixtureFileCar))
	var list []*DeviceFixture
	err := parser.ReadYamlFile(`fixtures/car_browser.yml`, &list)
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
