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
		User:       NewUserService(repo),
		Game:       NewGameService(repo),
		Team:       NewTeamService(repo),
		Player:     NewPlayerService(repo),
		GameDetail: NewGameDetailService(repo),
		Payment:    NewPaymentService(repo),
	}
}
