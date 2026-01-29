package repository

import (
	"context"

	"github.com/gbart/fcabl-api/internal/models"
)

type UserRepository interface {
	ListUsers(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, userID int64) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	CreateUser(ctx context.Context, user models.CreateUserRequest) (int64, error)
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

func (r *userRepository) GetUserByID(ctx context.Context, userID int64) (User, error) {
	userRow, err := r.queries.GetUserById(ctx, userID)
	if err != nil {
		return User{}, err
	}

	return userRow.ToUser(), nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	userRow, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return User{}, err
	}

	return userRow.ToUser(), nil
}

func (r *userRepository) CreateUser(ctx context.Context, user models.CreateUserRequest) (int64, error) {
	userID, err := r.queries.CreateUser(ctx, CreateUserParams{
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Role:         user.Role,
		PasswordHash: user.PasswordHash,
		PhoneNumber:  user.PhoneNumber,
	})

	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (u *GetUserByEmailRow) ToUser() User {
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

func (u *GetUserByIdRow) ToUser() User {
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
