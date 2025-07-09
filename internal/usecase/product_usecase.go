package usecase

import (
	"context"
	"errors"

	"ragnarok-multifinance/internal/model/database"
	"ragnarok-multifinance/internal/model/request"
	"ragnarok-multifinance/internal/repository"

	"github.com/go-playground/validator/v10"
	"github.com/oklog/ulid/v2"
)

type ProductUsecase struct {
	repo     repository.ProductRepository
	validate *validator.Validate
}

func NewProductUsecase(repo repository.ProductRepository) *ProductUsecase {
	return &ProductUsecase{
		repo:     repo,
		validate: validator.New(),
	}
}

func (u *ProductUsecase) List(ctx context.Context) ([]database.Product, error) {
	return u.repo.List(ctx)
}

func (u *ProductUsecase) Create(ctx context.Context, req request.ProductCreateRequest) (*database.Product, error) {
	if err := u.validate.Struct(req); err != nil {
		return nil, err
	}
	if req.Tenure != 1 && req.Tenure != 2 && req.Tenure != 3 && req.Tenure != 6 {
		return nil, errors.New("invalid tenure")
	}
	id := ulid.Make().String()
	product := &database.Product{
		ID:          id,
		Name:        req.Name,
		Tenure:      req.Tenure,
		MaxLimit:    req.MaxLimit,
		Interest:    req.Interest,
		Description: req.Description,
		IsActive:    true,
	}
	if err := u.repo.Create(ctx, product); err != nil {
		return nil, err
	}
	return product, nil
}
