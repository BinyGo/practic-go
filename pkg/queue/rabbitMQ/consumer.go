package main

import (
	"fmt"
)

// 消费者订阅
func ExecConsumer() {
	// 初始化mq
	mq := NewRabbitMQ("queue_publisher", "exchange_publisher", "key1")
	defer mq.ReleaseRes() // 完成任务释放资源

	// 1.声明队列（两端都要声明，原因在生产者处已经说明）
	_, err := mq.Channel.QueueDeclare( // 返回的队列对象内部记录了队列的一些信息，这里没什么用
		mq.QueueName, // 队列名
		true,         // 是否持久化
		false,        // 是否自动删除(前提是至少有一个消费者连接到这个队列，之后所有与这个队列连接的消费者都断开时，才会自动删除。注意：生产者客户端创建这个队列，或者没有消费者客户端与这个队列连接时，都不会自动删除这个队列)
		false,        // 是否为排他队列（排他的队列仅对“首次”声明的conn可见[一个conn中的其他channel也能访问该队列]，conn结束后队列删除）
		false,        // 是否阻塞
		nil,          // 额外属性（我还不会用）
	)
	if err != nil {
		fmt.Println("声明队列失败", err)
		return
	}

	// 2.从队列获取消息（消费者只关注队列）consume方式会不断的从队列中获取消息
	msgChanl, err := mq.Channel.Consume(
		mq.QueueName, // 队列名
		"",           // 消费者名，用来区分多个消费者，以实现公平分发或均等分发策略
		true,         // 是否自动应答
		false,        // 是否排他
		false,        // 是否接收只同一个连接中的消息，若为true，则只能接收别的conn中发送的消息
		true,         // 队列消费是否阻塞
		nil,          // 额外属性
	)
	if err != nil {
		fmt.Println("获取消息失败", err)
		return
	}

	for msg := range msgChanl {
		// 这里写你的处理逻辑
		// 获取到的消息是amqp.Delivery对象，从中可以获取消息信息
		fmt.Println(string(msg.Body))
		// msg.Ack(true) // 主动应答

	}

}
