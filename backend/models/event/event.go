package event

import (
	"context"

	"backend/ent"
	"backend/ent/event"

	m "backend/models"
)

func NewEvent(ctx context.Context, name string, url string) (*ent.Event, error) {
	return m.Client.Event.Create().
		SetName(name).
		SetURL(url).
		Save(ctx)
}

func GetEvent(ctx context.Context, name string) (*ent.Event, error) {
	return m.Client.Event.
		Query().
		Where(event.Name(name)).
		Only(ctx)
}
