package pokeapi

import (
	"fmt"
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
	var cache Cache
	cache.mu = &sync.Mutex{}
	cache.reapLoop(interval)
	return cache
}

func (cache Cache) Add(key string, val []byte) Cache {
	cache.mu.Lock()
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	cache.cache[key] = newEntry
	cache.mu.Unlock()
	return cache
}

func (cache Cache) Get(key string) (val []byte, b bool) {
	cache.mu.Lock()
	if cache.cache[key].val != nil {
		cache.mu.Unlock()
		return cache.cache[key].val, true
	}
	fmt.Printf("%s not found", key)
	cache.mu.Unlock()
	return nil, false

}

func (cache Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.mu.Lock()
		current := time.Now()
		for key := range cache.cache {
			initial := cache.cache[key].createdAt
			duration := current.Sub(initial)
			if duration > interval {
				delete(cache.cache, key)
			}
		}
		cache.mu.Unlock()
	}
}
