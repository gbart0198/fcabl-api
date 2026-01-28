package repository

import (
	"context"
)

type GameRepository interface {
	ListGames(ctx context.Context) ([]Game, error)
}

type gameRepository struct {
	queries *Queries
}

func NewGameRepository(queries *Queries) GameRepository {
	return &gameRepository{queries: queries}
}

func (r *gameRepository) ListGames(ctx context.Context) ([]Game, error) {
	return r.queries.ListGames(ctx)
}
