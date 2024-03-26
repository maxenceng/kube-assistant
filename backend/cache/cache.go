package cache

import (
	"errors"
	"sync"
)

type Cache[T any] interface {
	Get(key string) (*Value[T], error)
	GetSelected() (*Value[T], error)
	Set(key string, value *Value[T])
	Delete(key string)
}

type Value[T any] struct {
	Value    T
	Selected bool
}

type cache[T any] struct {
	store map[string]*Value[T]
	lock  sync.RWMutex
}

func New[T any]() Cache[T] {
	return &cache[T]{
		store: make(map[string]*Value[T]),
	}
}

func (c *cache[T]) Get(key string) (*Value[T], error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	value, ok := c.store[key]
	if !ok {
		return nil, errors.New("key not found")
	}
	return value, nil
}

func (c *cache[T]) GetSelected() (*Value[T], error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	for _, value := range c.store {
		if value.Selected {
			return value, nil
		}
	}

	return nil, errors.New("selected not found")
}

func (c *cache[T]) Set(key string, value *Value[T]) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.store[key] = value
}

func (c *cache[T]) Delete(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.store, key)
}
