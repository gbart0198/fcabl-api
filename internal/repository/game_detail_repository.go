package repository

import (
	"context"
)

type GameDetailRepository interface {
	ListGameDetails(ctx context.Context) ([]GameDetail, error)
}

type gameDetailRepository struct {
	queries *Queries
}

func NewGameDetailRepository(queries *Queries) GameDetailRepository {
	return &gameDetailRepository{queries: queries}
}

func (r *gameDetailRepository) ListGameDetails(ctx context.Context) ([]GameDetail, error) {
	return r.queries.ListGameDetails(ctx)
}
