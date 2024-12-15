package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/seckill/internal/handler"
	"github.com/yourusername/seckill/internal/repository"
	"github.com/yourusername/seckill/internal/service"
	"github.com/yourusername/seckill/internal/pkg/middleware"
	"github.com/yourusername/seckill/internal/pkg/db"
	"github.com/yourusername/seckill/internal/pkg/redis"
	"github.com/yourusername/seckill/internal/pkg/mq"
	"github.com/yourusername/seckill/config"
	"gorm.io/gorm"
	amqp "github.com/rabbitmq/amqp091-go"
)

func initDB() *gorm.DB {
	db, err := db.InitDB(config.GetConfig().Database)
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	redis, err := redis.InitRedis(config.GetConfig().Redis)
	if err != nil {
		panic(err)
	}
	return redis
}

func initMQ() *amqp.Channel {
	mq, err := mq.InitRabbitMQ(config.GetConfig().RabbitMQ)
	if err != nil {
		panic(err)
	}
	return mq
}

func main() {
	// 初始化数据库连接
	db := initDB()
	
	// 初始化Redis连接
	redis := initRedis()
	
	// 初始化RabbitMQ连接
	mq := initMQ()
	
	// 初始化repositories
	userRepo := repository.NewUserRepository(db)
	productRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	
	// 初始化services
	seckillService := service.NewSeckillService(productRepo, orderRepo, redis, mq)
	
	// 初始化handlers
	seckillHandler := handler.NewSeckillHandler(seckillService)
	
	// 设置路由
	r := gin.Default()
	
	// 添加中间件
	r.Use(middleware.RateLimit(100, 200))
	
	// API路由
	api := r.Group("/api")
	{
		api.GET("/seckill/products", seckillHandler.GetProducts)
		api.POST("/seckill/order", seckillHandler.PlaceOrder)
		// 其他路由...
	}
	
	r.Run(":8080")
} 