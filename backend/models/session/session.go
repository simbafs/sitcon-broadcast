package session

import (
	"context"

	"backend/ent"
	"backend/ent/session"
	"backend/internal/logger"

	m "backend/models"
)

var log = logger.New("models/session")

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
	s, err := m.Client.Session.Query().
		Where(session.Room(room), session.Finish(false)).
		Order(ent.Asc(session.FieldStart)).
		First(ctx)
	if ent.IsNotFound(err) {
		return m.Client.Session.Query().
			Order(ent.Desc(session.FieldStart)).
			First(ctx)
	}
	return s, err
}

// Next set the end time of current session and the start time of next session(if there is) to end
func Next(ctx context.Context, room string, id string, end int64) (*ent.Session, error) {
	curr, err := Get(ctx, room, id)
	if err != nil {
		return nil, err
	}

	next, err := Get(ctx, room, curr.Next)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	err = curr.Update().
		SetEnd(end).
		SetFinish(true).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	if next != nil {
		err = next.Update().
			SetStart(end).
			Exec(ctx)
		if err != nil {
			return nil, err
		}

	}

	return next, nil
}
