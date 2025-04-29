package usecase

import "context"

type EventGetInput struct {
	Name string `json:"name"`
}

type EventGetOutput struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Script string `json:"script"`
}

func (g *EventImpl) Get(ctx context.Context, input *EventGetInput) (*EventGetOutput, error) {
	e, err := g.event.Get(ctx, input.Name)
	if err != nil {
		return nil, err
	}

	output := &EventGetOutput{
		Name:   e.Name(),
		URL:    e.URL(),
		Script: e.Script(),
	}

	return output, nil
}
