// todo добавить удаление по ключу и просто чистку
package memcache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheDataTimeDuration time.Duration
	cachedData            map[string]*data

	mtx sync.Mutex

	cacheCleanerTicker   *time.Ticker
	cancelCacheCleanChan chan bool
}

type data struct {
	expirationDate time.Time
	dataValue      interface{}
}

func (cache *Cache) Set(key string, value interface{}) {
	cache.mtx.Lock()
	expirationForNewData := time.Now().Add(cache.cacheDataTimeDuration)
	cache.cachedData[key] = &data{
		expirationDate: expirationForNewData,
		dataValue:      value,
	}
	cache.mtx.Unlock()
}

func (cache *Cache) Get(key string) *data {
	cache.mtx.Lock()
	retValue, found := cache.cachedData[key]
	cache.mtx.Unlock()

	if !found {
		return nil
	}

	return retValue
}

func (cache *Cache) Keys() []string {
	keys := make([]string, 0, len(cache.cachedData))

	cache.mtx.Lock()
	for key := range cache.cachedData {
		keys = append(keys, key)
	}
	cache.mtx.Unlock()

	return keys
}
