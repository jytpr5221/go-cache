package cache

import (
	// "fmt"
	"sync"
)

var cache map[string]string = make(map[string]string)

type Entry struct{
	value string
}

type Cache struct{
	storage map[string]string
	lock sync.RWMutex
}

func NewCache() *Cache{

	return &Cache{
		storage: make(map[string]string),
	}
}

func (cache *Cache)Set(key string, val string){

	cache.lock.Lock()
	defer cache.lock.Unlock()

	cache.storage[key] = val
}

func (cache *Cache)Get(key string) (bool, string){

	cache.lock.RLock()
	defer cache.lock.RUnlock()

	val, ok := cache.storage[key]

	return ok, val
}

func (cache *Cache)Del(key string){

	cache.lock.Lock()
	defer cache.lock.Unlock()

	delete(cache.storage, key)
}