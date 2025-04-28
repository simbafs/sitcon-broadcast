package usecase

import "context"

// left empty, for future expansion
type EventListInput struct{}

type EventListOutput struct {
	Events []EventItem `json:"events"`
}

func (e *Event) List(ctx context.Context, input *EventListInput) (*EventListOutput, error) {
	events, err := e.event.List(ctx)
	if err != nil {
		return nil, err
	}

	output := &EventListOutput{
		Events: make([]EventItem, len(events)),
	}

	for i, e := range events {
		output.Events[i] = EventItem{
			Name:   e.Name(),
			URL:    e.URL(),
			Script: e.Script(),
		}
	}

	return output, nil
}
