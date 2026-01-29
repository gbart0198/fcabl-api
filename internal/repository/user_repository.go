package repository

import (
	"context"
)

type UserRepository interface {
	ListUsers(ctx context.Context) ([]User, error)
}

type userRepository struct {
	queries *Queries
}

func NewUserRepository(queries *Queries) UserRepository {
	return &userRepository{
		queries: queries,
	}
}

func (r *userRepository) ListUsers(ctx context.Context) ([]User, error) {
	userRows, err := r.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]User, len(userRows))
	for i, userRow := range userRows {
		users[i] = userRow.ToUser()
	}
	return users, nil
}

func (u *ListUsersRow) ToUser() User {
	return User{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
