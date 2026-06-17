package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cacheEntry map[string]CacheEntry
	mu         sync.RWMutex
	interval   time.Duration
}

func (c *Cache) Add(key string, value []byte) {
	newCachEntry := CacheEntry{
		val:       value,
		createdAt: time.Now(),
	}
	defer c.mu.Unlock()
	c.mu.Lock()
	c.cacheEntry[key] = newCachEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	defer c.mu.RUnlock()
	c.mu.RLock()
	if entry, ok := c.cacheEntry[key]; ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	defer c.mu.Unlock()
	c.mu.Lock()
	for entryKey, entryValue := range c.cacheEntry {
		if time.Since(entryValue.createdAt) > c.interval {
			delete(c.cacheEntry, entryKey)
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheEntry: make(map[string]CacheEntry),
		interval:   interval,
	}
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			cache.reapLoop()
		}
	}()
	return cache
}
