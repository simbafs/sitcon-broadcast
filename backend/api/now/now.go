package now

import (
	"context"

	"backend/internal/token"
	"backend/models/now"
	"backend/util"

	"github.com/danielgtaylor/huma/v2"
)

type NowOutput struct {
	Body now.Now `doc:"current time."`
}

type BodySetNow struct {
	Now int64 `json:"now" example:"1741393800" doc:"Current time in seconds since epoch."`
}

func Route(api huma.API, t *token.Token) {
	huma.Get(api, "/", func(ctx context.Context, input *struct{}) (*NowOutput, error) {
		n := now.GetNow()
		return &NowOutput{
			Body: n,
		}, nil
	}, util.APIDesp("Get Current Time", "Get the current time in unix timestamp in seconds", "now"))

	huma.Post(api, "/", func(ctx context.Context, input *struct {
		Body BodySetNow
	},
	) (*NowOutput, error) {
		now.SetNow(input.Body.Now)
		return &NowOutput{
			Body: now.GetNow(),
		}, nil
	},
		util.APIDesp("Set Current Time", "Set the current time in unix timestamp in seconds", "now"),
		t.AuthHuma(api),
	)

	huma.Delete(api, "/", func(ctx context.Context, input *struct{}) (*NowOutput, error) {
		now.ResetNow()
		return &NowOutput{
			Body: now.GetNow(),
		}, nil
	},
		util.APIDesp("Reset Current Time", "Reset the current time to the actual current time.", "now"),
		t.AuthHuma(api),
	)
}
