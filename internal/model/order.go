package model

import "time"

type Order struct {
    ID              uint      `json:"id" gorm:"primaryKey"`
    UserID          uint      `json:"user_id" gorm:"not null"`
    ProductID       uint      `json:"product_id" gorm:"not null"`
    ActivityID      uint      `json:"activity_id" gorm:"not null"`
    OrderNo         string    `json:"order_no" gorm:"unique;not null"`
    Amount          float64   `json:"amount" gorm:"not null"`
    Status          int       `json:"status" gorm:"not null"` // 1:未支付 2:已支付 3:已取消
    PaymentTime     time.Time `json:"payment_time"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
} 