package usecase

import (
	"context"
	"errors"
	"math"
	"time"

	"ragnarok-multifinance/internal/model/database"
	"ragnarok-multifinance/internal/model/request"
	"ragnarok-multifinance/internal/repository"

	"github.com/oklog/ulid/v2"
)

type TransactionUsecase struct {
	transactionRepo repository.TransactionRepository
	productRepo     repository.ProductRepository
	paymentRepo     repository.PaymentRepository
}

func NewTransactionUsecase(tr repository.TransactionRepository, pr repository.ProductRepository, payr repository.PaymentRepository) *TransactionUsecase {
	return &TransactionUsecase{
		transactionRepo: tr,
		productRepo:     pr,
		paymentRepo:     payr,
	}
}

func (u *TransactionUsecase) Create(ctx context.Context, req request.TransactionCreateRequest) (*database.Transaction, error) {
	product, err := u.productRepo.GetByID(ctx, req.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}
	if req.OTR > product.MaxLimit {
		return nil, errors.New("OTR exceeds product max limit")
	}
	if req.AdminFee < 0 {
		return nil, errors.New("invalid admin fee")
	}
	if req.OTR <= 0 {
		return nil, errors.New("invalid OTR")
	}
	if req.AssetName == "" {
		return nil, errors.New("asset name required")
	}
	if product.Tenure != 1 && product.Tenure != 2 && product.Tenure != 3 && product.Tenure != 6 {
		return nil, errors.New("invalid product tenure")
	}
	installment := calculateInstallment(req.OTR, req.AdminFee, product.Interest, product.Tenure)
	id := ulid.Make().String()
	tx := &database.Transaction{
		ID:                id,
		UserID:            req.UserID,
		ProductID:         req.ProductID,
		OTR:               req.OTR,
		AdminFee:          req.AdminFee,
		InstallmentAmount: installment,
		AssetName:         req.AssetName,
		Status:            "PENDING",
		CreatedAt:         time.Now(),
	}
	if err := u.transactionRepo.Create(ctx, tx); err != nil {
		return nil, err
	}
	// TODO: Generate payment schedule and save payments
	return tx, nil
}

func calculateInstallment(otr, adminFee, interest float64, tenure int) float64 {
	total := (otr + adminFee) * (1 + interest/100)
	return math.Round(total/float64(tenure)*100) / 100
}
