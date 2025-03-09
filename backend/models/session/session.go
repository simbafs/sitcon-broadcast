package session

import (
	"context"
	"time"

	"backend/ent"
	"backend/ent/session"

	m "backend/models"
)

func GetAllInRoom(ctx context.Context, room string) (ent.Sessions, error) {
	return m.Client.Session.Query().
		Where(session.Room(room)).
		Order(ent.Asc(session.FieldStart)).
		All(ctx)
}

func Get(ctx context.Context, room string, id string) (*ent.Session, error) {
	return m.Client.Session.Query().
		Where(session.Room(room), session.SessionID(id)).
		Only(ctx)
}

func GetCurrent(ctx context.Context, room string) (*ent.Session, error) {
	// TODO: now package
	tw, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return nil, err
	}
	n := time.Date(2025, 3, 8, 5, 50, 0, 0, tw).Unix()

	return m.Client.Session.Query().
		Where(session.Room(room), session.EndGT(n)).
		Order(ent.Asc(session.FieldStart)).
		First(ctx)
		// All(ctx)
}
