package memcache

import (
	"reflect"
	"testing"
	"time"
)

func prepareCache() *Cache {
	cleanUpInterval, _ := time.ParseDuration("5s")
	dataExpireAfter, _ := time.ParseDuration("3s")
	newCache := NewCache(cleanUpInterval, dataExpireAfter)
	newCache.Set("first", 123)
	newCache.Set("second", 234)

	return newCache
}

func TestCache_GetSet(t *testing.T) {
	newCache := prepareCache()

	expected := 123
	actual := newCache.Get("first")

	if actual.dataValue != expected {
		t.Errorf("General test failed. Got: %v", actual.dataValue)
	}
}

func TestCache_Keys(t *testing.T) {
	newCache := prepareCache()

	expected := []string{"first", "second"}
	actual := newCache.Keys()
	if !reflect.DeepEqual(actual, expected) {
		t.Error("Keys test failed. Got keys: ", actual)
	}
}

func TestCache_RemoveKeyC(t *testing.T) {
	newCache := prepareCache()
	newCache.RemoveKey("first")

	var expected *data = nil
	actual := newCache.Get("first")
	if actual != expected {
		t.Error("Remove test failed. Got: ", actual)
	}
}

func TestCache_CleanerOn(t *testing.T) {
	newCache := prepareCache()

	time.Sleep(10 * time.Second)

	actual := newCache.Get("first")
	if actual != nil {
		t.Error("CleanerOn test failed. Got: ", actual)
	}
}

func TestCache_StopCleaner(t *testing.T) {
	newCache := prepareCache()
	newCache.StopCleaner()

	time.Sleep(10 * time.Second)
	actual := newCache.Get("first")

	if actual == nil {
		t.Errorf("StopCleaner test failed. Got: %v", actual)
	}
}