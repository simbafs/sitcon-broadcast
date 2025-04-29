package usecase

import "context"

type EventCreateInput struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EventCreateOutput struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Script string `json:"script"`
}

func (e *EventImpl) Create(ctx context.Context, input *EventCreateInput) (*EventCreateOutput, error) {
	ev, err := e.event.Create(ctx, input.Name, input.URL, "")
	if err != nil {
		return nil, err
	}

	// TODO: fetch default script from github

	output := &EventCreateOutput{
		Name:   ev.Name(),
		URL:    ev.URL(),
		Script: ev.Script(),
	}

	return output, nil
}
