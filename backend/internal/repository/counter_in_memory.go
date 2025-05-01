package repository

import (
	"errors"
	"maps"
	"slices"
	"sync"

	"backend/internal/entity"
)

var (
	ErrCannotGetCounter = errors.New("can not get counter")
	ErrCounterExist     = errors.New("counter already exists")
)

var _ Counter = &CounterInMemory{}

type CounterInMemory struct {
	counters map[string]*entity.Counter
	mu       sync.RWMutex
}

func NewCounterInMemory(counters map[string]*entity.Counter) *CounterInMemory {
	if counters == nil {
		counters = make(map[string]*entity.Counter)
	}
	return &CounterInMemory{
		counters: counters,
	}
}

func (c *CounterInMemory) List() []*entity.Counter {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := slices.Sorted(maps.Keys(c.counters))
	counters := make([]*entity.Counter, len(keys))
	for i, k := range keys {
		counters[i] = c.counters[k]
	}
	return counters
}

func (c *CounterInMemory) Get(name string) (*entity.Counter, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	counter, ok := c.counters[name]
	if !ok {
		return nil, ErrCannotGetCounter
	}
	return counter, nil
}

func (c *CounterInMemory) New(name string, init int, callback func(*entity.Counter)) (*entity.Counter, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.counters[name]; ok {
		return nil, ErrCounterExist
	}
	counter := entity.NewCounter(init, callback)
	c.counters[name] = counter
	return counter, nil
}

func (c *CounterInMemory) Delete(name string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.counters[name]; !ok {
		return nil
	}
	delete(c.counters, name)
	return nil
}
