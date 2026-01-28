package service

import (
	"context"

	"github.com/gbart/fcabl-api/internal/repository"
)

type PlayerService interface {
	ListPlayers(ctx context.Context) ([]repository.Player, error)
}

type playerService struct {
	repository repository.PlayerRepository
}

func NewPlayerService(repository repository.PlayerRepository) PlayerService {
	return &playerService{repository: repository}
}

func (s *playerService) ListPlayers(ctx context.Context) ([]repository.Player, error) {
	return s.repository.ListPlayers(ctx)
}
