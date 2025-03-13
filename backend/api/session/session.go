package session

import (
	"context"

	"backend/ent"
	"backend/models/session"

	"github.com/danielgtaylor/huma/v2"
)

type Output[T any] struct {
	Body T
}

func Route(api huma.API) {
	huma.Get(api, "/api/session/{room}", func(ctx context.Context, input *struct {
		Room string `path:"room" example:"R0" doc:"Room ID"`
	},
	) (*Output[*ent.Session], error) {
		s, err := session.GetCurrent(ctx, input.Room)
		if err != nil {
			return nil, err
		}
		return &Output[*ent.Session]{
			Body: s,
		}, nil
	})

	huma.Get(api, "/api/session/{room}/all", func(ctx context.Context, input *struct {
		Room string `path:"room" example:"R0" doc:"Room ID"`
	},
	) (*Output[ent.Sessions], error) {
		s, err := session.GetAllInRoom(ctx, input.Room)
		if err != nil {
			return nil, err
		}
		return &Output[ent.Sessions]{
			Body: s,
		}, nil
	})

	huma.Get(api, "/api/session/{room}/{id}", func(ctx context.Context, input *struct {
		Room string `path:"room" example:"R0" doc:"Room ID"`
		ID   string `path:"id" example:"2d8a5e" doc:"Session ID"`
	},
	) (*Output[*ent.Session], error) {
		s, err := session.Get(ctx, input.Room, input.ID)
		if err != nil {
			return nil, err
		}
		return &Output[*ent.Session]{
			Body: s,
		}, nil
	})

	// trigger `next` on session with ID in Room, return the next session
	huma.Post(api, "/api/session/{room}/{id}", func(ctx context.Context, input *struct {
		Room string `path:"room" example:"R0" doc:"Room ID"`
		ID   string `path:"id" example:"2d8a5e" doc:"Current session ID"`
		Body struct {
			End int64 `json:"end" example:"1741393800" doc:"End time of current session in unix timestamp in seconds"`
		}
	},
	) (*Output[*ent.Session], error) {
		s, err := session.SetEnd(ctx, input.Room, input.ID, input.Body.End)
		if err != nil {
			return nil, err
		}
		return &Output[*ent.Session]{
			Body: s,
		}, nil
	})
}
