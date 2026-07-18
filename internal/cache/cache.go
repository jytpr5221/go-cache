package cache

import "time"

const MAX_CACHE_SIZE = 3

func NewCache(policy EvictionPolicy) *Cache {
	if policy == LRU{
	    return NewCacheWithEviction(NewLRUEviction())
	}else if policy == LFU{
		return NewCacheWithEviction((NewLFUEviction()))
	}

	return NewCacheWithEviction(NewLRUEviction())
}

func NewCacheWithEviction(eviction Eviction) *Cache {
    if eviction == nil {
        eviction = NewLRUEviction()
    }

    return &Cache{
        storage:  make(map[string]Entry),
        eviction: eviction,
    }
}

func (cache *Cache) Set(key string, val string, ttl uint) {
    cache.lock.Lock()
    defer cache.lock.Unlock()

    expiry := time.Now().Add(time.Duration(ttl) * time.Second)

    entry := cache.storage[key]
    entry.value = val
    entry.ttl = expiry

    if cache.eviction != nil {
        evictedKey, evicted := cache.eviction.OnSet(key, &entry.metadata)
        cache.storage[key] = entry

        if evicted {
            delete(cache.storage, evictedKey)
        }
        return
    }

    cache.storage[key] = entry
}

func (cache *Cache) Get(key string) (string, bool) {
    cache.lock.RLock()
    defer cache.lock.RUnlock()

    entry, ok := cache.storage[key]
    if !ok {
        return "", false
    }

    if time.Now().After(entry.ttl) {
        delete(cache.storage, key)

        if cache.eviction != nil {
            cache.eviction.OnDel(key, &entry.metadata)
        }

        return "", false
    }

    if cache.eviction != nil {
        cache.eviction.OnGet(key, &entry.metadata)
        cache.storage[key] = entry
    }

    return entry.value, true
}

func (cache *Cache) Del(key string) {
    cache.lock.Lock()
    defer cache.lock.Unlock()

    entry, ok := cache.storage[key]
    if !ok {
        return
    }

    delete(cache.storage, key)

    if cache.eviction != nil {
        cache.eviction.OnDel(key, &entry.metadata)
    }
}