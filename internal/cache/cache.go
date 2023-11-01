package cache

import (
	"sync"
	"time"
)

type CacheEntry[T any] struct {
	Value     T
	ExpiresAt time.Time
}

type Cache[K comparable, V any] struct {
	mu              sync.RWMutex
	data            map[K]CacheEntry[V]
	cleanupInterval time.Duration
	stop            chan struct{}
}

func NewCache[K comparable, V any]() *Cache[K, V] {
	cache := &Cache[K, V]{
		data:            make(map[K]CacheEntry[V]),
		cleanupInterval: 24 * time.Hour,
		stop:            make(chan struct{}),
	}
	if cache.cleanupInterval > 0 {
		go cache.cleanup()
	}
	return cache
}

func (c *Cache[K, V]) Set(key K, value V, expiry time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = CacheEntry[V]{
		Value:     value,
		ExpiresAt: time.Now().Add(expiry),
	}
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	entry, exists := c.data[key]
	if !exists {
		c.mu.RUnlock()
		var zero V
		return zero, false
	}
	if time.Now().After(entry.ExpiresAt) {
		c.mu.RUnlock()
		c.Delete(key)
		var zero V
		return zero, false
	}
	c.mu.RUnlock()
	return entry.Value, true
}

func (c *Cache[K, V]) Delete(key K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func (c *Cache[K, V]) Exist(key K) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	entry, exists := c.data[key]
	return exists && !time.Now().After(entry.ExpiresAt)
}

func (c *Cache[K, V]) cleanup() {
	ticker := time.NewTicker(c.cleanupInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			now := time.Now()
			for k, v := range c.data {
				if now.After(v.ExpiresAt) {
					delete(c.data, k)
				}
			}
			c.mu.Unlock()
		case <-c.stop:
			return
		}
	}
}

func (c *Cache[K, V]) Stop() {
	close(c.stop)
}
