package api

import (
	"context"

	"backend/models/now"
)

func (h *Handler) GetNow(ctx context.Context, input *struct{}) (*Output[now.Now], error) {
	n := now.GetNow()
	return &Output[now.Now]{
		Body: n,
	}, nil
}

type BodySetNow struct {
	Now int64 `json:"now" example:"1741393800" doc:"Current time in seconds since epoch."`
}

func (h *Handler) SetNow(ctx context.Context, input *struct {
	Body BodySetNow
},
) (*Output[now.Now], error) {
	now.SetNow(input.Body.Now, h.send)

	return &Output[now.Now]{
		Body: now.GetNow(),
	}, nil
}

func (h *Handler) ResetNow(ctx context.Context, input *struct{}) (*Output[now.Now], error) {
	now.ResetNow(h.send)

	return &Output[now.Now]{
		Body: now.GetNow(),
	}, nil
}
