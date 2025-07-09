package usecase

import (
	"context"
	"errors"
	"time"

	"ragnarok-multifinance/internal/model/database"
	"ragnarok-multifinance/internal/repository"
)

type PaymentUsecase struct {
	paymentRepo     repository.PaymentRepository
	transactionRepo repository.TransactionRepository
}

func NewPaymentUsecase(pr repository.PaymentRepository, tr repository.TransactionRepository) *PaymentUsecase {
	return &PaymentUsecase{
		paymentRepo:     pr,
		transactionRepo: tr,
	}
}

func (u *PaymentUsecase) RecordPayment(ctx context.Context, payment *database.Payment) error {
	// Validate transaction exists
	tx, err := u.transactionRepo.GetByID(ctx, payment.TransactionID)
	if err != nil || tx == nil {
		return errors.New("transaction not found")
	}
	payment.Status = "PAID"
	paidAt := time.Now()
	payment.PaidAt = &paidAt
	return u.paymentRepo.Create(ctx, payment)
}

func (u *PaymentUsecase) GetSchedule(ctx context.Context, transactionID string) ([]database.Payment, error) {
	return u.paymentRepo.ListByTransaction(ctx, transactionID)
}
