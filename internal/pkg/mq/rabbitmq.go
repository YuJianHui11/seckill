package mq

import (
	"fmt"
	"seckill/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func InitRabbitMQ(conf *config.RabbitMQConfig) (*amqp.Channel, error) {
	conn, err := amqp.Dial(conf.URL)
	if err != nil {
		return nil, err
	}
	
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	
	// 声明队列
	_, err = ch.QueueDeclare(
		"order_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	
	return ch, err
} 