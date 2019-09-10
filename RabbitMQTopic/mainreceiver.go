package main

import "../RabbitMQ"

func main() {
	//rabbitmq1 := RabbitMQ.NewRabbitMQTopic("exImoocTopic", "#")
	//rabbitmq1.ReceiverTopic()

	rabbitmq2 := RabbitMQ.NewRabbitMQTopic("exImoocTopic", "imooc.*.key2")
	rabbitmq2.ReceiverTopic()
}
