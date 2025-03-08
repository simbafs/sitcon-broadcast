package session

import (
	"context"
	_ "embed"
	"errors"
	"log"
	"sync"

	"backend/ent"
	"backend/ent/session"
	"backend/models/now"

	m "backend/models"

	_ "github.com/mattn/go-sqlite3"
)

// no Create

// ReadAll
func ReadAll(ctx context.Context) ([]*ent.Session, error) {
	return m.Client.Session.Query().
		Order(ent.Asc(session.FieldStart)).
		All(ctx)
}

func ReadAllInRoom(ctx context.Context, room string) ([]*ent.Session, error) {
	return m.Client.Session.Query().
		Where(session.Room(room)).
		Order(ent.Asc(session.FieldStart)).
		All(ctx)
}

// ReadByID
func ReadByID(ctx context.Context, room string, id string) (*ent.Session, error) {
	return m.Client.Session.Query().
		Where(session.And(session.Room(room), session.ID(id))).
		Only(ctx)
}

func ReadCurrentByRoom(ctx context.Context, room string) (*ent.Session, error) {
	sessions, err := ReadAllInRoom(ctx, room)
	if err != nil {
		return nil, err
	}

	n := now.Read()

	for _, v := range sessions {
		if v.End > int64(n) {
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
			log.Println("found", v.ID, k)
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

var mutex = sync.Mutex{}

func Update(ctx context.Context, u *ent.Session) error {
	mutex.Lock()
	defer mutex.Unlock()

	prev, curr, next, err := ReadPrevNext(ctx, u.Room, u.ID)
	if err != nil {
		return err
	}

	if prev != nil && u.Start <= prev.Start {
		return errors.New("start time cannot be before previous session's start time")
		// } else {
		// 	prev.End = u.Start
	}
	if next != nil && u.End >= next.End {
		return errors.New("end time cannot be after next session's end time")
		// } else {
		// 	next.Start = u.End
	}

	// 更新當前 Session
	err = curr.Update().
		SetTitle(u.Title).
		SetStart(u.Start).
		SetEnd(u.End).
		SetSpeaker(u.Speaker).
		SetQa(u.Qa).
		SetSlidoID(u.SlidoID).
		SetSlidoAdminLink(u.SlidoAdminLink).
		SetCoWrite(u.CoWrite).
		Exec(ctx)
	if err != nil {
		return err
	}

	// 更新前一個 Session 的 end 時間
	if prev != nil {
		err = prev.Update().
			SetEnd(u.Start).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	// 更新下一個 Session 的 start 時間
	if next != nil {
		err = next.Update().
			SetStart(u.End).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// no Delete
