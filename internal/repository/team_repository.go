package repository

import "context"

type TeamRepository interface {
	ListTeams(ctx context.Context) ([]Team, error)
}

type teamRepository struct {
	queries *Queries
}

func NewTeamRepository(queries *Queries) TeamRepository {
	return &teamRepository{queries: queries}
}

func (r *teamRepository) ListTeams(ctx context.Context) ([]Team, error) {
	return r.queries.ListTeams(ctx)
}
