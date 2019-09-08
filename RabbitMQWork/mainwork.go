package main

import (
	"../RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	//receiveMsg()
	publishMsg()
	/*go func() {
		time.Sleep(2 * time.Second)

	}()
	time.Sleep(5 * time.Second)*/

}

const QueueName = "imoocSimple"

//消费消息
func receiveMsg() {
	RabbitMQ.NewRabbitMQSimple(QueueName).Consume()
}

//生产消息
func publishMsg() {
	for i := 0; i < 100; i++ {
		RabbitMQ.NewRabbitMQSimple(QueueName).PublishSimple("Hello imooc! " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println("publish msg i=", i)
	}
}
