package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//url格式：amqp://账号:密码@RabbitMQ服务器地址:端口号/vHost
const MQURL = "amqp://root:root@127.0.0.1:5672/imooc"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//连接信息
	MqUrl string
}

func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		MqUrl:     MQURL,
	}
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)
	rabbitmq.failOnErr(err, "Failed to connect to RabbitMQ")

	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "Failed to open a channel")
	return rabbitmq
}

//端口channel和connection
func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s", message, err)
		panic(fmt.Sprintf("%s: %s", message, err))
	}
}

//1，创建简单模式RabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ(queueName, "", "")
	return rabbitmq
}

func (r *RabbitMQ) ApplyQueue() {
	//1,申请队列，如果队列不存在则创建，如果存在则调过创建
	//保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName, // name
		false,       // 消息是否持久化
		false,       // 最后一个监听失效是否自动删除
		false,       // 是否具有排他性（其他用户是否可见）
		false,       // 是否阻塞
		nil,         // 额外属性
	)
	r.failOnErr(err, "Failed to declare a queue")
}

//1.2生产简单模式消息
func (r *RabbitMQ) PublishSimple(message string) {
	//1,申请队列
	r.ApplyQueue()
	//2,发送消息到队列中
	err := r.channel.Publish(
		r.Exchange,  // 交换机
		r.QueueName, // 队列名称
		// 如果为true,根据Exchange类型和routKey规则，
		// 如果无法找到符合条件的队列那么会把发送的消息返回给发送者
		false,
		//如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者，
		// 则会把消息返还给发送者
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	r.failOnErr(err, "Failed to publish a message")
	log.Println("消息发送成功！")
}

//消费简单消息
func (r *RabbitMQ) ConsumeSimple() {
	//1,申请队列
	r.ApplyQueue()
	//2,接受消息
	msgs, err := r.channel.Consume(
		r.QueueName, // queue
		"",          // 用来区分多个消费者
		true,        // 是否自动应答，消费完成通知RabbitMQ删除该条消息
		false,       // 是否具有排他性（其他用户是否可见）
		false,       // 如果设置为true,表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,       // 队列消息是否阻塞
		nil,         // 其他参数
	)
	r.failOnErr(err, "Failed to register a consumer")

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d := range msgs {
			log.Printf("接收到消息 Received a message: %s", d.Body)
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
