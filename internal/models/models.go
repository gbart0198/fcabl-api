// Package models provides structs to be used for API request binding and validation.
package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateUserRequest struct {
	Email        string `json:"email" binding:"required,email"`
	PhoneNumber  string `json:"phoneNumber" binding:"required,e164"`
	FirstName    string `json:"firstName" binding:"required,min=1"`
	LastName     string `json:"lastName" binding:"required,min=1"`
	Role         string `json:"role" binding:"required,oneofci=admin user"`
	Password     string `json:"password" binding:"required,min=1"`
	PasswordHash string
}

type PartialUpdateUserRequest struct {
	Email       *string `json:"email,omitempty" binding:"omitempty,email"`
	PhoneNumber *string `json:"phoneNumber,omitempty" binding:"omitempty,e164"`
	FirstName   *string `json:"firstName,omitempty" binding:"omitempty,min=1"`
	LastName    *string `json:"lastName,omitempty" binding:"omitempty,min=1"`
	Role        *string `json:"role,omitempty" binding:"omitempty,oneofci=admin user"`
	Password    *string `json:"password,omitempty" binding:"omitempty,min=1"`
}

type UpdateUserRequest struct {
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phoneNumber" binding:"required,e164"`
	FirstName   string `json:"firstName" binding:"required,min=1"`
	LastName    string `json:"lastName" binding:"required,min=1"`
	Role        string `json:"role" binding:"required,oneofci=admin user"`
}

type CreatePlayerRequest struct {
	UserID       int64       `json:"userId" binding:"required"`
	TeamID       pgtype.Int8 `json:"teamId" binding:"required"`
	FeeRemainder int32       `json:"feeRemainder" binding:"required"`
	JerseyNumber pgtype.Int4 `json:"jerseyNumber" binding:"required"`
}

type UpdatePlayerRequest struct {
	ID           int64       `json:"id" binding:"required"`
	TeamID       pgtype.Int8 `json:"teamId" binding:"required"`
	FeeRemainder int32       `json:"registrationFeeDue" binding:"required"`
	JerseyNumber pgtype.Int4 `json:"jerseyNumber" binding:"required"`
}

type UpdatePlayerTeamRequest struct {
	ID     int64       `json:"id" binding:"required"`
	TeamID pgtype.Int8 `json:"teamId" binding:"required"`
}

type UpdatePlayerRegistrationFeeRequest struct {
	ID           int64 `json:"id" binding:"required"`
	FeeRemainder int32 `json:"feeRemainder" binding:"required"`
}

// Game request models
type CreateGameWithoutScoreRequest struct {
	HomeTeamID int64            `json:"homeTeamId" binding:"required"`
	AwayTeamID int64            `json:"awayTeamId" binding:"required"`
	GameTime   pgtype.Timestamp `json:"gameTime" binding:"required"`
}

type CreateGameWithScoreRequest struct {
	HomeTeamID int64            `json:"homeTeamId" binding:"required"`
	AwayTeamID int64            `json:"awayTeamId" binding:"required"`
	GameTime   pgtype.Timestamp `json:"gameTime" binding:"required"`
	HomeScore  int32            `json:"homeScore" binding:"required"`
	AwayScore  int32            `json:"awayScore" binding:"required"`
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

type UpdateGameTimeRequest struct {
	ID       int64            `json:"id" binding:"required"`
	GameTime pgtype.Timestamp `json:"gameTime" binding:"required"`
}

type UpdateGameScoreAndStatusRequest struct {
	ID        int64  `json:"id" binding:"required"`
	HomeScore int32  `json:"homeScore" binding:"required"`
	AwayScore int32  `json:"awayScore" binding:"required"`
	Status    string `json:"status" binding:"required"`
}

// Payment request models

type CreatePaymentRequest struct {
	PlayerID int64  `json:"playerId" binding:"required"`
	StripeID string `json:"stripeId" binding:"required"`
	Amount   int32  `json:"amount" binding:"required"`
	Status   string `json:"status" binding:"required"`
}

type UpdatePaymentStatusRequest struct {
	ID     int64  `json:"id" binding:"required"`
	Status string `json:"status" binding:"required"`
}

type CreateTeamRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTeamRequest struct {
	Name string `json:"name" binding:"required"`
	ID   int64
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

type PlayerGameStats struct {
	PlayerID        int64       `json:"playerId"`
	PlayerFirstName string      `json:"playerFirstName"`
	PlayerLastName  string      `json:"playerLastName"`
	Number          pgtype.Int4 `json:"number"`
	Score           int32       `json:"score"`
}

type GameOptionsParams struct {
	Includes string `json:"includes"`
}
