package delivery

import (
	"context"

	"backend/internal/usecase"
)

var _ EventUsecase = &usecase.Event{}

type EventUsecase interface {
	Create(ctx context.Context, input *usecase.EventCreateInput) (*usecase.EventCreateOutput, error)
	Delete(ctx context.Context, input *usecase.EventDeleteInput) (*usecase.EventDeleteOutput, error)
	Execute(ctx context.Context, input *usecase.EventSetScriptInput) (*usecase.EventSetScriptOutput, error)
	Get(ctx context.Context, input *usecase.EventGetInput) (*usecase.EventGetOutput, error)
	List(ctx context.Context, input *usecase.EventListInput) (*usecase.EventListOutput, error)
}
