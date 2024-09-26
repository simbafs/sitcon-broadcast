package session

import (
	"backend/models/now"
	"encoding/json"
	"fmt"
	"io"
)

type Room []SessionItem

func (r Room) IsOverlap(start now.Time, end now.Time) bool {
	for _, s := range r {
		if (start >= s.Start && start < s.End) || (end > s.Start && end <= s.End) {
			return true
		}
	}

	return false
}

type (
	Sessions map[string]Room
)

func (s Sessions) Get(room string, index int) (SessionItem, bool) {
	if items, ok := s[room]; ok && index >= 0 && index < len(items) {
		return items[index], true
	}

	return SessionItem{}, false
}

type IDMap map[string]IDMapItem

func (m IDMap) Get(id string) (string, int, bool) {
	if item, ok := m[id]; ok {
		return item.Room, item.Index, true
	}

	return "", -1, false
}

type Data struct {
	Sessions Sessions `json:"sessions"`
	IDMap    IDMap    `json:"idMap"`
	NextID   int      `json:"nextID"`
}

func (d *Data) GetNextID() string {
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

type IDMapItem struct {
	Room  string `json:"room"`
	Index int    `json:"index"`
}

func GetSessions(file io.Reader) (*Data, error) {
	var data Data
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
