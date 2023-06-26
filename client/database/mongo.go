package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.Background()

type MongoCache[V any] struct {
	cache      *InMemoryCache[string, V]
	client     *Database
	collection string
}

func NewMongoCache[V any](db *Database, collection string, timeout int64) *MongoCache[V] {
	m := &MongoCache[V]{
		NewInMemoryCache[string, V](timeout, true, 0),
		db, collection,
	}
	go m.check()
	return m
}

func (c *MongoCache[V]) check() {
	for {
		select {
		case key := <-c.cache.onDelete:
			c.invalidate(key)
		}
	}
}

func (c *MongoCache[V]) Get(key string) V {
	value, ok := c.cache.Get(key)
	if !ok {
		var v V
		c.client.Read(c.collection, bson.M{"id": key}, &v)
		c.cache.Set(key, v)
		return v
	}
	return value
}

func (c *MongoCache[V]) Set(key string, value V) {
	c.cache.Set(key, value)
}

// invalidate will delete the cache entry and update the database with the new value
func (c *MongoCache[V]) invalidate(key string) {
	value, _ := c.cache.Get(key)
	c.client.Replace(c.collection, bson.M{"id": key}, value)
	c.cache.Delete(key)
}
