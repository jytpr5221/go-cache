package cache

import (
	"sync"
	"time"
)

type Entry struct {
	value string
	ttl   time.Time 
	//TODO: for future extension: add a heap based ttlTracker, that keeps entries in exprationTime order and then removes them in order.
}

type Cache struct {
	storage map[string]Entry
	lock    sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		storage: make(map[string]Entry),
	}
}

func (cache *Cache) Set(key string, val string, ttl uint) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	expiry := time.Now().Add(time.Duration(ttl) * time.Second)

	cache.storage[key] = Entry{
		value: val,
		ttl:   expiry,
	}
}

func (cache *Cache) Get(key string) (bool, string) {
	cache.lock.RLock()

	entry, ok := cache.storage[key]
	if !ok {
		cache.lock.RUnlock()
		return false, ""
	}

	// Lazy expiration
	if time.Now().After(entry.ttl) {
		cache.lock.RUnlock()

		cache.Del(key)

		return false, ""
	}

	value := entry.value
	cache.lock.RUnlock()

	return true, value
}

func (cache *Cache) Del(key string) {
	cache.lock.Lock()
	defer cache.lock.Unlock()

	delete(cache.storage, key)
}