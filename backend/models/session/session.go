package session

import (
	"backend/models/now"
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
)

//go:embed sessions.json
var file []byte
var Data *DataType

func init() {
	data, err := GetSessions(bytes.NewReader(file))
	if err != nil {
		panic(err)
	}

	Data = data
}

type Room []SessionItem

func (r Room) IsOverlap(idx int, start now.Time, end now.Time) bool {
	for i, s := range r {
		if i == idx {
			continue
		}
		if (start >= s.Start && start < s.End) || (end > s.Start && end <= s.End) {
			return true
		}
	}

	return false
}

func (r Room) GetNow() (SessionItem, bool) {
	for _, s := range r {
		if s.Start > now.GetNow() {
			return s, true
		}
	}

	return SessionItem{}, false
}

type Rooms map[string]Room

func (r Rooms) Get(room string, index int) (SessionItem, bool) {
	if room, ok := r[room]; ok && index >= 0 && index < len(room) {
		return room[index], true
	}

	return SessionItem{}, false
}

type DataType struct {
	Rooms  Rooms `json:"sessions"`
	NextID int   `json:"nextID"`
}

func (d *DataType) GetNextID() string {
	d.NextID++
	return fmt.Sprintf("id-%d", d.NextID-1)
}

type SessionItem struct {
	ID        string   `json:"id"`
	Title     string   `json:"title"`
	Type      string   `json:"type"`
	Speakers  []string `json:"speakers"`
	Room      string   `json:"room"`
	Broadcast []string `json:"broadcast"`
	Start     now.Time `json:"start"`
	End       now.Time `json:"end"`
}

func GetSessions(file io.Reader) (*DataType, error) {
	var data DataType
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
