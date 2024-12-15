package model

import "time"

type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Username  string    `json:"username" gorm:"unique;not null"`
    Password  string    `json:"-" gorm:"not null"`  // 密码不返回给前端
    Email     string    `json:"email" gorm:"unique"`
    Phone     string    `json:"phone"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
} 