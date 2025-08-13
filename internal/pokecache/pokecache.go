package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cache -
type Cache struct {
	mux    *sync.Mutex
	values map[string]cacheEntry
}

// NewCache -
func NewCache(interval time.Duration) Cache {
	cache := Cache{
		mux:    &sync.Mutex{},
		values: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.values[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	value, ok := c.values[key]

	return value.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for k, v := range c.values {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.values, k)
		}
	}
}
