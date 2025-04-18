package util

import "github.com/danielgtaylor/huma/v2"

func apiDesp(summary string, description string, tags []string) func(op *huma.Operation) {
	return func(op *huma.Operation) {
		op.Summary = summary
		op.Description = description
		op.Tags = tags
	}
}
