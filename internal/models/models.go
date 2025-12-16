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

// Team request models
type CreateTeamRequest struct {
	Name          string `json:"name" binding:"required"`
	Wins          int32  `json:"wins"`
	Losses        int32  `json:"losses"`
	Draws         int32  `json:"draws"`
	PointsFor     int32  `json:"pointsFor"`
	PointsAgainst int32  `json:"pointsAgainst"`
}

func (rq *CreateTeamRequest) IntoDBModel() repository.CreateTeamParams {
	return repository.CreateTeamParams{
		Name:          rq.Name,
		Wins:          rq.Wins,
		Losses:        rq.Losses,
		Draws:         rq.Draws,
		PointsFor:     rq.PointsFor,
		PointsAgainst: rq.PointsAgainst,
	}
}

type UpdateTeamRequest struct {
	ID            int64  `json:"id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Wins          int32  `json:"wins" binding:"required"`
	Losses        int32  `json:"losses" binding:"required"`
	Draws         int32  `json:"draws" binding:"required"`
	PointsFor     int32  `json:"pointsFor" binding:"required"`
	PointsAgainst int32  `json:"pointsAgainst" binding:"required"`
}

func (rq *UpdateTeamRequest) IntoDBModel() repository.UpdateTeamParams {
	return repository.UpdateTeamParams{
		ID:            rq.ID,
		Name:          rq.Name,
		Wins:          rq.Wins,
		Losses:        rq.Losses,
		Draws:         rq.Draws,
		PointsFor:     rq.PointsFor,
		PointsAgainst: rq.PointsAgainst,
	}
}

// Player request models
// BUG: Using binding:"required" on bool types makes validation fail if a false
// value is passed, because the validation cannot distinguish between a false and no value because
// of the zero value of bool.

type CreatePlayerRequest struct {
	UserID             int64          `json:"userId" binding:"required"`
	TeamID             pgtype.Int8    `json:"teamId" binding:"required"`
	RegistrationFeeDue pgtype.Numeric `json:"registrationFeeDue" binding:"required"`
	IsFullyRegistered  bool           `json:"isFullyRegistered"`
	IsActive           bool           `json:"isActive"`
	JerseyNumber       pgtype.Int4    `json:"jerseyNumber" binding:"required"`
}

func (rq *CreatePlayerRequest) IntoDBModel() repository.CreatePlayerParams {
	return repository.CreatePlayerParams{
		UserID:             rq.UserID,
		TeamID:             rq.TeamID,
		RegistrationFeeDue: rq.RegistrationFeeDue,
		IsFullyRegistered:  rq.IsFullyRegistered,
		IsActive:           rq.IsActive,
		JerseyNumber:       rq.JerseyNumber,
	}
}

type UpdatePlayerRequest struct {
	ID                 int64          `json:"id" binding:"required"`
	TeamID             pgtype.Int8    `json:"teamId" binding:"required"`
	RegistrationFeeDue pgtype.Numeric `json:"registrationFeeDue" binding:"required"`
	IsFullyRegistered  bool           `json:"isFullyRegistered"`
	IsActive           bool           `json:"isActive"`
	JerseyNumber       pgtype.Int4    `json:"jerseyNumber" binding:"required"`
}

func (rq *UpdatePlayerRequest) IntoDBModel() repository.UpdatePlayerParams {
	return repository.UpdatePlayerParams{
		ID:                 rq.ID,
		TeamID:             rq.TeamID,
		RegistrationFeeDue: rq.RegistrationFeeDue,
		IsFullyRegistered:  rq.IsFullyRegistered,
		IsActive:           rq.IsActive,
		JerseyNumber:       rq.JerseyNumber,
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

type UpdatePlayerRegistrationStatusRequest struct {
	ID                 int64          `json:"id" binding:"required"`
	RegistrationFeeDue pgtype.Numeric `json:"registrationFeeDue" binding:"required"`
	IsFullyRegistered  bool           `json:"isFullyRegistered"`
}

func (rq *UpdatePlayerRegistrationStatusRequest) IntoDBModel() repository.UpdatePlayerRegistrationStatusParams {
	return repository.UpdatePlayerRegistrationStatusParams{
		ID:                 rq.ID,
		RegistrationFeeDue: rq.RegistrationFeeDue,
		IsFullyRegistered:  rq.IsFullyRegistered,
	}
}

// Game request models
type CreateGameRequest struct {
	HomeTeamID int64            `json:"homeTeamId" binding:"required"`
	AwayTeamID int64            `json:"awayTeamId" binding:"required"`
	GameTime   pgtype.Timestamp `json:"gameTime" binding:"required"`
}

func (rq *CreateGameRequest) IntoDBModel() repository.CreateGameParams {
	return repository.CreateGameParams{
		HomeTeamID: rq.HomeTeamID,
		AwayTeamID: rq.AwayTeamID,
		GameTime:   rq.GameTime,
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
	PlayerID int64          `json:"playerId" binding:"required"`
	StripeID string         `json:"stripeId" binding:"required"`
	Amount   pgtype.Numeric `json:"amount" binding:"required"`
	Status   string         `json:"status" binding:"required"`
}

func (rq *CreatePaymentRequest) IntoDBModel() repository.CreatePaymentParams {
	return repository.CreatePaymentParams{
		PlayerID: rq.PlayerID,
		StripeID: rq.StripeID,
		Amount:   rq.Amount,
		Status:   rq.Status,
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
