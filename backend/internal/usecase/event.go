package usecase

import "backend/internal/repository"

type EventImpl struct {
	event repository.Event
}

func NewEvent(event repository.Event) *EventImpl {
	return &EventImpl{
		event: event,
	}
}

type EventItem struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Script string `json:"script"`
}
