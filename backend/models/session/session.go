package session

import (
	"context"
	"errors"
	"fmt"

	"backend/ent"
	"backend/ent/session"
	"backend/internal/logger"
	"backend/models/now"

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
	n := int64(now.GetNow())

	s, err := m.Client.Session.Query().
		Where(session.Room(room), session.EndGT(n)).
		Order(ent.Asc(session.FieldStart)).
		First(ctx)
	if ent.IsNotFound(err) {
		return m.Client.Session.Query().
			Order(ent.Desc(session.FieldStart)).
			First(ctx)
	}
	return s, err
}

var (
	ErrEndBeforeStart = errors.New("end time is before start time")
	ErrStartAfterEnd  = errors.New("start time is after end time")
)

func SetEnd(ctx context.Context, room string, id string, end int64) (*ent.Session, error) {
	curr, err := Get(ctx, room, id)
	if err != nil {
		return nil, err
	}

	next, err := Get(ctx, room, curr.Next)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	// TODO: check if the time is valid
	if end <= curr.Start {
		return nil, fmt.Errorf("%s %s, %w", curr.Room, curr.SessionID, ErrEndBeforeStart)
	}
	if next != nil && end >= next.End {
		return nil, fmt.Errorf("%s %s, %w", next.Room, next.SessionID, ErrStartAfterEnd)
	}

	err = curr.Update().
		SetEnd(end).
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
