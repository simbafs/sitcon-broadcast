package usecase

import (
	"context"

	"backend/internal/entity"
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

type Counter interface {
	New(name string, init int) (*entity.Counter, error)
	List() []*entity.Counter
	Get(name string) (*entity.Counter, error)

	Start(name string) error
	Stop(name string) error
	SetInit(name string, init int) error
}
