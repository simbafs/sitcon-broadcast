package admin

import (
	"context"

	"backend/internal/refresh"
	"backend/internal/token"
	"backend/models"

	"github.com/danielgtaylor/huma/v2"
)

func Route(api huma.API, t *token.Token) {
	huma.Post(api, "/refresh/db", func(ctx context.Context, input *struct {
		Body struct {
			URL string `json:"url" example:"https://sitcon.org/2025/sessions.json" doc:"URL of sessions.json"`
		}
	},
	) (*struct{}, error) {
		ss, err := refresh.FromURL(input.Body.URL)
		if err != nil {
			return nil, err
		}

		return nil, refresh.SaveToDB(ctx, models.Client, ss)
	}, func(op *huma.Operation) {
		op.Tags = []string{"admin"}
		op.Summary = "Refresh Database"
		op.Description = "Fetch latest sessions.json and parse, then write to database. Note that this will overwrite all current records in database"
		t.AuthHuma(api, op)
	})
}
