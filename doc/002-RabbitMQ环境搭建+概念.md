## 一，RabbitMQ介绍
    1，RabbitMQ是面向消息的中间件，用户组件间解耦，主要体现在消息发送者和消息消费者之间无强依赖关系；
    2，RabbitMQ特点：高可用，扩展性，多语言客户端，管理界面等；
    3，主要使用场景：流浪削峰，异步处理，应用解耦
## 二，安装RabbitMQ:
        1，拉取镜像：docker pull rabbitmq:3.7.17-management
        2，根据下载的镜像创建和启动容器：docker run -d --name rabbitmq3.7.17 -p 5672:5672 -p 15672:15672 -v /data:/var/lib/rabbitmq --hostname myRabbit -e RABBITMQ_DEFAULT_VHOST=my_vhost  -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=admin rabbitmq:3.7.17-management
            说明：
                -d 后台运行容器；
                --name 指定容器名；
                -p 指定服务运行的端口（5672：应用访问端口；15672：控制台Web端口号）；
                --hostname  主机名（RabbitMQ的一个重要注意事项是它根据所谓的 “节点名称” 存储数据，默认为主机名）；
                -e 指定环境变量；（RABBITMQ_DEFAULT_VHOST：默认虚拟机名；RABBITMQ_DEFAULT_USER：默认的用户名；RABBITMQ_DEFAULT_PASS：默认用户名的密码）
                容器名为：rabbitmq:3.7.17-management
## 三，RabbitMQ 常用命令
    1,启动容器：docker start 容器名称
    2，停止容器：docker stop 容器名称
    
## 四，RabbitMQ插件管理界面(rabbitmq-plugins)
    1， virtual hosts：可以用于隔离多个环节
    
## 五，RabbitMQ核心概念
    1，Virtual Hosts
    2，Connections
    3，Channels
    4，Exchanges
    5，Queues
    6，Binding(Queues->Exchanges)

## RabbitMQ 6种工作模式
1，Simple模式（最简单的模式）
    ![Image text](images/003-rabbitmq-mode-simple.png)
2，
3，
4，
5，
6，