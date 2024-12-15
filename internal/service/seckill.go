package service

import (
	"context"
	"seckill/internal/model"
	"seckill/internal/repository"
	"github.com/go-redis/redis/v8"
	amqp "github.com/rabbitmq/amqp091-go"
	"encoding/json"
	"time"
)

type SeckillService interface {
	CreateActivity(ctx context.Context, activity *model.SeckillActivity) error
	PlaceOrder(ctx context.Context, userID, activityID uint) error
	GetProducts(ctx context.Context) ([]*model.Product, error)
}

type seckillService struct {
	productRepo repository.ProductRepository
	orderRepo   repository.OrderRepository
	redis       *redis.Client
	mq          *amqp.Channel
}

func NewSeckillService(
	productRepo repository.ProductRepository,
	orderRepo repository.OrderRepository,
	redis *redis.Client,
	mq *amqp.Channel,
) SeckillService {
	return &seckillService{
		productRepo: productRepo,
		orderRepo:   orderRepo,
		redis:       redis,
		mq:          mq,
	}
}

func (s *seckillService) PlaceOrder(ctx context.Context, userID, activityID uint) error {
	// 1. 检查活动是否有效
	product, err := s.productRepo.GetByID(ctx, activityID)
	if err != nil {
		return err
	}
	if product == nil {
		return ErrActivityNotFound
	}

	// 2. 使用Redis检查库存
	stockKey := fmt.Sprintf("seckill:stock:%d", activityID)
	stock, err := s.redis.Get(ctx, stockKey).Int()
	if err != nil {
		return err
	}
	if stock <= 0 {
		return ErrStockNotEnough
	}

	// 3. 预扣减库存(使用Redis的原子操作)
	decrResult := s.redis.Decr(ctx, stockKey)
	if decrResult.Err() != nil {
		return decrResult.Err()
	}
	if decrResult.Val() < 0 {
		// 库存不足,恢复库存
		s.redis.Incr(ctx, stockKey)
		return ErrStockNotEnough
	}

	// 4. 创建订单
	order := &model.Order{
		UserID:     userID,
		ProductID:  activityID,
		Status:     model.OrderStatusPending,
		CreateTime: time.Now(),
	}
	if err := s.orderRepo.Create(ctx, order); err != nil {
		// 创建订单失败,恢复库存
		s.redis.Incr(ctx, stockKey)
		return err
	}

	// 5. 发送消息到MQ进行异步处理
	msg := OrderMessage{
		OrderID:    order.ID,
		UserID:     userID,
		ProductID:  activityID,
		CreateTime: order.CreateTime,
	}
	msgBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = s.mq.Publish(
		"",              // exchange
		"order_queue",   // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msgBody,
		},
	)
	if err != nil {
		// 发送MQ失败,需要回滚
		s.redis.Incr(ctx, stockKey)
		s.orderRepo.Delete(ctx, order.ID)
		return err
	}

	return nil
}

func (s *seckillService) GetProducts(ctx context.Context) ([]*model.Product, error) {
	// 先尝试从Redis获取
	products, err := s.getProductsFromCache(ctx)
	if err == nil {
		return products, nil
	}

	// Redis获取失败，从数据库获取
	products, err = s.productRepo.ListProducts(ctx, 0, 100)
	if err != nil {
		return nil, err
	}

	// 存入Redis缓存
	s.cacheProducts(ctx, products)
	
	return products, nil
}

func (s *seckillService) getProductsFromCache(ctx context.Context) ([]*model.Product, error) {
	key := "seckill:products"
	data, err := s.redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	var products []*model.Product
	if err := json.Unmarshal(data, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (s *seckillService) cacheProducts(ctx context.Context, products []*model.Product) error {
	key := "seckill:products"
	data, err := json.Marshal(products)
	if err != nil {
		return err
	}

	return s.redis.Set(ctx, key, data, time.Hour).Err()
}