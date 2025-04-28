package repository

import (
	"context"
	"errors"

	"backend/internal/entity"
)

type EventInMemoryRepository struct {
	events map[string]*entity.Event
}

func (e *EventInMemoryRepository) List(ctx context.Context) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0, len(e.events))
	for _, event := range e.events {
		events = append(events, event)
	}
	return events, nil
}

func (e *EventInMemoryRepository) Get(ctx context.Context, name string) (*entity.Event, error) {
	event, exists := e.events[name]
	if !exists {
		return nil, entity.ErrCannotGetSessions
	}
	return event, nil
}

var ErrEventAlreadyExists = errors.New("event already exists")

func (e *EventInMemoryRepository) Create(ctx context.Context, name string, url string, script string) (*entity.Event, error) {
	if _, ok := e.events[name]; ok {
		return nil, ErrEventAlreadyExists
	}
	event := entity.NewEvent(name, url, script)
	e.events[name] = event
	return event, nil
}

func (e *EventInMemoryRepository) Update(ctx context.Context, name string, url string, script string) error {
	event, ok := e.events[name]
	if !ok {
		return entity.ErrCannotGetSessions
	}

	event.SetURL(url)
	event.SetScript(script)

	e.events[name] = event

	return nil
}

func (e *EventInMemoryRepository) Delete(ctx context.Context, name string) error {
	if _, ok := e.events[name]; !ok {
		return nil
	}

	delete(e.events, name)
	return nil
}

func NewEventInMemoryRepository() *EventInMemoryRepository {
	return &EventInMemoryRepository{
		events: make(map[string]*entity.Event),
	}
}
