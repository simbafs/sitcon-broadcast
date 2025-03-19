package now

import (
	"context"

	"backend/internal/token"
	"backend/models/now"

	"github.com/danielgtaylor/huma/v2"
)

type NowOutput struct {
	Body now.Now `doc:"current time."`
}

func Route(api huma.API, t *token.Token) {
	huma.Get(api, "/", func(ctx context.Context, input *struct{}) (*NowOutput, error) {
		n := now.GetNow()
		return &NowOutput{
			Body: n,
		}, nil
	}, func(op *huma.Operation) {
		op.Tags = []string{"now"}
		op.Summary = "Get Current Time"
		op.Description = "Get the current time in unix timestamp in seconds."
	})

	huma.Post(api, "/", func(ctx context.Context, input *struct {
		Body struct {
			Now int64 `json:"now" example:"1741393800" doc:"Current time in seconds since epoch."`
			// Token http.Cookie `cookie:"token"`
		}
	},
	) (*NowOutput, error) {
		now.SetNow(input.Body.Now)
		return &NowOutput{
			Body: now.GetNow(),
		}, nil
	}, func(op *huma.Operation) {
		op.Tags = []string{"now"}
		op.Summary = "Set Current Time"
		op.Description = "Set the current time in unix timestamp in seconds."
		t.AuthHuma(api, op)
	})

	huma.Delete(api, "/", func(ctx context.Context, input *struct{}) (*NowOutput, error) {
		now.ResetNow()
		return &NowOutput{
			Body: now.GetNow(),
		}, nil
	}, func(op *huma.Operation) {
		op.Tags = []string{"now"}
		op.Summary = "Reset Current Time"
		op.Description = "Reset the current time to the actual current time."
		t.AuthHuma(api, op)
	})
}
