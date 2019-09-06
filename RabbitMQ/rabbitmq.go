package RabbitMQ

//url格式	amqp://账号:密码@RabbitMQ服务器地址:端口号/vHost
const MQURL = "amqp://root:root@127.0.0.1:5672/imooc"

type RabbitMQ struct {
	conn *amqp.Con
}
