package repository

import (
	"context"
	"dev/models"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]*models.Product, error)
	GetById(ctx context.Context, id string) (*models.Product, error)
}
