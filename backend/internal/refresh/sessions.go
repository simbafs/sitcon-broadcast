// this package provide functions to update sessions in db to the latest data from the given URL.
package refresh

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"time"

	"backend/ent"
	"backend/internal/logger"

	lo "github.com/samber/lo"
	lop "github.com/samber/lo/parallel"

	_ "github.com/mattn/go-sqlite3"
)

var log = logger.New("refresh")

func parseInput(speakers map[string]string, rooms map[string]string) func(Session, int) []*ent.Session {
	return func(s Session, _ int) []*ent.Session {
		extraData := map[string]any{
			"speaker":     s.Speakers,
			"qa":          s.QA,
			"slide":       s.Slide,
			"co_write":    s.CoWrite,
			"record":      s.Record,
			"live":        s.Live,
			"tags":        s.Tags,
			"uri":         s.URI,
			"description": s.Zh.Description,
		}

		if len(s.Broadcast) == 0 {
			return []*ent.Session{
				{
					Idx:       0,
					Start:     parseTime(s.Start),
					End:       parseTime(s.End),
					SessionID: s.ID,
					Room:      rooms[s.Room],
					Next:      "",
					Title:     s.Zh.Title,
					Data:      extraData,
				},
			}
		} else {
			return lop.Map(s.Broadcast, func(room string, idx int) *ent.Session {
				return &ent.Session{
					Idx:       0,
					Start:     parseTime(s.Start),
					End:       parseTime(s.End),
					SessionID: s.ID,
					Room:      rooms[room],
					Next:      "",
					Title:     s.Zh.Title,
					Data:      extraData,
				}
			})
		}
	}
}

// mergeSameTitle is used after groupByRoom and before setNext
func mergeSameTitle(sessions []*ent.Session, curr *ent.Session, _ int) []*ent.Session {
	l := len(sessions)
	if l == 0 {
		return []*ent.Session{curr}
	}
	if sessions[l-1].Title == curr.Title {
		sessions[l-1].End = curr.End
		return sessions
	}
	return append(sessions, curr)
}

func removeSpeakerFromRest(s *ent.Session, _ int) *ent.Session {
	if slices.Contains([]string{"休息", "午休", "點心"}, s.Title) {
		s.Data["speaker"] = []string{}
	}
	return s
}

func setNextAndIdx(ss []*ent.Session) func(s *ent.Session, idx int) *ent.Session {
	return func(s *ent.Session, idx int) *ent.Session {
		if idx < len(ss)-1 {
			s.Next = ss[idx+1].Title
		}
		s.Idx = int8(idx)
		return s
	}
}

func sortByStart(sessions []*ent.Session) []*ent.Session {
	return slices.SortedFunc(slices.Values(sessions), func(a, b *ent.Session) int {
		return int(a.Start) - int(b.Start)
	})
}

func FromURL(url string) ([]*ent.Session, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data Input
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	rooms := lo.FromPairs(lop.Map(data.Rooms, func(r Room, idx int) lo.Entry[string, string] {
		return lo.Entry[string, string]{
			Key:   r.ID,
			Value: r.Zh.Name,
		}
	}))

	speakers := lo.FromPairs(lop.Map(data.Speakers, func(s Speaker, idx int) lo.Entry[string, string] {
		return lo.Entry[string, string]{
			Key:   s.ID,
			Value: s.Zh.Name,
		}
	}))

	roomSession := lo.GroupBy(
		lop.Map(
			lo.FlatMap[Session, *ent.Session](
				data.Sessions,
				parseInput(speakers, rooms),
			),
			removeSpeakerFromRest,
		),
		func(s *ent.Session) string {
			return s.Room
		},
	)

	roomSession = lo.MapValues(roomSession, func(ss []*ent.Session, _ string) []*ent.Session {
		ss = lo.Reduce(
			sortByStart(ss),
			mergeSameTitle,
			[]*ent.Session{},
		)

		return lop.Map(ss, setNextAndIdx(ss))
	})

	return lo.Flatten(lo.Values(roomSession)), nil
}

func parseTime(timeStr string) int64 {
	t, _ := time.Parse(time.RFC3339, timeStr)
	return t.Unix()
}

// with the rollback error if occurred.
func rollback(tx *ent.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	return err
}

func SaveToDB(ctx context.Context, client *ent.Client, sessions []*ent.Session) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}

	// delete all sessions
	n, err := tx.Session.Delete().Exec(ctx)
	if err != nil {
		return rollback(tx, err)
	}
	log.Printf("delete %d sessions", n)

	// insert new sessions
	for _, s := range sessions {
		_, err := tx.Session.Create().
			SetIdx(s.Idx).
			SetStart(s.Start).
			SetEnd(s.End).
			SetRoom(s.Room).
			SetSessionID(s.SessionID).
			SetNext(s.Next).
			SetTitle(s.Title).
			SetData(s.Data).
			Save(ctx)
		if err != nil {
			return rollback(tx, err)
		}
	}

	return tx.Commit()
}
