package cache

import (
	"sync"
	"time"
)

type Eviction interface {
	OnSet(key string, meta *Metadata) (evictedKey string, evicted bool)
	OnGet(key string, meta *Metadata)
	OnDel(key string, meta *Metadata)
}

type Metadata struct {
	EvictionNode any
}

type Entry struct {
	value    string
	ttl      time.Time
	metadata Metadata
}

type Cache struct {
	storage  map[string]Entry
	lock     sync.Mutex
	eviction Eviction
}

type EvictionPolicy int

const (
	_                   = iota
	LRU  EvictionPolicy = 1
	LFU  EvictionPolicy = 2
	FIFO EvictionPolicy = 3
)
