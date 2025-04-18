package counter

import (
	"context"
	"net/http"

	"backend/internal/token"
	"backend/models/counter"
	"backend/util"

	"github.com/danielgtaylor/huma/v2"
)

var Counters = counter.NewGroup([]string{
	"R0",
	"R1",
	"R2",
	"R3",
	"S",
})

type Output[T any] struct {
	Body T `doc:"response body"`
}

func Route(api huma.API, t *token.Token) {
	huma.Get(api, "/", func(ctx context.Context, input *struct{}) (*Output[counter.CounterGroup], error) {
		return &Output[counter.CounterGroup]{
			Body: Counters,
		}, nil
	}, util.APIDesp("Get All Counter", "Get all counters.", "counter"))

	huma.Get(api, "/{name}", func(ctx context.Context, input *struct {
		Name string `path:"name" doc:"counter name" example:"R0"`
	},
	) (*Output[counter.Counter], error) {
		c := Counters.Get(input.Name)
		if c == nil {
			return nil, huma.NewError(http.StatusNotFound, "counter not found")
		}

		return &Output[counter.Counter]{
			Body: *c,
		}, nil
	}, util.APIDesp("Get Counter", "Get a counter by name.", "counter"))

	huma.Put(api, "/{name}", func(ctx context.Context, input *struct {
		Name string `path:"name" doc:"counter name" example:"R0"`
		Body struct {
			Init int `json:"init" doc:"initial value" example:"10"`
		}
	},
	) (*Output[counter.Counter], error) {
		c := Counters.Get(input.Name)
		if c == nil {
			return nil, huma.NewError(http.StatusNotFound, "counter not found")
		}

		c.Set(input.Body.Init)

		return &Output[counter.Counter]{
			Body: *c,
		}, nil
	},
		util.APIDesp("Set Init Value", "Set the initial value of a counter. It will reset the counter.", "counter"),
		t.AuthHuma(api),
	)

	huma.Put(api, "/{name}/start", func(ctx context.Context, input *struct {
		Name string `path:"name" doc:"counter name" example:"R0"`
	},
	) (*Output[counter.Counter], error) {
		c := Counters.Get(input.Name)
		if c == nil {
			return nil, huma.NewError(http.StatusNotFound, "counter not found")
		}

		c.Start()

		return &Output[counter.Counter]{
			Body: *c,
		}, nil
	},
		util.APIDesp("Start Counter", "Start a counter. It will reset the counter depend on the state.", "counter"),
		t.AuthHuma(api),
	)

	huma.Put(api, "/{name}/stop", func(ctx context.Context, input *struct {
		Name string `path:"name" doc:"counter name" example:"R0"`
	},
	) (*Output[counter.Counter], error) {
		c := Counters.Get(input.Name)
		if c == nil {
			return nil, huma.NewError(http.StatusNotFound, "counter not found")
		}

		c.Stop()

		return &Output[counter.Counter]{
			Body: *c,
		}, nil
	},
		util.APIDesp("Stop Counter", "Stop a counter. It will reset the counter when start it again.", "counter"),
		t.AuthHuma(api),
	)

	huma.Put(api, "/{name}/pause", func(ctx context.Context, input *struct {
		Name string `path:"name" doc:"counter name" example:"R0"`
	},
	) (*Output[counter.Counter], error) {
		c := Counters.Get(input.Name)
		if c == nil {
			return nil, huma.NewError(http.StatusNotFound, "counter not found")
		}

		c.Pause()

		return &Output[counter.Counter]{
			Body: *c,
		}, nil
	},
		util.APIDesp("Pause Counter", "Pause a counter. It will not reset the counter when start it again.", "counter"),
		t.AuthHuma(api),
	)
}
