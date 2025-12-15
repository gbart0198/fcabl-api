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
