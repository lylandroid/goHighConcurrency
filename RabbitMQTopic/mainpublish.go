package main

import (
	"../RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	rabbitmq1 := RabbitMQ.NewRabbitMQTopic("exImoocTopic", "imooc.topic.key1")
	rabbitmq2 := RabbitMQ.NewRabbitMQTopic("exImoocTopic", "imooc.topic.key2")

	for i := 0; i < 10; i++ {
		rabbitmq1.PublishTopic("topic-1- " + strconv.Itoa(i))
		rabbitmq2.PublishTopic("topic-2- " + strconv.Itoa(i))
		time.Sleep(time.Second)
		fmt.Println("main-topic-for-", i)
	}
}
