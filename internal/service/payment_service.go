package service

import (
	"context"

	"github.com/gbart/fcabl-api/internal/repository"
)

type PaymentService interface {
	ListPayments(ctx context.Context) ([]repository.Payment, error)
}

type paymentService struct {
	repository repository.PaymentRepository
}

func NewPaymentService(repository repository.PaymentRepository) PaymentService {
	return &paymentService{repository: repository}
}

func (s *paymentService) ListPayments(ctx context.Context) ([]repository.Payment, error) {
	return s.repository.ListPayments(ctx)
}
