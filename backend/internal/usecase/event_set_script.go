package usecase

import "context"

type EventSetScriptInput struct {
	Name   string `json:"name"`
	Script string `json:"script"`
}

// left empty, for future expansion
type EventSetScriptOutput struct{}

func (e *Event) Execute(ctx context.Context, input *EventSetScriptInput) (*EventSetScriptOutput, error) {
	err := e.event.Update(ctx, input.Name, "", input.Script)
	return nil, err
}
