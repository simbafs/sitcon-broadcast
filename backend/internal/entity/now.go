package entity

import (
	"sync"
	"time"
)

type Now struct {
	mu  sync.RWMutex
	now int64
}

func NewNow(now int64) *Now {
	return &Now{
		mu:  sync.RWMutex{},
		now: now,
	}
}

func (n *Now) Get() int64 {
	n.mu.RLock()
	defer n.mu.RUnlock()

	if n.now == 0 {
		return time.Now().Unix()
	}
	return n.now
}

func (n *Now) Set(now int64) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.now = now
}

func (n *Now) Reset() {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.now = 0
}
