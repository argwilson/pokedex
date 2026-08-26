package pokeapi

import (
	"bytes"
	"sync"
	"time"
)

// Cache -
type Cache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// NewCache -
func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}

func (cache Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	cache.cache[key] = newEntry
}

func (cache Cache) Get(key string) (val []byte, b bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	checkEntry := cache.cache[key]
	if bytes.Equal(checkEntry.val, nil) {
		return nil, false
	}
	return checkEntry.val, true

}

func (cache Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.mu.Lock()
		defer cache.mu.Unlock()
		current := time.Now()
		for key := range cache.cache {
			initial := cache.cache[key].createdAt
			duration := current.Sub(initial)
			if duration >= interval {
				delete(cache.cache, key)
			}
		}
	}
}
