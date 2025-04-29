package usecase

import "context"

type EventDeleteInput struct {
	Name string `json:"name"`
}

// left empty, for future expansion
type EventDeleteOutput struct{}

func (e *EventImpl) Delete(ctx context.Context, input *EventDeleteInput) (*EventDeleteOutput, error) {
	err := e.event.Delete(ctx, input.Name)
	return nil, err
}
