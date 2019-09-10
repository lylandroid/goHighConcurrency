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
	ExchangeName string
	//key
	Key string
	//连接信息
	MqUrl string
}

func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName:    queueName,
		ExchangeName: exchange,
		Key:          key,
		MqUrl:        MQURL,
	}
	rabbitmq.dial()
	return rabbitmq
}

func (r *RabbitMQ) dial() {
	var err error
	r.conn, err = amqp.Dial(r.MqUrl)
	r.failOnErr(err, "Failed to connect to RabbitMQ")

	r.channel, err = r.conn.Channel()
	r.failOnErr(err, "Failed to open a channel")
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

func (r *RabbitMQ) applyQueueArgs(exclusive bool) {
	//1,申请队列，如果队列不存在则创建，如果存在则调过创建
	//保证队列存在，消息能发送到队列中
	q, err := r.channel.QueueDeclare(
		r.QueueName, // name
		false,       // 消息是否持久化
		false,       // 最后一个监听失效是否自动删除消息
		exclusive,   // 是否具有排他性（其他用户是否可见）
		false,       // 是否阻塞
		nil,         // 额外属性
	)
	r.QueueName = q.Name
	r.failOnErr(err, "Failed to declare a queue")
}

//试探性申请队列
func (r *RabbitMQ) ApplyQueue() *RabbitMQ{
	r.applyQueueArgs(false)
	return r
}

//发送消息
func (r *RabbitMQ) publish(message string) {
	err := r.channel.Publish(
		r.ExchangeName, // 交换机
		r.Key,
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

//1.2生产简单模式消息
func (r *RabbitMQ) PublishSimple(message string) {
	//1,申请队列
	r.ApplyQueue()
	//2,发送消息到队列中
	r.publish(message)

}

//消费消息
func (r *RabbitMQ) Consume() {
	//1,申请队列
	//r.ApplyQueue()
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

//创建订阅模式RabbitMQ实例
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	return rabbitmq
}

//试探性创建交换机(广播类型:"fanout",路由模式："direct")
func (r *RabbitMQ) applyExchange(kind string) {
	err := r.channel.ExchangeDeclare(
		r.ExchangeName, // name
		kind,           // type 交换机类型（fanout：广播类型）
		true,           // durable 是否持久化
		false,          // auto-deleted 是否自动删除
		//true表示这个exchange不可以被client用来推送消息，
		// 仅用来进行exchange和exchange之间的绑定
		false, // internal
		false, // no-wait
		nil,   // arguments
	)
	r.failOnErr(err, "Failed to declare an exchange")
}

//队列绑定到交换机
func (r *RabbitMQ) bindingQueueExchange() {
	err := r.channel.QueueBind(
		r.QueueName,    // queue name
		r.Key,          // routing key,订阅模式下key必须为空
		r.ExchangeName, // exchange
		false,
		nil,
	)
	r.failOnErr(err, "Failed to bind a queue")
}

//订阅模式生产
func (r *RabbitMQ) PublishPub(message string) {
	//1,试探性创建交换机
	r.applyExchange("fanout")
	r.publish(message)
}

func (r *RabbitMQ) ReceiverSub() {
	//1,试探性创建交换机
	r.applyExchange("fanout")
	//queueName=""表示队列随机生成
	r.applyQueueArgs(true)
	r.bindingQueueExchange()
	//消费消息
	r.Consume()
}

func NewRabbitMQRouting(exchangeName string, routingKey string) *RabbitMQ {
	rabbitMQ := NewRabbitMQ("", exchangeName, routingKey)
	return rabbitMQ
}

func (r *RabbitMQ) PublishRouting(message string) {
	r.applyExchange("direct")
	r.publish(message)
}

func (r *RabbitMQ) ReceiverRouting() {
	r.applyExchange("direct")
	r.applyQueueArgs(true)
	r.bindingQueueExchange()
	r.Consume()
}

func NewRabbitMQTopic(exchangeName string, routingKey string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
	return rabbitmq
}

//话题模式发送消息
func (r *RabbitMQ) PublishTopic(message string) {
	r.applyExchange("topic")
	r.publish(message)
}

//话题模式接收消息
//要注意key规则
//其中“*”表示匹配imooc.hello,
// 但是imooc.hello.one需要用imooc.#才能匹配
func (r *RabbitMQ) ReceiverTopic() {
	r.applyExchange("topic")
	r.applyQueueArgs(true)
	r.bindingQueueExchange()
	r.Consume()
}
