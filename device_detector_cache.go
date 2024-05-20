package devicedetector

type Cache struct {
	cache map[string]*DeviceInfo
}

func NewCache() *Cache {
	return &Cache{
		make(map[string]*DeviceInfo),
	}
}

// Associate a deviceInfo element with the userAgent
func (d *Cache) Add(ua string, deviceInfo *DeviceInfo) {
	d.cache[ua] = deviceInfo
}

// Look for a cached userAgent: if found, hit is true.
func (d *Cache) Lookup(ua string) (deviceInfo *DeviceInfo, hit bool) {
	deviceInfo, hit = d.cache[ua]

	return deviceInfo, hit
}

// Purge the cache
func (d *Cache) Purge() {
	d.cache = make(map[string]*DeviceInfo)
}
