package main

import (
	"log"

	"github.com/streadway/amqp"
) //导入mq包

// /资料原文:https://mp.weixin.qq.com/s/HprV7_JjXLjBc4tbH6ri0w

// MQURL 格式 amqp://账号：密码@rabbitmq服务器地址：端口号/vhost (默认是5672端口)
// 端口可在 /etc/rabbitmq/rabbitmq-env.conf 配置文件设置，也可以启动后通过netstat -tlnp查看
const MQURL = "amqp://admin:huan91uncc@172.21.138.131:5672/"

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// routing Key
	RoutingKey string
	//MQ链接字符串
	Mqurl string
}

// 创建结构体实例
func NewRabbitMQ(queueName, exchange, routingKey string) *RabbitMQ {
	rabbitMQ := RabbitMQ{
		QueueName:  queueName,
		Exchange:   exchange,
		RoutingKey: routingKey,
		Mqurl:      MQURL,
	}
	var err error
	//创建rabbitmq连接
	rabbitMQ.Conn, err = amqp.Dial(rabbitMQ.Mqurl)
	checkErr(err, "创建连接失败")

	//创建Channel
	rabbitMQ.Channel, err = rabbitMQ.Conn.Channel()
	checkErr(err, "创建channel失败")

	return &rabbitMQ

}

// 释放资源,建议NewRabbitMQ获取实例后 配合defer使用
func (mq *RabbitMQ) ReleaseRes() {
	mq.Conn.Close()
	mq.Channel.Close()
}

func checkErr(err error, meg string) {
	if err != nil {
		log.Fatalf("%s:%s\n", meg, err)
	}
}
