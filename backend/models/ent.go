package models

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"backend/ent"
	"backend/models/event"
	"backend/models/session"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(db string) (*event.Event, *session.Session, error) {
	if err := os.MkdirAll(filepath.Dir(db), 0o755); err != nil {
		return nil, nil, fmt.Errorf("failed to create directory for db: %w", err)
	}
	var err error
	client, err := ent.Open("sqlite3", "file:"+db+"?cache=shared&_fk=1")
	if err != nil {
		return nil, nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, nil, err
	}
	return event.New(client), session.New(client), nil
}
