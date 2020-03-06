package memcache

import (
	"time"
)

func NewCache(cleanUpInterval, dataExpireAfter time.Duration) (newCache *Cache) {
	newCache = &Cache{
		cacheDataTimeDuration: dataExpireAfter,
		cachedData:            make(map[string]*Data),
		cancelCacheCleanChan:  make(chan bool),
		cacheCleanerTicker:    time.NewTicker(cleanUpInterval),
	}

	go cleaner(newCache)

	return
}

func DeleteCache(cache *Cache) {
	cache.stopCache()
	cache = nil
}

func cleaner(cache *Cache) {
LOOP:
	for {
		select {
		case <-cache.cancelCacheCleanChan:
			break LOOP
		case <-cache.cacheCleanerTicker.C:
			cache.cleanExpired()
		}
	}
}
