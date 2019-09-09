package main

import (
	"../RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	rabbitmq1 := RabbitMQ.NewRabbitMQRouting("exImoocRouting", "imooc_routing_key1")
	rabbitmq2 := RabbitMQ.NewRabbitMQRouting("exImoocRouting", "imooc_routing_key2")

	for i := 0; i < 10; i++ {
		rabbitmq1.PublishRouting("routing-1- " + strconv.Itoa(i))
		rabbitmq2.PublishRouting("routing-2- " + strconv.Itoa(i))
		time.Sleep(time.Second)
		fmt.Println("main-for-", i)
	}
}
