package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		entries: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}

	go cache.reapLoop(interval)
	return cache
}

//Cache struct methods

// Add
func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.entries[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}

}

// Get
func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	data, ok := cache.entries[key]
	return data.val, ok
}

// reapLoop
func (cache *Cache) reapLoop(interval time.Duration) {
	//init a ticker with the interval
	ticker := time.NewTicker(interval)

	//calls reap every given interval time
	for range ticker.C {
		cache.reap(time.Now().UTC(), interval)
	}
}

func (cache *Cache) reap(now time.Time, lastCall time.Duration) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	for index, value := range cache.entries {
		//check if the cache was create before the -- now time *minus* the interval
		if value.createdAt.Before(now.Add(-lastCall)) {
			delete(cache.entries, index)
		}
	}
}
