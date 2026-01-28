package service

import (
	"context"

	"github.com/gbart/fcabl-api/internal/repository"
)

type TeamService interface {
	ListTeams(ctx context.Context, includes []string) ([]repository.Team, error)
}

type teamService struct {
	repo repository.TeamRepository
}

func NewTeamService(repo repository.TeamRepository) TeamService {
	return &teamService{repo: repo}
}

func (s *teamService) ListTeams(ctx context.Context, includes []string) ([]repository.Team, error) {
	return s.repo.ListTeams(ctx)
}
