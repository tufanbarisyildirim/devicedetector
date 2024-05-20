package devicedetector

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCache(t *testing.T) {
	// Initialize the cache
	cache := NewCache()

	require.NotNil(t, cache)
	require.Equal(t, map[string]*DeviceInfo{}, cache.cache)
}

func TestAddToCache(t *testing.T) {
	// Initialize the cache
	cache := NewCache()

	require.NotNil(t, cache)
	require.Equal(t, map[string]*DeviceInfo{}, cache.cache)

	// Add an element to the cache
	deviceInfo := &DeviceInfo{
		userAgent: "test-user-agent",
	}

	cache.Add("test-user-agent", deviceInfo)

	// Verify the element has been cached
	require.NotEmpty(t, cache.cache)

	cachedDeviceInfo, hit := cache.Lookup("test-user-agent")

	require.NotNil(t, cachedDeviceInfo)
	require.True(t, hit)
	require.EqualValues(t, deviceInfo, cachedDeviceInfo)
}

func TestLookupCache(t *testing.T) {
	// Initialize the cache
	cache := NewCache()

	require.NotNil(t, cache)
	require.Equal(t, map[string]*DeviceInfo{}, cache.cache)

	// Add an element to the cache
	deviceInfo := &DeviceInfo{
		userAgent: "test-user-agent",
	}

	cache.Add("test-user-agent", deviceInfo)

	// Verify the element has been cached
	require.NotEmpty(t, cache.cache)

	cachedDeviceInfo, hit := cache.Lookup("test-user-agent")

	require.NotNil(t, cachedDeviceInfo)
	require.True(t, hit)
	require.EqualValues(t, deviceInfo, cachedDeviceInfo)

	// Verify the cache in case of miss of the elmenent
	cachedDeviceInfo, hit = cache.Lookup("not-cached-user-agent")

	require.Nil(t, cachedDeviceInfo)
	require.False(t, hit)
}

func TestPurgeCache(t *testing.T) {
	// Initialize the cache
	cache := NewCache()

	require.NotNil(t, cache)
	require.Equal(t, map[string]*DeviceInfo{}, cache.cache)

	// Add an element to the cache
	deviceInfo := &DeviceInfo{
		userAgent: "test-user-agent",
	}

	cache.Add("test-user-agent", deviceInfo)

	// Verify the element has been cached
	require.NotEmpty(t, cache.cache)

	cachedDeviceInfo, hit := cache.Lookup("test-user-agent")

	require.NotNil(t, cachedDeviceInfo)
	require.True(t, hit)
	require.EqualValues(t, deviceInfo, cachedDeviceInfo)

	// Verify the cache after the purge
	cache.Purge()

	require.Equal(t, map[string]*DeviceInfo{}, cache.cache)
}
