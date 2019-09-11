## RabbitMQ阶段总结

一、RabbitMQ主要使用场景：
    <ol>
        <li>程序解耦利器；</li>
        <li>流量消峰；</li>
        <li>异步处理；</li>
    </ol>

二、RabbitMQ核心概念：
    <ul>
        <li>VirtualHost</li>
        <li>Connection</li>
        <li>Exchange</li>
        <li>Channel</li>
        <li>Queue</li>
        <li>Binding</li>
    </ul>
    
三、RabbitMQ 工作中最常用的工作模式：
    <ol>
        <li>
        <p>Simple模式，最简单常用的模式：<br>
        <img src="http://product.luckyhomemart.com/public/img/simple.png" alt="image"></p>
        </li>
        <li>
        <p>Work，工作模式,一个消息只能被一个消费者获取。<br>
        <img src="http://product.luckyhomemart.com/public/img/work.png" alt="image"></p>
        </li>
        <li>
        <p>Publish/Subscribe，订阅模式,消息被路由投递给多个队列，一个消息被多个消费者获取。<br>
        <img src="http://product.luckyhomemart.com/public/img/pub-sub.png" alt="image"></p>
        </li>
        <li>
        <p>Routing，路由模式,一个消息被多个消费者获取。并且消息的目标队列可被生产者指定。<br>
        <img src="http://product.luckyhomemart.com/public/img/routing.png" alt="image"></p>
        </li>
        <li>
        <p>Topic，话题模式,一个消息被多个消费者获取。消息的目标queue可用BindingKey     以通配符， (#：一个或多个词，*：一个词)的方式指定。</p>
        </li>
    </ol>