// Package service provides business logic for the application
package service

import (
	"context"

	"github.com/gbart/fcabl-api/internal/models"
	"github.com/gbart/fcabl-api/internal/repository"
)

type UserService interface {
	ListUsers(ctx context.Context, includes []string) ([]repository.User, error)
	GetUserByID(ctx context.Context, userID int64) (repository.User, error)
	GetUserByEmail(ctx context.Context, email string) (repository.User, error)
	CreateUser(ctx context.Context, user models.CreateUserRequest) (int64, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) ListUsers(ctx context.Context, includes []string) ([]repository.User, error) {
	return s.repo.ListUsers(ctx)
}

func (s *userService) GetUserByID(ctx context.Context, userID int64) (repository.User, error) {
	return s.repo.GetUserByID(ctx, userID)
}
func (s *userService) GetUserByEmail(ctx context.Context, email string) (repository.User, error) {
	return s.GetUserByEmail(ctx, email)
}
func (s *userService) CreateUser(ctx context.Context, user models.CreateUserRequest) (int64, error) {
	return s.CreateUser(ctx, user)
}
