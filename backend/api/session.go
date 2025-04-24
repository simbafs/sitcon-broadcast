package api

import (
	"context"

	"backend/ent"
	"backend/models/session"
)

func (h *Handler) GetAllSession(ctx context.Context, input *struct {
	Room string `path:"room" example:"R0" doc:"Room ID"`
},
) (*Output[*ent.Session], error) {
	s, err := h.session.GetCurrent(ctx, input.Room)
	if err != nil {
		return nil, err
	}
	return &Output[*ent.Session]{
		Body: s,
	}, nil
}

func (h *Handler) GetSessionInRoom(ctx context.Context, input *struct {
	Room string `path:"room" example:"R0" doc:"Room ID"`
},
) (*Output[ent.Sessions], error) {
	s, err := h.session.GetAllInRoom(ctx, input.Room)
	if err != nil {
		return nil, err
	}
	return &Output[ent.Sessions]{
		Body: s,
	}, nil
}

func (h *Handler) GetSessionByID(ctx context.Context, input *struct {
	Room string `path:"room" example:"R0" doc:"Room ID"`
	ID   string `path:"id" example:"2d8a5e" doc:"Session ID"`
},
) (*Output[*ent.Session], error) {
	s, err := h.session.Get(ctx, input.Room, input.ID)
	if err != nil {
		return nil, err
	}
	return &Output[*ent.Session]{
		Body: s,
	}, nil
}

type BodyNext struct {
	End int64 `json:"end" example:"1741393800" doc:"End time of current session in unix timestamp in seconds"`
}

func (h *Handler) NextSession(ctx context.Context, input *struct {
	Room string `path:"room" example:"R0" doc:"Room ID"`
	ID   string `path:"id" example:"2d8a5e" doc:"Current session ID"`
	Body BodyNext
},
) (*Output[*ent.Session], error) {
	s, err := h.session.Next(ctx, input.Room, input.ID, input.Body.End, h.send) // TODO: too many arguments
	if err != nil {
		return nil, err
	}
	return &Output[*ent.Session]{
		Body: s,
	}, nil
}

type BodySetSession struct {
	Sessions []session.SessionWithoutID
}

func (h *Handler) SetAllSession(ctx context.Context, input *struct {
	Body BodySetSession
},
) (*Output[string], error) {
	err := h.session.UpdateAll(ctx, input.Body.Sessions)

	return &Output[string]{
		Body: "ok",
	}, err
}
