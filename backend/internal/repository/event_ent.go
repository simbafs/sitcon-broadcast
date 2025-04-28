package repository

import (
	"context"

	"backend/ent"
	"backend/ent/event"
	"backend/internal/entity"
)

type EventEntRepository struct {
	client *ent.Client
}

func NewEventEntRepository(client *ent.Client) *EventEntRepository {
	return &EventEntRepository{
		client: client,
	}
}

func (r *EventEntRepository) List(ctx context.Context) ([]*entity.Event, error) {
	events, err := r.client.Event.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*entity.Event, len(events))
	for i, e := range events {
		results[i] = entity.NewEvent(e.Name, e.URL, e.Script)
	}

	return results, nil
}

func (r *EventEntRepository) Get(ctx context.Context, name string) (*entity.Event, error) {
	e, err := r.client.Event.Query().Where(event.Name(name)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return entity.NewEvent(e.Name, e.URL, e.Script), nil
}

func (r *EventEntRepository) Create(ctx context.Context, name, url, script string) (*entity.Event, error) {
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

func (r *EventEntRepository) Update(ctx context.Context, name, url, script string) error {
	update := r.client.Event.Update().
		Where(event.Name(name))

	if url != "" {
		update = update.SetURL(url)
	}

	if script != "" {
		update = update.SetScript(script)
	}

	err := update.Exec(ctx)
	return err
}

func (r *EventEntRepository) Delete(ctx context.Context, name string) error {
	_, err := r.client.Event.Delete().
		Where(event.Name(name)).
		Exec(ctx)
	return err
}
