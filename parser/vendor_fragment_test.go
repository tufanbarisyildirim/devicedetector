package parser

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVendorParse(t *testing.T) {
	type VendorFixture struct {
		Vendor    string `yaml:"vendor"`
		UserAgent string `yaml:"useragent"`
	}
	var vendorParser, _ = NewVendor(filepath.Join(dir, FixtureFileVendor))
	var list []VendorFixture
	err := ReadYamlFile(`fixtures/vendorfragments.yml`, &list)
	if err != nil {
		t.Error(err)
	}

	for _, item := range list {
		ua := item.UserAgent
		r := vendorParser.Parse(ua)
		require.Equal(t, item.Vendor, r, ua)
	}
}
