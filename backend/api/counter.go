package api

import (
	"context"
	"net/http"

	"backend/models/counter"

	"github.com/danielgtaylor/huma/v2"
)

func (h *Handler) GetAllCounter(ctx context.Context, input *struct{}) (*Output[counter.CounterGroup], error) {
	return &Output[counter.CounterGroup]{
		Body: h.counters,
	}, nil
}

func (h *Handler) GetCounter(ctx context.Context, input *struct {
	Name string `path:"name" doc:"counter name" example:"R0"`
},
) (*Output[counter.Counter], error) {
	c := h.counters.Get(input.Name)
	if c == nil {
		return nil, huma.NewError(http.StatusNotFound, "counter not found")
	}

	return &Output[counter.Counter]{
		Body: *c,
	}, nil
}

type CreateInit struct {
	Init int `json:"init" doc:"initial value" example:"10"`
}

func (h *Handler) SetCounterInit(ctx context.Context, input *struct {
	Name string     `path:"name" doc:"counter name" example:"R0"`
	Body CreateInit `json:"body" doc:"counter init value"`
},
) (*Output[counter.Counter], error) {
	c := h.counters.Get(input.Name)
	if c == nil {
		return nil, huma.NewError(http.StatusNotFound, "counter not found")
	}

	c.SetInit(input.Body.Init)

	return &Output[counter.Counter]{
		Body: *c,
	}, nil
}

func (h *Handler) CounterStart(ctx context.Context, input *struct {
	Name string `path:"name" doc:"counter name" example:"R0"`
},
) (*Output[counter.Counter], error) {
	c := h.counters.Get(input.Name)
	if c == nil {
		return nil, huma.NewError(http.StatusNotFound, "counter not found")
	}

	go c.Start()

	return &Output[counter.Counter]{
		Body: *c,
	}, nil
}

func (h *Handler) CounterStop(ctx context.Context, input *struct {
	Name string `path:"name" doc:"counter name" example:"R0"`
},
) (*Output[counter.Counter], error) {
	c := h.counters.Get(input.Name)
	if c == nil {
		return nil, huma.NewError(http.StatusNotFound, "counter not found")
	}

	c.Stop()

	return &Output[counter.Counter]{
		Body: *c,
	}, nil
}

func (h *Handler) CounterReset(ctx context.Context, input *struct {
	Name string `path:"name" doc:"counter name" example:"R0"`
},
) (*Output[counter.Counter], error) {
	c := h.counters.Get(input.Name)
	if c == nil {
		return nil, huma.NewError(http.StatusNotFound, "counter not found")
	}

	c.Reset()

	return &Output[counter.Counter]{
		Body: *c,
	}, nil
}
