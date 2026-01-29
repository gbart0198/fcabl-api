package service

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/gbart/fcabl-api/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

// internal/service/user_service_test.go

type MockUserRepo struct {
	// We use fields to store what we want the mock to return
	mockUsers []repository.User
	mockErr   error
}

func (m *MockUserRepo) ListUsers(ctx context.Context) ([]repository.User, error) {
	return m.mockUsers, m.mockErr
}

func TestListUsers_Success(t *testing.T) {
	mockData := createMockUsers(1)
	mockRepo := &MockUserRepo{mockUsers: mockData, mockErr: nil}
	svc := NewUserService(mockRepo)
	ctx := context.Background()

	results, err := svc.ListUsers(ctx, []string{})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if len(results) != 1 {
		t.Errorf("expected 1 user, got %d", len(results))
	}
	if results[0].FirstName != "John1" {
		t.Errorf("expected username John, got %s", results[0].FirstName)
	}
}

func createMockUsers(numUsers int) []repository.User {
	users := make([]repository.User, numUsers)
	for i := range numUsers {
		users[i] = repository.User{
			ID:          int64(i + 1),
			Email:       fmt.Sprintf("test_user_%d@gmail.com", i+1),
			PhoneNumber: fmt.Sprintf("+123456789%d", i+1),
			FirstName:   fmt.Sprintf("John%d", i+1),
			LastName:    fmt.Sprintf("Doe%d", i+1),
			Role:        "admin",
			CreatedAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
			UpdatedAt: pgtype.Timestamp{
				Time:  time.Now(),
				Valid: true,
			},
		}
	}
	return users
}
