package database

import "sync"

type InMemoryCache[T, K comparable] struct {
	mu   sync.RWMutex
	data map[T]K
}

func NewInMemoryCache[T, K comparable]() *InMemoryCache[T, K] {
	return &InMemoryCache[T, K]{mu: sync.RWMutex{}, data: make(map[T]K)}
}

func (c *InMemoryCache[T, K]) Get(key T) (value K, ok bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok = c.data[key]
	return
}

func (c *InMemoryCache[T, K]) Set(key T, value K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *InMemoryCache[T, K]) Delete(key T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}
