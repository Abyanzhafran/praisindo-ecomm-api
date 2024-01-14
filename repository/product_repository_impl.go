package repository

import (
	"context"
	"dev/models"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{DB: db}
}

func (r *ProductRepositoryImpl) GetAll(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product
	// write ORM
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepositoryImpl) GetById(ctx context.Context, id string) (*models.Product, error) {
	var product *models.Product

	if err := r.DB.Where("id_product = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}
