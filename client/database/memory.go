package database

import (
	"container/heap"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type InMemoryCache[T, K comparable] struct {
	mu    sync.RWMutex
	data  *sync.Map
	queue *priorityQueue[T]

	timeoutEat time.Duration
	stop       int32
}

const (
	eatInterval      = time.Second
	defaultQueueSize = 1024
)

func NewInMemoryCache[T, K comparable](timeout int) *InMemoryCache[T, K] {
	c := &InMemoryCache[T, K]{
		data:       &sync.Map{},
		queue:      newPriorityQueue[T](defaultQueueSize),
		timeoutEat: time.Duration(timeout) * time.Millisecond,
	}
	go c.start()
	return c
}

func (c *InMemoryCache[T, K]) Get(key T) (K, bool) {
	value, ok := c.data.Load(any(key))
	if ok {
		c.updateExpiration(key)
	}
	return value.(K), ok
}

func (c *InMemoryCache[T, K]) Set(key T, value K) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data.Store(key, value)
	c.updateExpiration(key)
}

func (c *InMemoryCache[T, K]) Delete(key T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data.Delete(key)
}

func (c *InMemoryCache[T, K]) Close() {
	atomic.StoreInt32(&c.stop, 1)
}

func (c *InMemoryCache[T, K]) start() {
	eatTicker := time.NewTicker(eatInterval)
	defer eatTicker.Stop()

	for {
		select {
		case <-eatTicker.C:
			if atomic.LoadInt32(&c.stop) == 1 {
				return
			}
			c.devour()
		}
	}
}

func (c *InMemoryCache[T, K]) updateExpiration(key T) {
	expiration := time.Now().Add(c.timeoutEat)
	c.removeCachedDevorer(key)
	c.queue.push(&devourItem[T]{value: key, t: expiration})
}

func (c *InMemoryCache[T, K]) removeCachedDevorer(key T) {
	for i := 0; i < c.queue.length; i++ {
		if c.queue.items[i].value == key {
			heap.Remove(c.queue, i)
			break
		}
	}
}

func (c *InMemoryCache[T, K]) devour() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for {
		item := c.queue.peek()
		if item == nil || item.t.After(now) {
			break
		}
		fmt.Println("Deleting", item.value)
		c.data.Delete(item.value)
		c.queue.pop()
	}
}

type devourItem[K comparable] struct {
	value K
	t     time.Time
}

type priorityQueue[K comparable] struct {
	items  []*devourItem[K]
	length int
}

func newPriorityQueue[K comparable](capacity int) *priorityQueue[K] {
	return &priorityQueue[K]{
		items:  make([]*devourItem[K], 0, capacity),
		length: 0,
	}
}

func (pq *priorityQueue[K]) push(item *devourItem[K]) {
	heap.Push(pq, item)
}

func (pq *priorityQueue[K]) pop() *devourItem[K] {
	if pq.length == 0 {
		return nil
	}
	return heap.Pop(pq).(*devourItem[K])
}

func (pq *priorityQueue[K]) peek() *devourItem[K] {
	if pq.length == 0 {
		return nil
	}
	return pq.items[0]
}

func (pq *priorityQueue[K]) Len() int { return pq.length }

func (pq *priorityQueue[K]) Less(i, j int) bool {
	return pq.items[i].t.Before(pq.items[j].t)
}

func (pq *priorityQueue[K]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func (pq *priorityQueue[K]) Push(x interface{}) {
	item := x.(*devourItem[K])
	pq.items = append(pq.items, item)
	pq.length++
}

func (pq *priorityQueue[K]) Pop() interface{} {
	item := pq.items[pq.length-1]
	pq.items = pq.items[:pq.length-1]
	pq.length--
	return item
}
