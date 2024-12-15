package model

import "time"

type Product struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Name        string    `json:"name" gorm:"not null"`
    Description string    `json:"description"`
    Price       float64   `json:"price" gorm:"not null"`
    Stock       int       `json:"stock" gorm:"not null"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type SeckillActivity struct {
    ID         uint      `json:"id" gorm:"primaryKey"`
    ProductID  uint      `json:"product_id" gorm:"not null"`
    StartTime  time.Time `json:"start_time" gorm:"not null"`
    EndTime    time.Time `json:"end_time" gorm:"not null"`
    SeckillPrice float64 `json:"seckill_price" gorm:"not null"`
    Stock     int       `json:"stock" gorm:"not null"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
} 