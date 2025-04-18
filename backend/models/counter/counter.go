package counter

type Status int

const (
	StatusStopped Status = iota // when trigger with start(), it will reset Count = Init
	StatusPause                 // when trigger with start(), it won't reset Count
	StaatusCounting
)

type Counter struct {
	Name     string    `json:"name" doc:"counter name"`
	Init     int       `json:"init" doc:"initial value"`
	Count    int       `json:"count" doc:"current value"`
	Status   Status    `json:"status" doc:"counter status, 0: stopped, 1: pause, 2: counting"`
	Callback func(int) `json:"-"` // Callback will be called when the counter change
}

func (c *Counter) Start() {
	if c.Status == StatusStopped {
		c.Count = c.Init
	}
	c.Status = StaatusCounting

	if c.Callback != nil {
		c.Callback(c.Count)
	}
}

func (c *Counter) Pause() {
	if c.Status == StaatusCounting {
		c.Status = StatusPause
	}

	if c.Callback != nil {
		c.Callback(c.Count)
	}
}

func (c *Counter) Stop() {
	if c.Status == StaatusCounting {
		c.Status = StatusStopped
	}

	if c.Callback != nil {
		c.Callback(c.Count)
	}
}

// Update will return true if the counter is finished
func (c *Counter) Update() {
	p := c.Count
	if c.Status == StaatusCounting {
		c.Count--
	}
	if c.Count <= 0 {
		c.Count = 0
		c.Status = StatusStopped
	}

	if p != c.Count {
		if c.Callback != nil {
			c.Callback(c.Count)
		}
	}
}

func (c *Counter) Set(n int) {
	c.Status = StatusStopped
	c.Count = n
	c.Init = n

	if c.Callback != nil {
		c.Callback(c.Count)
	}
}
