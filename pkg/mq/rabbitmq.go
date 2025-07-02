package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"quanfuxia/pkg/config"
)

var MQConn *amqp.Connection
var MQChannel *amqp.Channel

func Init() {
	var err error
	MQConn, err = amqp.Dial(config.Cfg.RabbitMQ.URL)
	if err != nil {
		log.Fatalf("RabbitMQ连接失败: %v", err)
	}

	MQChannel, err = MQConn.Channel()
	if err != nil {
		log.Fatalf("创建Channel失败: %v", err)
	}

	log.Println("✅ RabbitMQ 已连接")
}
