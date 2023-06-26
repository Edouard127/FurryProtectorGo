package database

import (
	"sync"
	"time"
)

type InMemoryCache[T, K comparable] struct {
	mu    sync.RWMutex
	data  map[T]*K
	queue map[T]int64 // may be nil if purge is false

	doPurge    bool
	timeoutEat int64
	stop       int32
}

const (
	eatInterval      = time.Millisecond * 1000
	defaultQueueSize = 0
)

func NewInMemoryCache[T, K comparable](timeout int64, purge bool, size int64) *InMemoryCache[T, K] {
	if size < 0 {
		size = defaultQueueSize
	}
	c := &InMemoryCache[T, K]{
		data:       make(map[T]*K, size),
		doPurge:    purge,
		timeoutEat: timeout,
	}
	if purge {
		c.queue = make(map[T]int64, size)
		go c.start()
	}
	return c
}

func (c *InMemoryCache[T, K]) Get(key T) (*K, bool) {
	c.mu.RLock()
	value, ok := c.data[key]
	if ok {
		c.updateExpiration(key)
	}
	c.mu.RUnlock()
	return value, ok
}

func (c *InMemoryCache[T, K]) Set(key T, value K) {
	c.mu.Lock()
	c.data[key] = &value
	if c.doPurge {
		c.queue[key] = time.Now().UnixMilli() + c.timeoutEat
	}
	c.mu.Unlock()
}

func (c *InMemoryCache[T, K]) Delete(key T) {
	c.mu.Lock()
	delete(c.data, key)
	delete(c.queue, key)
	c.mu.Unlock()
}

func (c *InMemoryCache[T, K]) Close() {
	c.stop = 1
	c.data = nil
	c.queue = nil
}

func (c *InMemoryCache[T, K]) start() {
	eatTicker := time.NewTicker(eatInterval)

	for {
		select {
		case <-eatTicker.C:
			if c.stop == 1 {
				goto exit
			}
			c.devour()
		}
	}

exit:
	eatTicker.Stop()
}

func (c *InMemoryCache[T, K]) updateExpiration(key T) {
	if c.doPurge {
		c.queue[key] = time.Now().UnixMilli() + c.timeoutEat
	}
}

func (c *InMemoryCache[T, K]) devour() {
	c.mu.Lock()
	now := time.Now().UnixMilli()
	for key := range c.queue {
		if c.queue[key] < now {
			delete(c.queue, key)
			delete(c.data, key)
		}
	}
	c.mu.Unlock()
}
