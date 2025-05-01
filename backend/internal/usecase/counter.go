package usecase

import (
	"backend/internal/entity"
	"backend/internal/repository"
	"backend/sse"
)

var _ Counter = &CounterImpl{}

type CounterImpl struct {
	counter repository.Counter
	sse     *sse.SSE
}

func NewCounter(counter repository.Counter, sse *sse.SSE) *CounterImpl {
	return &CounterImpl{
		counter: counter,
		sse:     sse,
	}
}

func (c *CounterImpl) New(name string, init int) (*entity.Counter, error) {
	callback := func(counter *entity.Counter) {
		c.sse.Send <- sse.Msg{
			Topic: []string{"counter/" + name},
			Data: struct {
				Count      int  `json:"count"`
				IsCounting bool `json:"is_counting"`
			}{
				Count:      counter.Count(),
				IsCounting: counter.Counting(),
			},
		}
	}
	return c.counter.New(name, init, callback)
}

func (c *CounterImpl) List() []*entity.Counter {
	return c.counter.List()
}

func (c *CounterImpl) Get(name string) (*entity.Counter, error) {
	return c.counter.Get(name)
}

func (c *CounterImpl) Start(name string) error {
	counter, err := c.counter.Get(name)
	if err != nil {
		return err
	}

	counter.Start()
	return nil
}

func (c *CounterImpl) Stop(name string) error {
	counter, err := c.counter.Get(name)
	if err != nil {
		return err
	}

	counter.Stop()
	return nil
}

func (c *CounterImpl) SetInit(name string, init int) error {
	counter, err := c.counter.Get(name)
	if err != nil {
		return err
	}

	counter.SetInit(init)
	return nil
}
