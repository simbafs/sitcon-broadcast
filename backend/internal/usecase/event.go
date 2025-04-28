package usecase

type Event struct {
	event EventRepository
}

func NewEvent(event EventRepository) *Event {
	return &Event{
		event: event,
	}
}

type EventItem struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Script string `json:"script"`
}
