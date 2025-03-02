package session

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"time"

	"backend/ent"
	"backend/ent/session"
	"backend/models/now"

	_ "github.com/mattn/go-sqlite3"
)

// TODO: maybe move the initialize of ent to another files?
var Client *ent.Client

func init() {
	var err error
	Client, err = ent.Open("sqlite3", "file:sessions.db?cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}

	if err = Client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}
}

// no Create

// ReadAll
func ReadAll(ctx context.Context) ([]*ent.Session, error) {
	return Client.Session.Query().
		All(ctx)
}

// ReadByID
func ReadByID(ctx context.Context, id string) (*ent.Session, error) {
	return Client.Session.Query().
		Where(session.ID(id)).
		Only(ctx)
}

// ReadCurrentByRoom
func ReadCurrentByRoom(ctx context.Context, room string) (*ent.Session, error) {
	sessions, err := Client.Session.Query().
		Where(session.Room(room)).
		Order(ent.Asc(session.FieldStart)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	now := now.Read()

	var target *ent.Session
	for _, s := range sessions {
		if s.End.After(now) {
			target = s
			break
		}
	}

	if target == nil {
		target = sessions[len(sessions)-1]
	}

	return target, nil
}

func Update(ctx context.Context, id string, start, end time.Time) error {
	if start.After(end) {
		return fmt.Errorf("start time cannot be after end time")
	}

	sess, err := Client.Session.Get(ctx, id)
	if err != nil {
		return err
	}

	sameRoomSessions, err := Client.Session.
		Query().
		Where(session.Room(sess.Room)).
		Order(ent.Asc("start")).
		All(ctx)
	if err != nil {
		return err
	}

	var prevSession, nextSession *ent.Session
	for i, s := range sameRoomSessions {
		if s.ID == id {
			if i > 0 {
				prevSession = sameRoomSessions[i-1]
			}
			if i < len(sameRoomSessions)-1 {
				nextSession = sameRoomSessions[i+1]
			}
			break
		}
	}

	if prevSession != nil && start.Before(prevSession.Start) {
		return errors.New("start time cannot be before previous session's start time")
	}
	if nextSession != nil && end.After(nextSession.End) {
		return errors.New("end time cannot be after next session's end time")
	}

	// 更新當前 Session
	err = Client.Session.UpdateOneID(id).
		SetStart(start).
		SetEnd(end).
		Exec(ctx)
	if err != nil {
		return err
	}

	// 更新前一個 Session 的 end 時間
	if prevSession != nil {
		err = Client.Session.UpdateOneID(prevSession.ID).
			SetEnd(start).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	// 更新下一個 Session 的 start 時間
	if nextSession != nil {
		err = Client.Session.UpdateOneID(nextSession.ID).
			SetStart(end).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// no Delete
