package main

import "../RabbitMQ"

func main() {
	//发送消息
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("newProduce")
	rabbitmq.PublishPub("订阅模式，发送消息")

}
