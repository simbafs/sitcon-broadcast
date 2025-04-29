package repository

import (
	"context"
	"errors"

	"backend/internal/entity"
)

var ErrCannotGetEvent = errors.New("can not get event")

type EventInMemory struct {
	events map[string]*entity.Event
}

func NewEventInMemory() *EventInMemory {
	return &EventInMemory{
		events: make(map[string]*entity.Event),
	}
}

func (e *EventInMemory) List(ctx context.Context) ([]*entity.Event, error) {
	events := make([]*entity.Event, 0, len(e.events))
	for _, event := range e.events {
		events = append(events, event)
	}
	return events, nil
}

func (e *EventInMemory) Get(ctx context.Context, name string) (*entity.Event, error) {
	event, exists := e.events[name]
	if !exists {
		return nil, ErrCannotGetEvent
	}
	return event, nil
}

var ErrEventAlreadyExists = errors.New("event already exists")

func (e *EventInMemory) Create(ctx context.Context, name string, url string, script string) (*entity.Event, error) {
	if _, ok := e.events[name]; ok {
		return nil, ErrEventAlreadyExists
	}
	event := entity.NewEvent(name, url, script)
	e.events[name] = event
	return event, nil
}

func (e *EventInMemory) Update(ctx context.Context, name string, url string, script string) error {
	event, ok := e.events[name]
	if !ok {
		return ErrCannotGetEvent
	}

	event.SetURL(url)
	event.SetScript(script)

	e.events[name] = event

	return nil
}

func (e *EventInMemory) Delete(ctx context.Context, name string) error {
	if _, ok := e.events[name]; !ok {
		return nil
	}

	delete(e.events, name)
	return nil
}
