package memcache

import (
	"fmt"
	"time"
)

func NewCache(cleanUpInterval, invalidAfter time.Duration) (newCache *Cache) {
	newCache = &Cache{
		cacheDataTimeDuration: invalidAfter,
		cachedData:            make(map[string]*data),
		cancelCacheCleanChan:  make(chan bool),
		cacheCleanerTicker:    time.NewTicker(cleanUpInterval),
	}

	go cleaner(newCache)

	return
}

func cleaner(cache *Cache) {
LOOP:
	for {
		select {
		case <-cache.cancelCacheCleanChan:
			break LOOP
		case <-cache.cacheCleanerTicker.C:
			fmt.Println("CleanUp!")
		}
	}
}