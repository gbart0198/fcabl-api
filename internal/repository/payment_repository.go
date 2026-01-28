package repository

import (
	"context"
)

type PaymentRepository interface {
	ListPayments(ctx context.Context) ([]Payment, error)
}

type paymentRepository struct {
	queries *Queries
}

func NewPaymentRepository(queries *Queries) PaymentRepository {
	return &paymentRepository{queries: queries}
}

func (r *paymentRepository) ListPayments(ctx context.Context) ([]Payment, error) {
	return r.queries.ListPayments(ctx)
}
