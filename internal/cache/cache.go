package cache

import (
	"sync"
	"time"
)

type item struct {
	value     any
	expiresAt time.Time
}

func (i item) isExpired() bool {
	return time.Now().After(i.expiresAt)
}

// Cache is a thread-safe in-memory key-value store.
type Cache struct {
	mu    sync.RWMutex
	items map[string]item
}

// New initializes an empty Cache.
func New() *Cache {
	return &Cache{
		items: make(map[string]item),
	}
}

// Set saves an item with a Time-To-Live (TTL).
func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = item{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}
}

// Get retrieves an item by key.
// Returns (value, true) if found and not expired.
// Returns (nil, false) if missing or expired.
func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	it, exists := c.items[key]
	if !exists {
		return nil, false
	}

	if it.isExpired() {
		return nil, false
	}

	return it.value, true
}

// Delete immediately removes a key from the cache.
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
}