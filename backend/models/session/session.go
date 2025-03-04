package session

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"log"
	"slices"
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
		Order(ent.Asc(session.FieldStart)).
		All(ctx)
}

func ReadAllInRoom(ctx context.Context, room string) ([]*ent.Session, error) {
	s, err := ReadAll(ctx)
	if err != nil {
		return nil, err
	}

	r := []*ent.Session{}
	for _, v := range s {
		if v.Room == room || slices.Contains(v.Broadcast, room) {
			r = append(r, v)
		}
	}

	return r, nil
}

// ReadByID
func ReadByID(ctx context.Context, id string) (*ent.Session, error) {
	return Client.Session.Query().
		Where(session.ID(id)).
		Only(ctx)
}

func ReadCurrentByRoom(ctx context.Context, room string) (*ent.Session, error) {
	sessions, err := ReadAllInRoom(ctx, room)
	if err != nil {
		return nil, err
	}

	n := now.Read()

	for _, v := range sessions {
		if v.End.After(n.Time) {
			return v, nil
		}
	}

	return sessions[len(sessions)-1], nil
}

// prev, current, next, error
func ReadPrevNext(ctx context.Context, room string, id string) (*ent.Session, *ent.Session, *ent.Session, error) {
	sessions, err := ReadAllInRoom(ctx, room)
	if err != nil {
		return nil, nil, nil, err
	}

	var current *ent.Session
	var prev *ent.Session
	var next *ent.Session

	for k, v := range sessions {
		if v.ID == id {
			current = v
			if k > 0 {
				prev = sessions[k-1]
			}
			if k < len(sessions)-1 {
				next = sessions[k+1]
			}
			break
		}
	}

	if current == nil {
		return nil, nil, nil, errors.New("session not found")
	}

	return prev, current, next, nil
}

func Update(ctx context.Context, room string, id string, start, end time.Time) error {
	log.Println(id, start, end)
	if start.After(end) {
		return fmt.Errorf("start time cannot be after end time")
	}

	prev, _, next, err := ReadPrevNext(ctx, room, id)
	if err != nil {
		return err
	}

	if prev != nil && start.Before(prev.Start) {
		return errors.New("start time cannot be before previous session's start time")
	}
	if next != nil && end.After(next.End) {
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
	if prev != nil {
		err = Client.Session.UpdateOneID(prev.ID).
			SetEnd(start).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	// 更新下一個 Session 的 start 時間
	if next != nil {
		err = Client.Session.UpdateOneID(next.ID).
			SetStart(end).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// no Delete
