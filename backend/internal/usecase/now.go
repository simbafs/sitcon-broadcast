package usecase

import (
	"backend/internal/entity"
)

var _ Now = &NowImpl{}

type NowImpl struct {
	now *entity.Now
}

func NewNow(now *entity.Now) *NowImpl {
	return &NowImpl{
		now: now,
	}
}

type NowOutput struct {
	Now int64 `json:"now"`
}

type NowInput NowOutput

func (n *NowImpl) Get() *NowOutput {
	return &NowOutput{
		Now: n.now.Get(),
	}
}

func (n *NowImpl) Set(input *NowInput) *NowOutput {
	n.now.Set(input.Now)
	return &NowOutput{
		Now: n.now.Get(),
	}
}

func (n *NowImpl) Reset() *NowOutput {
	n.now.Reset()
	return &NowOutput{
		Now: n.now.Get(),
	}
}
