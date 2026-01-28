package service

import (
	"context"

	"github.com/gbart/fcabl-api/internal/repository"
)

type GameDetailService interface {
	ListGameDetails(ctx context.Context) ([]repository.GameDetail, error)
}

type gameDetailService struct {
	repository repository.GameDetailRepository
}

func NewGameDetailService(repository repository.GameDetailRepository) GameDetailService {
	return &gameDetailService{repository: repository}
}

func (s *gameDetailService) ListGameDetails(ctx context.Context) ([]repository.GameDetail, error) {
	return s.repository.ListGameDetails(ctx)
}
