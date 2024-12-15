package handler

import (
	"github.com/gin-gonic/gin"
	"seckill/internal/service"
	"net/http"
	"strconv"
	"go.uber.org/zap"
)

type SeckillHandler struct {
	seckillService service.SeckillService
	logger *zap.Logger
}

func NewSeckillHandler(seckillService service.SeckillService, logger *zap.Logger) *SeckillHandler {
	return &SeckillHandler{
		seckillService: seckillService,
		logger: logger,
	}
}

func (h *SeckillHandler) PlaceOrder(c *gin.Context) {
	// 1. 参数验证
	// 2. 调用service
	// 3. 返回结果
}

func (h *SeckillHandler) GetProducts(c *gin.Context) {
	// 获取查询参数
	offset := 0
	limit := 100
	if page, err := strconv.Atoi(c.DefaultQuery("page", "1")); err == nil {
		offset = (page - 1) * limit
	}

	products, err := h.seckillService.GetProducts(c.Request.Context())
	if err != nil {
		h.logger.Error("Failed to get products", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 转换为响应格式
	var response []gin.H
	for _, p := range products {
		response = append(response, gin.H{
			"id": p.ID,
			"name": p.Name,
			"description": p.Description,
			"price": p.Price,
			"seckillPrice": p.Price * 0.5, // 示例：秒杀���为原价的一半
			"stock": p.Stock,
			"image": p.Image,
		})
	}

	c.JSON(http.StatusOK, response)
} 