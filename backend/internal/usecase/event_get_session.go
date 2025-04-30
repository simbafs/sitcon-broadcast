package usecase

import (
	"context"
	"fmt"
	"net/http"
)

type EventGetSessionInput struct {
	Name string `json:"name"`
}

type EventGetSessionOutput = http.Response

func (e *EventImpl) GetSession(ctx context.Context, input *EventGetSessionInput) (*EventGetSessionOutput, error) {
	event, err := e.event.Get(ctx, input.Name)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(event.URL())
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get session: %s", resp.Status)
	}

	return resp, nil
}
