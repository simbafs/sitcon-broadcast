package repository

import (
	"context"
	"errors"

	"backend/ent"
	"backend/ent/event"
	"backend/internal/entity"
)

var _ Event = &EventEnt{}

type EventEnt struct {
	client *ent.Client
}

func NewEventEnt(client *ent.Client) *EventEnt {
	return &EventEnt{
		client: client,
	}
}

func (r *EventEnt) List(ctx context.Context) ([]*entity.Event, error) {
	events, err := r.client.Event.Query().
		Order(event.ByName()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*entity.Event, len(events))
	for i, e := range events {
		results[i] = entity.NewEvent(e.Name, e.URL, e.Script)
	}

	return results, nil
}

func (r *EventEnt) Get(ctx context.Context, name string) (*entity.Event, error) {
	e, err := r.client.Event.Query().Where(event.Name(name)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return entity.NewEvent(e.Name, e.URL, e.Script), nil
}

func (r *EventEnt) Create(ctx context.Context, name, url, script string) (*entity.Event, error) {
	e, err := r.client.Event.Create().
		SetName(name).
		SetURL(url).
		SetScript(script).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity.NewEvent(e.Name, e.URL, e.Script), nil
}

var ErrEventNotFound = errors.New("event not found")

func (r *EventEnt) Update(ctx context.Context, name, url, script string) error {
	update := r.client.Event.Update().
		Where(event.Name(name))

	if url != "" {
		update = update.SetURL(url)
	}

	if script != "" {
		update = update.SetScript(script)
	}

	n, err := update.Save(ctx)
	if n == 0 {
		return ErrEventNotFound
	}
	return err
}

func (r *EventEnt) Delete(ctx context.Context, name string) error {
	_, err := r.client.Event.Delete().
		Where(event.Name(name)).
		Exec(ctx)
	return err
}
