package models

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"backend/ent"

	_ "github.com/mattn/go-sqlite3"
)

var Client *ent.Client

func InitDB(db string) error {
	if err := os.MkdirAll(filepath.Dir(db), 0o755); err != nil {
		return fmt.Errorf("failed to create directory for db: %w", err)
	}
	var err error
	Client, err = ent.Open("sqlite3", "file:"+db+"?cache=shared&_fk=1")
	if err != nil {
		return err
	}

	if err := Client.Schema.Create(context.Background()); err != nil {
		return err
	}
	return nil
}
