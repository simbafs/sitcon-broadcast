package repository

import (
	"context"

	"backend/internal/entity"
)

type Event interface {
	List(ctx context.Context) ([]*entity.Event, error)
	Get(ctx context.Context, name string) (*entity.Event, error)

	Create(ctx context.Context, name, url, script string) (*entity.Event, error)
	Update(ctx context.Context, name, url, script string) error

	Delete(ctx context.Context, name string) error
}

type Now interface {
	Get() int64
	Set(int64)
	Reset()
}

type Counter interface {
	List() []*entity.Counter
	Get(name string) (*entity.Counter, error)
	New(name string, init int, callback func(*entity.Counter)) (*entity.Counter, error)
}
