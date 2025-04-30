package usecase

import (
	"backend/internal/repository"
)

var _ Now = &NowImpl{}

type NowImpl struct {
	now repository.Now
}

func NewNow(now repository.Now) *NowImpl {
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
