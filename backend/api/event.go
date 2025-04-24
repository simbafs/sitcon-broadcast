package api

import (
	"context"
	"io"
	"net/http"

	"backend/ent"

	"github.com/danielgtaylor/huma/v2"
)

func (h *Handler) GetAllEvent(ctx context.Context, input *struct{}) (*Output[ent.Events], error) {
	e, err := h.event.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return &Output[ent.Events]{
		Body: e,
	}, nil
}

func (h *Handler) GetEvent(ctx context.Context, input *struct {
	Name string `path:"name" example:"SITCON2025" doc:"Event Name"`
},
) (*Output[*ent.Event], error) {
	e, err := h.event.Get(ctx, input.Name)
	if err != nil {
		return nil, err
	}

	return &Output[*ent.Event]{
		Body: e,
	}, nil
}

func (h *Handler) GetEventSession(ctx context.Context, input *struct {
	Name string `path:"name" example:"SITCON2025" doc:"Event Name"`
},
) (*huma.StreamResponse, error) {
	e, err := h.event.Get(ctx, input.Name)
	if err != nil {
		return nil, err
	}

	res, err := http.Get(e.URL)
	if err != nil {
		return nil, err
	}

	return &huma.StreamResponse{
		Body: func(ctx huma.Context) {
			ctx.SetHeader("Content-Type", "application/json")
			w := ctx.BodyWriter()
			io.Copy(w, res.Body)
		},
	}, nil
}

type BodyCreateEvent struct {
	Name   string `json:"name" example:"SITCON2025" doc:"Event Name"`
	URL    string `json:"url" example:"https://sitcon.org/2025/sessions.json" doc:"URl to sessions.json"`
	Script string `json:"script" doc:"javascript to process sessions.json"`
}

func (h *Handler) CreateEvent(ctx context.Context, input *struct {
	Body BodyCreateEvent
},
) (*Output[*ent.Event], error) {
	e, err := h.event.NewEvent(ctx, input.Body.Name, input.Body.URL, input.Body.Script)
	if err != nil {
		return nil, err
	}

	return &Output[*ent.Event]{
		Body: e,
	}, nil
}

type BodyUpdateScript struct {
	Script string `json:"script" doc:"javascript to process sessions.json"`
}

func (h *Handler) UpdateEventScript(ctx context.Context, input *struct {
	Name string `path:"name" example:"SITCON2025" doc:"Event Name"`
	Body BodyUpdateScript
},
) (*Output[*ent.Event], error) {
	e, err := h.event.UpdateScript(ctx, input.Name, input.Body.Script)
	if err != nil {
		return nil, err
	}

	return &Output[*ent.Event]{
		Body: e,
	}, nil
}
