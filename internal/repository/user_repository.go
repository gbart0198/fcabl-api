package repository

import (
	"context"
)

type UserRepository interface {
	ListUsers(ctx context.Context) ([]ListUsersRow, error)
}

type userRepository struct {
	queries *Queries
}

func NewUserRepository(queries *Queries) UserRepository {
	return &userRepository{
		queries: queries,
	}
}

func (r *userRepository) ListUsers(ctx context.Context) ([]ListUsersRow, error) {
	return r.queries.ListUsers(ctx)
}
