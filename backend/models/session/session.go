package session

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"log"
	"slices"
	"sync"
	"time"

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
	return m.Client.Session.Query().
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
		if v.End > n.Time.Unix() {
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

func Update(ctx context.Context, room string, id string, start, end time.Time) error {
	mutex.Lock()
	defer mutex.Unlock()
	
	startT := start.Unix() * 1000
	endT := end.Unix() * 1000

	log.Println(id, start, end)
	if start.After(end) {
		return fmt.Errorf("start time cannot be after end time")
	}

	prev, _, next, err := ReadPrevNext(ctx, room, id)
	if err != nil {
		return err
	}

	if prev != nil {
		log.Println("start", startT, prev.Start)
		log.Println(endT <= prev.Start)

	} else {
		log.Println("prev is nil")
	}
	if next != nil {
		log.Println("end", endT, next.End)
		log.Println(startT >= next.End)
	} else {
		log.Println("next is nil")
	}

	if prev != nil && startT <= prev.Start {
		return errors.New("start time cannot be before previous session's start time")
	}
	if next != nil && endT >= next.End {
		return errors.New("end time cannot be after next session's end time")
	}

	// 更新當前 Session
	err = m.Client.Session.UpdateOneID(id).
		SetStart(startT).
		SetEnd(endT).
		Exec(ctx)
	if err != nil {
		return err
	}

	// 更新前一個 Session 的 end 時間
	if prev != nil {
		err = m.Client.Session.UpdateOneID(prev.ID).
			SetEnd(startT).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	// 更新下一個 Session 的 start 時間
	if next != nil {
		err = m.Client.Session.UpdateOneID(next.ID).
			SetStart(endT).
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// no Delete
