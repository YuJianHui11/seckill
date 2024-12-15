package repository

import (
	"context"
	"gorm.io/gorm"
	"seckill/internal/model"
)

type OrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
	GetByID(ctx context.Context, id uint) (*model.Order, error)
	GetByOrderNo(ctx context.Context, orderNo string) (*model.Order, error)
	UpdateStatus(ctx context.Context, id uint, status int) error
	Delete(ctx context.Context, id uint) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

// 实现OrderRepository接口的方法
func (r *orderRepository) Create(ctx context.Context, order *model.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *orderRepository) GetByID(ctx context.Context, id uint) (*model.Order, error) {
	var order model.Order
	err := r.db.WithContext(ctx).First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) GetByOrderNo(ctx context.Context, orderNo string) (*model.Order, error) {
	var order model.Order
	err := r.db.WithContext(ctx).Where("order_no = ?", orderNo).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) UpdateStatus(ctx context.Context, id uint, status int) error {
	return r.db.WithContext(ctx).Model(&model.Order{}).Where("id = ?", id).
		Update("status", status).Error
}

func (r *orderRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&model.Order{}, id).Error
} 