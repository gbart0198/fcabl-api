// Package service provides business logic for the application
package service

import (
	"context"

	"github.com/gbart/fcabl-api/internal/repository"
)

type UserService interface {
	ListUsers(ctx context.Context, includes []string) ([]repository.ListUsersRow, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) ListUsers(ctx context.Context, includes []string) ([]repository.ListUsersRow, error) {
	return s.repo.ListUsers(ctx)
}
