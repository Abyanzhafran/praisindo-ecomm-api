package repository

import (
	"context"
	"dev/models/domain"

	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{DB: db}
}

func (r *ProductRepositoryImpl) GetAll(ctx context.Context) ([]*domain.Product, error) {
	var products []*domain.Product
	// write ORM
	if err := r.DB.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepositoryImpl) GetById(ctx context.Context, id string) (*domain.Product, error) {
	var product *domain.Product

	if err := r.DB.Where("id_product = ?", id).First(&product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (r *ProductRepositoryImpl) Add(ctx context.Context, product *domain.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepositoryImpl) Update(ctx context.Context, product *domain.Product) error {
	return r.DB.Model(&domain.Product{}).Where("id_product = ?", product.IDProduct).Updates(product).Error
}

func (r *ProductRepositoryImpl) Delete(ctx context.Context, id string) error {
	return r.DB.Where("id_product = ?", id).Delete(&domain.Product{}).Error
}
