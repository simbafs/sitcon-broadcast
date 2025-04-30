package usecase

import (
	"context"
)

type Event interface {
	Create(ctx context.Context, input *EventCreateInput) (*EventCreateOutput, error)
	Delete(ctx context.Context, input *EventDeleteInput) (*EventDeleteOutput, error)
	Execute(ctx context.Context, input *EventSetScriptInput) (*EventSetScriptOutput, error)
	Get(ctx context.Context, input *EventGetInput) (*EventGetOutput, error)
	List(ctx context.Context, input *EventListInput) (*EventListOutput, error)
	GetSession(ctx context.Context, input *EventGetSessionInput) (*EventGetSessionOutput, error)
}

type Now interface {
	Get() *NowOutput
	Set(input *NowInput) *NowOutput
	Reset() *NowOutput
}
