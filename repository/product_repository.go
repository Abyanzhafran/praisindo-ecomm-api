package repository

import (
	"context"
	"dev/models/domain"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]*domain.Product, error)
	GetById(ctx context.Context, id string) (*domain.Product, error)
	Add(ctx context.Context, product *domain.Product) error
	Update(ctx context.Context, product *domain.Product) error
}
