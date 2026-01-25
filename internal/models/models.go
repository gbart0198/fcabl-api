// Package models provides structs to be used for API request binding and validation.
package models

import (
	"github.com/gbart/fcabl-api/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateUserRequest struct {
	Email        string `json:"email" binding:"required"`
	PhoneNumber  string `json:"phoneNumber" binding:"required"`
	PasswordHash string `json:"passwordHash" binding:"required"`
	FirstName    string `json:"firstName" binding:"required"`
	LastName     string `json:"lastName" binding:"required"`
	Role         string `json:"role" binding:"required"`
}

func (rq *CreateUserRequest) IntoDBModel() repository.CreateUserParams {
	return repository.CreateUserParams{
		Email:        rq.Email,
		PhoneNumber:  rq.PhoneNumber,
		PasswordHash: rq.PasswordHash,
		FirstName:    rq.FirstName,
		LastName:     rq.LastName,
		Role:         rq.Role,
	}
}

type UpdateUserRequest struct {
	Email       string           `json:"email" binding:"required"`
	PhoneNumber string           `json:"phoneNumber" binding:"required"`
	FirstName   string           `json:"firstName" binding:"required"`
	LastName    string           `json:"lastName" binding:"required"`
	Role        string           `json:"role" binding:"required"`
	UpdatedAt   pgtype.Timestamp `json:"updatedAt" binding:"required"`
	ID          int64            `json:"id" binding:"required"`
}

func (rq *UpdateUserRequest) IntoDBModel() repository.UpdateUserParams {
	return repository.UpdateUserParams{
		Email:       rq.Email,
		PhoneNumber: rq.PhoneNumber,
		FirstName:   rq.FirstName,
		LastName:    rq.LastName,
		Role:        rq.Role,
		UpdatedAt:   rq.UpdatedAt,
		ID:          rq.ID,
	}
}

type CreatePlayerRequest struct {
	UserID       int64       `json:"userId" binding:"required"`
	TeamID       pgtype.Int8 `json:"teamId" binding:"required"`
	FeeRemainder int32       `json:"feeRemainder" binding:"required"`
	JerseyNumber pgtype.Int4 `json:"jerseyNumber" binding:"required"`
}

func (rq *CreatePlayerRequest) IntoDBModel() repository.CreatePlayerParams {
	return repository.CreatePlayerParams{
		UserID:       rq.UserID,
		TeamID:       rq.TeamID,
		FeeRemainder: rq.FeeRemainder,
		JerseyNumber: rq.JerseyNumber,
	}
}

type UpdatePlayerRequest struct {
	ID           int64       `json:"id" binding:"required"`
	TeamID       pgtype.Int8 `json:"teamId" binding:"required"`
	FeeRemainder int32       `json:"registrationFeeDue" binding:"required"`
	JerseyNumber pgtype.Int4 `json:"jerseyNumber" binding:"required"`
}

func (rq *UpdatePlayerRequest) IntoDBModel() repository.UpdatePlayerParams {
	return repository.UpdatePlayerParams{
		ID:           rq.ID,
		TeamID:       rq.TeamID,
		FeeRemainder: rq.FeeRemainder,
		JerseyNumber: rq.JerseyNumber,
	}
}

type UpdatePlayerTeamRequest struct {
	ID     int64       `json:"id" binding:"required"`
	TeamID pgtype.Int8 `json:"teamId" binding:"required"`
}

func (rq *UpdatePlayerTeamRequest) IntoDBModel() repository.UpdatePlayerTeamParams {
	return repository.UpdatePlayerTeamParams{
		ID:     rq.ID,
		TeamID: rq.TeamID,
	}
}

type UpdatePlayerRegistrationFeeRequest struct {
	ID           int64 `json:"id" binding:"required"`
	FeeRemainder int32 `json:"feeRemainder" binding:"required"`
}

func (rq *UpdatePlayerRegistrationFeeRequest) IntoDBModel() repository.UpdatePlayerRegistrationFeeParams {
	return repository.UpdatePlayerRegistrationFeeParams{
		ID:           rq.ID,
		FeeRemainder: rq.FeeRemainder,
	}
}

// Game request models
type CreateGameWithoutScoreRequest struct {
	HomeTeamID int64            `json:"homeTeamId" binding:"required"`
	AwayTeamID int64            `json:"awayTeamId" binding:"required"`
	GameTime   pgtype.Timestamp `json:"gameTime" binding:"required"`
}

func (rq *CreateGameWithoutScoreRequest) IntoDBModel() repository.CreateGameWithoutScoreParams {
	return repository.CreateGameWithoutScoreParams{
		HomeTeamID: rq.HomeTeamID,
		AwayTeamID: rq.AwayTeamID,
		GameTime:   rq.GameTime,
	}
}

type CreateGameWithScoreRequest struct {
	HomeTeamID int64            `json:"homeTeamId" binding:"required"`
	AwayTeamID int64            `json:"awayTeamId" binding:"required"`
	GameTime   pgtype.Timestamp `json:"gameTime" binding:"required"`
	HomeScore  int32            `json:"homeScore" binding:"required"`
	AwayScore  int32            `json:"awayScore" binding:"required"`
}

func (rq *CreateGameWithScoreRequest) IntoDBModel() repository.CreateGameWithScoreParams {
	return repository.CreateGameWithScoreParams{
		HomeTeamID: rq.HomeTeamID,
		AwayTeamID: rq.AwayTeamID,
		GameTime:   rq.GameTime,
		HomeScore:  rq.HomeScore,
		AwayScore:  rq.AwayScore,
	}
}

