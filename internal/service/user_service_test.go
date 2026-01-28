package service

import (
	"context"
	"testing"
	"time"

	"github.com/gbart/fcabl-api/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

// internal/service/user_service_test.go

type MockUserRepo struct {
	// We use fields to store what we want the mock to return
	mockUsers []repository.ListUsersRow
	mockErr   error
}

func (m *MockUserRepo) ListUsers(ctx context.Context) ([]repository.ListUsersRow, error) {
	return m.mockUsers, m.mockErr
}

func TestListUsers_Success(t *testing.T) {
	// 1. Arrange
	mockData := []repository.ListUsersRow{
		{
			ID:          1,
			Email:       "test_user@gmail.com",
			PhoneNumber: "+1234567890",
			FirstName:   "John",
			LastName:    "Doe",
			Role:        "admin",
			CreatedAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
			UpdatedAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
		},
	}
	mockRepo := &MockUserRepo{mockUsers: mockData, mockErr: nil}

	// Inject mock into the real service
	svc := NewUserService(mockRepo)

	// 2. Act
	ctx := context.Background()
	results, err := svc.ListUsers(ctx, []string{})

	// 3. Assert
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(results) != 1 {
		t.Errorf("expected 1 user, got %d", len(results))
	}
	if results[0].FirstName != "John" {
		t.Errorf("expected username John, got %s", results[0].FirstName)
	}
}
