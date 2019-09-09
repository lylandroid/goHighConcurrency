package main

import "../RabbitMQ"

func main() {
	//rabbitmq1 := RabbitMQ.NewRabbitMQRouting("exImoocRouting", "imooc_routing_key1")
	//rabbitmq1.ReceiverRouting()

	rabbitmq2 := RabbitMQ.NewRabbitMQRouting("exImoocRouting", "imooc_routing_key2")
	rabbitmq2.ReceiverRouting()
}
