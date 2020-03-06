package memcache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheDataTimeDuration time.Duration
	cachedData            map[string]*Data

	mtx sync.Mutex

	cacheCleanerTicker   *time.Ticker
	cancelCacheCleanChan chan bool
}

type Data struct {
	ExpirationDate time.Time
	DataValue      interface{}
}

func (cache *Cache) Set(key string, value interface{}) {
	cache.mtx.Lock()
	expirationForNewData := time.Now().Add(cache.cacheDataTimeDuration)
	cache.cachedData[key] = &Data{
		ExpirationDate: expirationForNewData,
		DataValue:      value,
	}
	cache.mtx.Unlock()
}

func (cache *Cache) Get(key string) *Data {
	cache.mtx.Lock()
	retValue, found := cache.cachedData[key]
	cache.mtx.Unlock()

	if !found {
		return nil
	}

	return retValue
}

func (cache *Cache) Keys() []string {
	cache.mtx.Lock()
	keys := make([]string, 0, len(cache.cachedData))
	for key := range cache.cachedData {
		keys = append(keys, key)
	}
	cache.mtx.Unlock()

	return keys
}

func (cache *Cache) stopCache() {
	cache.stopCleaner()
	cache.mtx.Lock()
	for k := range cache.cachedData {
		cache.unsafeRemove(k)
	}
	cache.mtx.Unlock()
}

func (cache *Cache) stopCleaner() {
	go func() {
		cache.cancelCacheCleanChan <- true
	}()
	cache.cacheCleanerTicker.Stop()

}

func (cache *Cache) RemoveKey(key string) {
	cache.mtx.Lock()
	cache.unsafeRemove(key)
	cache.mtx.Unlock()
}

func (cache *Cache) cleanExpired() {
	cache.mtx.Lock()
	for k, v := range cache.cachedData {
		if time.Now().After(v.ExpirationDate) {
			cache.unsafeRemove(k)
		}
	}
	cache.mtx.Unlock()
}

func (cache *Cache) unsafeRemove(key string) {
	delete(cache.cachedData, key)
}
