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

	Cb Callback

	ticker *time.Ticker
	stop   chan struct{}
}

func NewCounter(init int, callback Callback) *Counter {
	return &Counter{
		Init:     init,
		Count:    init,
		Counting: false,

		Cb: callback,

		ticker: time.NewTicker(time.Second),
		stop:   make(chan struct{}),
	}
}

func (c *Counter) Start() {
	c.Counting = true
	defer func() { c.Counting = false }()

	for {
		select {
		case <-c.stop:
			log.Println("counter stopped")
			return
		case <-c.ticker.C:
			c.Count--
			c.Cb(c)
			log.Println("counter", "count", c.Count)
			if c.Count <= 0 {
				log.Println("counter", "count", c.Count)
				return
			}
		}
	}
}

func (c *Counter) Stop() {
	if !c.Counting {
		return
	}
	c.stop <- struct{}{}
}

func (c *Counter) Reset() {
	c.Stop()
	c.Count = c.Init
	c.Cb(c)
}

func (c *Counter) SetInit(init int) {
	c.Stop()
	c.Init = init
	c.Count = init
	c.Cb(c)
}
