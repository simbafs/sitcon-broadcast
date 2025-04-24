package event

import (
	"context"

	"backend/ent"
	"backend/ent/event"
)

type Event struct {
	client *ent.Client
}

func New(client *ent.Client) *Event {
	return &Event{
		client: client,
	}
}

func (e *Event) NewEvent(ctx context.Context, name, url, script string) (*ent.Event, error) {
	if script == "" {
		script = "function main(data) {\n  return []\n}"
	}

	return e.client.Event.Create().
		SetName(name).
		SetURL(url).
		SetScript(script).
		Save(ctx)
}

func (e *Event) GetAll(ctx context.Context) ([]*ent.Event, error) {
	return e.client.Event.Query().All(ctx)
}

func (e *Event) Get(ctx context.Context, name string) (*ent.Event, error) {
	return e.client.Event.
		Query().
		Where(event.Name(name)).
		Only(ctx)
}

func (e *Event) UpdateScript(ctx context.Context, name, script string) (*ent.Event, error) {
	// TODO:
	event, err := e.client.Event.
		Query().
		Where(event.Name(name)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return event.Update().
		SetScript(script).
		Save(ctx)
}
