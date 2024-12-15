package db

import (
	"fmt"
	"seckill/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(conf *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
	)
	
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
} 