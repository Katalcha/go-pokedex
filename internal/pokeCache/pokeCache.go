package pokeCache

import (
	"sync"
	"time"
)

// CACHE
type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// NEWCACHE
func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	go cache.readLoop(interval)

	return cache
}

// ADD
func (cache *Cache) Add(key string, value []byte) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	cache.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

// GET
func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	val, ok := cache.cache[key]
	return val.val, ok
}

// READLOOP
func (cache *Cache) readLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.reap(time.Now().UTC(), interval)
	}
}

// REAP
func (cache *Cache) reap(now time.Time, last time.Duration) {
	cache.mux.Lock()
	defer cache.mux.Unlock()
	for key, value := range cache.cache {
		if value.createdAt.Before(now.Add(-last)) {
			delete(cache.cache, key)
		}
	}
}
