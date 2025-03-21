package event

import (
	"context"

	"backend/ent"
	"backend/ent/event"

	m "backend/models"
)

func NewEvent(ctx context.Context, name, url, script string) (*ent.Event, error) {
	if script == "" {
		script = "function main(data) {\n  return []\n}"
	}

	return m.Client.Event.Create().
		SetName(name).
		SetURL(url).
		SetScript(script).
		Save(ctx)
}

func GetAll(ctx context.Context) ([]*ent.Event, error) {
	return m.Client.Event.Query().All(ctx)
}

func Get(ctx context.Context, name string) (*ent.Event, error) {
	return m.Client.Event.
		Query().
		Where(event.Name(name)).
		Only(ctx)
}

func UpdateScript(ctx context.Context, name, script string) (*ent.Event, error) {
	e, err := m.Client.Event.
		Query().
		Where(event.Name(name)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return e.Update().
		SetScript(script).
		Save(ctx)
}
