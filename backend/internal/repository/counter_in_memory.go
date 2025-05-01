package repository

import (
	"errors"
	"maps"
	"slices"

	"backend/internal/entity"
)

var (
	ErrCannotGetCounter = errors.New("can not get counter")
	ErrCounterExist     = errors.New("counter already exists")
)

var _ Counter = &CounterInMemory{}

type CounterInMemory struct {
	counters map[string]*entity.Counter
}

func NewCounterInMemory() *CounterInMemory {
	return &CounterInMemory{
		counters: make(map[string]*entity.Counter),
	}
}

func (c *CounterInMemory) List() []*entity.Counter {
	keys := slices.Sorted(maps.Keys(c.counters))
	counters := make([]*entity.Counter, len(keys))
	for i, k := range keys {
		counters[i] = c.counters[k]
	}
	return counters
}

func (c *CounterInMemory) Get(name string) (*entity.Counter, error) {
	counter, ok := c.counters[name]
	if !ok {
		return nil, ErrCannotGetCounter
	}
	return counter, nil
}

func (c *CounterInMemory) New(name string, init int, callback func(*entity.Counter)) (*entity.Counter, error) {
	if _, ok := c.counters[name]; ok {
		return nil, ErrCounterExist
	}
	counter := entity.NewCounter(init, callback)
	c.counters[name] = counter
	return counter, nil
}
