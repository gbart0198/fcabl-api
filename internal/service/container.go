package service

import (
	"github.com/gbart/fcabl-api/internal/config"
	"github.com/gbart/fcabl-api/internal/repository"
)

type Container struct {
	User       UserService
	Game       GameService
	Team       TeamService
	Player     PlayerService
	GameDetail GameDetailService
	Payment    PaymentService
}

func NewContainer(repo *repository.Queries, cfg *config.Config) *Container {
	return &Container{
		User:       NewUserService(repository.NewUserRepository(repo)),
		Game:       NewGameService(repository.NewGameRepository(repo)),
		Team:       NewTeamService(repository.NewTeamRepository(repo)),
		Player:     NewPlayerService(repository.NewPlayerRepository(repo)),
		GameDetail: NewGameDetailService(repository.NewGameDetailRepository(repo)),
		Payment:    NewPaymentService(repository.NewPaymentRepository(repo)),
	}
}
