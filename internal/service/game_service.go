package service

import (
	"context"

	"github.com/gbart/fcabl-api/internal/repository"
)

type GameService interface {
	ListGames(ctx context.Context, includes []string) ([]repository.Game, error)
}

type gameService struct {
	repo repository.GameRepository
}

func NewGameService(repo repository.GameRepository) GameService {
	return &gameService{repo: repo}
}

func (s *gameService) ListGames(ctx context.Context, includes []string) ([]repository.Game, error) {
	return s.repo.ListGames(ctx)
}
