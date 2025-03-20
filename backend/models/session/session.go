package session

import (
	"context"
	"errors"

	"backend/ent"
	"backend/ent/session"
	"backend/internal/logger"

	m "backend/models"
)

var log = logger.New("models/session")

func GetAllInRoom(ctx context.Context, room string) (ent.Sessions, error) {
	return m.Client.Session.Query().
		Where(session.Room(room)).
		Order(ent.Asc(session.FieldIdx)).
		All(ctx)
}

func Get(ctx context.Context, room string, id string) (*ent.Session, error) {
	return m.Client.Session.Query().
		Where(session.Room(room), session.SessionID(id)).
		Only(ctx)
}

func GetCurrent(ctx context.Context, room string) (*ent.Session, error) {
	s, err := m.Client.Session.Query().
		Where(session.And(
			session.Room(room),
			session.Or(
				session.Finish(false),
				session.Next(""),
			),
		)).
		Order(ent.Asc(session.FieldIdx)).
		First(ctx)
	if ent.IsNotFound(err) {
		// get last
		return m.Client.Session.Query().
			Order(ent.Desc(session.FieldIdx)).
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

func rollback(tx *ent.Tx, err error) error {
	if e := tx.Rollback(); err != nil {
		return errors.Join(err, e)
	}
	return err
}

func UpdateAll(ctx context.Context, sessions ent.Sessions) error {
	_, err := m.Client.Session.Delete().Exec(ctx)
	if err != nil {
		return err
	}

	tx, err := m.Client.Tx(ctx)
	if err != nil {
		return err
	}

	for _, s := range sessions {
		_, err := tx.Session.Create().
			SetIdx(s.Idx).
			SetSessionID(s.SessionID).
			SetFinish(s.Finish).
			SetStart(s.Start).
			SetEnd(s.End).
			SetRoom(s.Room).
			SetNext(s.Next).
			SetTitle(s.Title).
			SetData(s.Data).
			Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
