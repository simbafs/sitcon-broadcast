package models

import (
	"context"

	"backend/ent"

	_ "github.com/mattn/go-sqlite3"
)

var Client *ent.Client

func InitDB(db string) error {
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