type UpdateGameRequest struct {
	ID         int64            `json:"id" binding:"required"`
	HomeTeamID int64            `json:"homeTeamId" binding:"required"`
	AwayTeamID int64            `json:"awayTeamId" binding:"required"`
	GameTime   pgtype.Timestamp `json:"gameTime" binding:"required"`
	HomeScore  int32            `json:"homeScore" binding:"required"`
	AwayScore  int32            `json:"awayScore" binding:"required"`
	Status     string           `json:"status" binding:"required"`
}

func (rq *UpdateGameRequest) IntoDBModel() repository.UpdateGameParams {
	return repository.UpdateGameParams{
		HomeTeamID: rq.HomeTeamID,
		AwayTeamID: rq.AwayTeamID,
		GameTime:   rq.GameTime,
		HomeScore:  rq.HomeScore,
		AwayScore:  rq.AwayScore,
		Status:     rq.Status,
		ID:         rq.ID,
	}
}

type UpdateGameTimeRequest struct {
	ID       int64            `json:"id" binding:"required"`
	GameTime pgtype.Timestamp `json:"gameTime" binding:"required"`
}

func (rq *UpdateGameTimeRequest) IntoDBModel() repository.UpdateGameTimeParams {
	return repository.UpdateGameTimeParams{
		ID:       rq.ID,
		GameTime: rq.GameTime,
	}
}

type UpdateGameScoreAndStatusRequest struct {
	ID        int64  `json:"id" binding:"required"`
	HomeScore int32  `json:"homeScore" binding:"required"`
	AwayScore int32  `json:"awayScore" binding:"required"`
	Status    string `json:"status" binding:"required"`
}

func (rq *UpdateGameScoreAndStatusRequest) IntoDBModel() repository.UpdateGameScoreAndStatusParams {
	return repository.UpdateGameScoreAndStatusParams{
		ID:        rq.ID,
		HomeScore: rq.HomeScore,
		AwayScore: rq.AwayScore,
		Status:    rq.Status,
	}
}

// Payment request models

type CreatePaymentRequest struct {
	PlayerID int64  `json:"playerId" binding:"required"`
	StripeID string `json:"stripeId" binding:"required"`
	Amount   int32  `json:"amount" binding:"required"`
	Status   string `json:"status" binding:"required"`
}

func (rq *CreatePaymentRequest) IntoDBModel() repository.CreatePaymentParams {
	return repository.CreatePaymentParams{
		PlayerID:      rq.PlayerID,
		TransactionID: rq.StripeID,
		Amount:        rq.Amount,
		Status:        rq.Status,
	}
}

type UpdatePaymentStatusRequest struct {
	ID     int64  `json:"id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

func (rq *UpdatePaymentStatusRequest) IntoDBModel() repository.UpdatePaymentStatusParams {
	return repository.UpdatePaymentStatusParams{
		ID:     rq.ID,
		Status: rq.Status,
	}
}

type CreateTeamRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTeamRequest struct {
	Name string `json:"name" binding:"required"`
	ID   int64
}

func (rq *UpdateTeamRequest) IntoDBModel() repository.UpdateTeamNameParams {
	return repository.UpdateTeamNameParams{
		ID:   rq.ID,
		Name: rq.Name,
	}
}

type TeamWithPlayers struct {
	ID            int64                 `json:"id"`
	Name          string                `json:"name"`
	Wins          int32                 `json:"wins"`
	Losses        int32                 `json:"losses"`
	Draws         int32                 `json:"draws"`
	PointsFor     int32                 `json:"pointsFor"`
	PointsAgainst int32                 `json:"pointsAgainst"`
	CreatedAt     pgtype.Timestamp      `json:"createdAt"`
	UpdatedAt     pgtype.Timestamp      `json:"updatedAt"`
	Players       []PlayerSimpleDetails `json:"players"`
}

type PlayerSimpleDetails struct {
	JerseyNumber pgtype.Int4 `json:"jerseyNumber"`
	FirstName    string      `json:"firstName"`
	LastName     string      `json:"lastName"`
}

type GameWithDetails struct {
	repository.ListGamesWithTeamsRow
	HomePlayerStats []PlayerGameStats `json:"homePlayerStats"`
	AwayPlayerStats []PlayerGameStats `json:"awayPlayerStats"`
}

func CreateGameWithDetails[T repository.ListGamesWithTeamsRow | repository.ListTeamScheduleRow](
	game T,
	homeStats []PlayerGameStats,
	awayStats []PlayerGameStats,
) GameWithDetails {
	var g any = game

	var gameRow repository.ListGamesWithTeamsRow
	switch v := g.(type) {
	case repository.ListGamesWithTeamsRow:
		gameRow = v
	case repository.ListTeamScheduleRow:
		gameRow = repository.ListGamesWithTeamsRow(v)
	}

	return GameWithDetails{
		ListGamesWithTeamsRow: gameRow,
		HomePlayerStats:       homeStats,
		AwayPlayerStats:       awayStats,
	}
}

type PlayerGameStats struct {
	PlayerID        int64       `json:"playerId"`
	PlayerFirstName string      `json:"playerFirstName"`
	PlayerLastName  string      `json:"playerLastName"`
	Number          pgtype.Int4 `json:"number"`
	Score           int32       `json:"score"`
}
