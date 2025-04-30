package repository

import (
	"sync"

	"backend/internal/entity"
)

var _ Now = &NowImpl{}

var (
	once sync.Once
	now  *entity.Now
)

// prevent usecase layer directly depending on entity layer, maybe this is redundant, but I left it anyway.
type NowImpl struct{}

func NewNow() *NowImpl {
	once.Do(func() {
		now = entity.NewNow(0)
	})
	return &NowImpl{}
}

func (n *NowImpl) Get() int64 {
	return now.Get()
}

func (n *NowImpl) Set(t int64) {
	now.Set(t)
}

func (n *NowImpl) Reset() {
	now.Reset()
}
