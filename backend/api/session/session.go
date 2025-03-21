package session

import (
	"context"

	"backend/ent"
	"backend/internal/logger"
	"backend/internal/token"
	"backend/models/session"

	"github.com/danielgtaylor/huma/v2"
)

var log = logger.New("api/session")

type Output[T any] struct {
	Body T
}

type BodyNext struct {
	End int64 `json:"end" example:"1741393800" doc:"End time of current session in unix timestamp in seconds"`
}

type BodySetSession struct {
	Sessions []session.SessionWithoutID
}

func Route(api huma.API, t *token.Token) {
	huma.Get(api, "/{room}", func(ctx context.Context, input *struct {
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
	}, func(op *huma.Operation) {
		op.Tags = []string{"session"}
		op.Summary = "Get Current Session in Room"
	})

	huma.Get(api, "/{room}/all", func(ctx context.Context, input *struct {
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
	}, func(op *huma.Operation) {
		op.Tags = []string{"session"}
		op.Summary = "Get All Sessions in Room"
		op.Description = "Get all sessions in a room."
	})

	huma.Get(api, "/{room}/{id}", func(ctx context.Context, input *struct {
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
	}, func(op *huma.Operation) {
		op.Tags = []string{"session"}
		op.Summary = "Get Session by ID in Room"
		op.Description = "Get a session by its ID in a room."
	})

	// trigger `next` on session with ID in Room, return the next session
	huma.Post(api, "/{room}/{id}", func(ctx context.Context, input *struct {
		Room string `path:"room" example:"R0" doc:"Room ID"`
		ID   string `path:"id" example:"2d8a5e" doc:"Current session ID"`
		Body BodyNext
	},
	) (*Output[*ent.Session], error) {
		s, err := session.Next(ctx, input.Room, input.ID, input.Body.End)
		if err != nil {
			return nil, err
		}
		return &Output[*ent.Session]{
			Body: s,
		}, nil
	}, func(op *huma.Operation) {
		op.Tags = []string{"session"}
		op.Summary = "Set End Time of Session"
		op.Description = "Set the end time of the current session and start time of the next session."
		t.AuthHuma(api, op)
	})

	huma.Put(api, "/", func(ctx context.Context, input *struct {
		Body BodySetSession
	},
	) (*Output[string], error) {
		err := session.UpdateAll(ctx, input.Body.Sessions)

		return &Output[string]{
			Body: "ok",
		}, err
	}, func(op *huma.Operation) {
		op.Tags = []string{"session"}
		op.Summary = "Set Sessions"
		op.Description = "Clear all sessions abd set new sessions in database. Note that this API will not check if the session is valid."
		t.AuthHuma(api, op)
	})
}
