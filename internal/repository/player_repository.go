package repository

import (
	"context"
)

type PlayerRepository interface {
	ListPlayers(ctx context.Context) ([]Player, error)
}

type playerRepository struct {
	queries *Queries
}

func NewPlayerRepository(queries *Queries) PlayerRepository {
	return &playerRepository{queries: queries}
}

func (r *playerRepository) ListPlayers(ctx context.Context) ([]Player, error) {
	return r.queries.ListPlayers(ctx)
}
