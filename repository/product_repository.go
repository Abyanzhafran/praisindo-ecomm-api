package repository

import (
	"context"
	"dev/models/domain"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]*domain.Product, error)
	GetById(ctx context.Context, id string) (*domain.Product, error)
	GetByProductName(ctx context.Context, productName string) (*domain.Product, error)
	Add(ctx context.Context, product *domain.Product) error
	Update(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, id string) error
	GetPaginated(ctx context.Context, pageSize int64, offset int) ([]*domain.Product, int, error)
}
