package main

import (
	"./RabbitMQ"
)

func main() {
	//publishMsg()
	receiveMsg()
}

const QueueName = "imoocSimple"

//消费消息
func receiveMsg() {
	RabbitMQ.NewRabbitMQSimple(QueueName).ApplyQueue().Consume()
}

//生产消息
func publishMsg() {
	RabbitMQ.NewRabbitMQSimple(QueueName).PublishSimple("Hello imooc!")
}
