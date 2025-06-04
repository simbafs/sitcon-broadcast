package counter

import (
	"time"

	"backend/internal/logger"
)

var log = logger.New("counter")

type Counter struct {
	Init     int  `json:"init"`
	Count    int  `json:"count"`
	Counting bool `json:"counting"`

	cb   Callback
	stop chan struct{}
}

func NewCounter(init int, callback Callback) *Counter {
	return &Counter{
		Init:     init,
		Count:    init,
		Counting: false,

		cb: callback,

		stop: make(chan struct{}),
	}
}

func (c *Counter) Tick() bool {
	c.Count--
	if c.Count < 0 {
		return true
	}
	c.cb(c)
	return false
}

func (c *Counter) Start() {
	c.Counting = true
	defer func() { c.Counting = false }()

	t := time.Tick(1 * time.Second)
	if c.Tick() {
		return
	}
	for {
		select {
		case <-c.stop:
			log.Println("counter stopped")
			return
		case <-t:
			if c.Tick() {
				return
			}
		}
	}
}

func (c *Counter) Stop() {
	if !c.Counting {
		return
	}
	// use select to avoid blocking if Stop is called multiple times
	select {
	case c.stop <- struct{}{}:
	default:
	}
}

func (c *Counter) Reset() {
	c.Stop()
	c.Count = c.Init
	c.cb(c)
}

func (c *Counter) SetInit(init int) {
	c.Stop()
	c.Init = init
	c.Count = init
	c.cb(c)
}
