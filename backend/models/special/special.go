package special

import (
	"context"

	"backend/ent"
	"backend/ent/special"

	m "backend/models"
)

func Read(ctx context.Context, id string) (*ent.Special, error) {
	return m.Client.Special.
		Query().
		Where(special.ID(id)).
		Only(ctx)
}

func ReadAll(ctx context.Context) ([]*ent.Special, error) {
	return m.Client.Special.
		Query().
		All(ctx)
}

func Update(ctx context.Context, id string, data string) (*ent.Special, error) {
	return m.Client.Special.
		UpdateOneID(id).
		SetData(data).
		Save(ctx)
}
