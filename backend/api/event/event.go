package event

import (
	"context"
	"io"
	"net/http"

	"backend/ent"
	"backend/internal/token"
	"backend/models/event"
	"backend/util"

	"github.com/danielgtaylor/huma/v2"
)

type Output[T any] struct {
	Body T
}

type BodyCreateEvent struct {
	Name   string `json:"name" example:"SITCON2025" doc:"Event Name"`
	URL    string `json:"url" example:"https://sitcon.org/2025/sessions.json" doc:"URl to sessions.json"`
	Script string `json:"script" doc:"javascript to process sessions.json"`
}

type BodyUpdateScript struct {
	Script string `json:"script" doc:"javascript to process sessions.json"`
}

func Route(api huma.API, t *token.Token) {
	huma.Get(api, "/", func(ctx context.Context, input *struct{}) (*Output[ent.Events], error) {
		e, err := event.GetAll(ctx)
		if err != nil {
			return nil, err
		}

		return &Output[ent.Events]{
			Body: e,
		}, nil
	}, util.APIDesp("Get All Events", "Get all events", "event"))

	huma.Get(api, "/{name}", func(ctx context.Context, input *struct {
		Name string `path:"name" example:"SITCON2025" doc:"Event Name"`
	},
	) (*Output[*ent.Event], error) {
		e, err := event.Get(ctx, input.Name)
		if err != nil {
			return nil, err
		}

		return &Output[*ent.Event]{
			Body: e,
		}, nil
	}, util.APIDesp("Get Event", "Get event by name.", "event"))

	huma.Get(api, "/{name}/session", func(ctx context.Context, input *struct {
		Name string `path:"name" example:"SITCON2025" doc:"Event Name"`
	},
	) (*huma.StreamResponse, error) {
		e, err := event.Get(ctx, input.Name)
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
	}, util.APIDesp("Get Event Sessions", "Get sessions of event.", "event"))

	huma.Post(api, "/", func(ctx context.Context, input *struct {
		Body BodyCreateEvent
	},
	) (*Output[*ent.Event], error) {
		e, err := event.NewEvent(ctx, input.Body.Name, input.Body.URL, input.Body.Script)
		if err != nil {
			return nil, err
		}

		return &Output[*ent.Event]{
			Body: e,
		}, nil
	},
		util.APIDesp("Create Event", "Create a new event.", "event"),
		t.AuthHuma(api),
	)

	huma.Put(api, "/{name}", func(ctx context.Context, input *struct {
		Name string `path:"name" example:"SITCON2025" doc:"Event Name"`
		Body BodyUpdateScript
	},
	) (*Output[*ent.Event], error) {
		e, err := event.UpdateScript(ctx, input.Name, input.Body.Script)
		if err != nil {
			return nil, err
		}

		return &Output[*ent.Event]{
			Body: e,
		}, nil
	},
		util.APIDesp("Update Event Script", "Update event script by name", "event"),
		t.AuthHuma(api),
	)
}
