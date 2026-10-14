package global

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	RDB    *redis.Client
	DB     *gorm.DB
	CH     *amqp.Channel
	MQConn *amqp.Connection
)
