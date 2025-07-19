package graph

import (
	"context"

	"github.com/kprueger/circumitus-api/graph/model"
	"github.com/kprueger/circumitus-api/internal/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r *mutationResolver) Chip(ctx context.Context, id string) (*model.Chip, error) {
	var chip model.Chip
	err := db.Pool.QueryRow(ctx, "INSERT INTO chip (id) VALUES ($1) RETURNING id", id).Scan(&chip.ID)
	if err != nil {
		return nil, err
	}
	return &chip, nil
}
