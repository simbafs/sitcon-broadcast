package entity

import (
	"errors"
	"io"
	"net/http"
)

type Event struct {
	name   string
	url    string
	script string
}

func NewEvent(name, url, script string) *Event {
	return &Event{
		name:   name,
		url:    url,
		script: script,
	}
}

func (e *Event) Name() string {
	return e.name
}

func (e *Event) URL() string {
	return e.url
}

func (e *Event) Script() string {
	return e.script
}

func (e *Event) SetURL(url string) {
	e.url = url
}

func (e *Event) SetScript(script string) {
	e.script = script
}

var ErrCannotGetSessions = errors.New("cannot get sessions")

func (e *Event) GetSessions() (string, error) {
	resp, err := http.Get(e.url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", ErrCannotGetSessions
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
