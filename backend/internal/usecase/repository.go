package usecase

import (
	"context"

	"backend/internal/entity"
	"backend/internal/repository"
)

type EventRepository interface {
	List(ctx context.Context) ([]*entity.Event, error)
	Get(ctx context.Context, name string) (*entity.Event, error)

	Create(ctx context.Context, name, url, script string) (*entity.Event, error)
	Update(ctx context.Context, name, url, script string) error

	Delete(ctx context.Context, name string) error
}

var _ EventRepository = &repository.EventEntRepository{}
