package entity

import "time"

type Counter struct {
	count    int
	init     int
	counting bool
	stop     chan struct{}
	callback func(*Counter)
}

func NewCounter(init int, callback func(*Counter)) *Counter {
	if callback == nil {
		callback = func(c *Counter) {}
	}
	return &Counter{
		count:    init,
		init:     init,
		counting: false,
		stop:     make(chan struct{}),
		callback: callback,
	}
}

func (c *Counter) Count() int {
	return c.count
}

func (c *Counter) Counting() bool {
	return c.counting
}

// TODO: check if this function ok
func (c *Counter) SetInit(init int) {
	defer c.callback(c)

	if c.counting {
		c.Stop()
	}
	c.init = init
	c.count = init
}

func (c *Counter) tick() {
	c.count--
	if c.count <= 0 {
		c.Stop()
	}
}

func (c *Counter) Start() {
	defer c.callback(c)

	c.counting = true
	go func() {
		defer func() { c.counting = false }()

		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				c.tick()
			case <-c.stop:
				return
			}
		}
	}()
}

func (c *Counter) Stop() {
	defer c.callback(c)

	if !c.counting {
		return
	}
	c.stop <- struct{}{}
}

func (c *Counter) Reset() {
	defer c.callback(c)

	c.Stop()
	c.count = c.init
}
