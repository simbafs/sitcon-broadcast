package models

import (
	"backend/ent"
)

// TODO: maybe move the initialize of ent to another files?
var Client *ent.Client

func init() {
	// var err error
	// Client, err = ent.Open("sqlite3", "file:sessions.db?cache=shared&_fk=1")
	// if err != nil {
	// 	panic(err)
	// }
	//
	// if err = Client.Schema.Create(context.Background()); err != nil {
	// 	panic(err)
	// }
}
