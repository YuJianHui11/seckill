package repository

import (
	"gorm.io/gorm"
	"seckill/internal/model"
)

type ProductRepository interface {
	Create(ctx context.Context, product *model.Product) error
	GetByID(ctx context.Context, id uint) (*model.Product, error)
	UpdateStock(ctx context.Context, id uint, stock int) error
	ListProducts(ctx context.Context, offset, limit int) ([]*model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// 实现ProductRepository接口的方法... 

func (r *productRepository) ListProducts(ctx context.Context, offset, limit int) ([]*model.Product, error) {
	var products []*model.Product
	result := r.db.WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Find(&products)
	
	if result.Error != nil {
		return nil, result.Error
	}
	
	return products, nil
}