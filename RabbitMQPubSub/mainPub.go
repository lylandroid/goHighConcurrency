package main

import (
	"../RabbitMQ"
	"strconv"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		//发送消息
		rabbitmq := RabbitMQ.NewRabbitMQPubSub("newProduce")
		rabbitmq.PublishPub("订阅模式，发送消息 " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
