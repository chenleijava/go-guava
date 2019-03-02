package c

import (
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

//https://github.com/rabbitmq/rabbitmq-tutorials/tree/master/go
var Consumers []*Consumer

const (
	prefetchCount = 16
)

//消费者
type Consumer struct {
	QueueName           string                   // consumer listen queue
	Handle              func(data *[]byte)  // hand mq message
	ConcurrentConsumers int                      // 并发消费个数
}

//初始化mq链接
func RabbitmqConn(url string) (*amqp.Connection, *error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		//尝试重新链接，直到成功
		var waitConn sync.WaitGroup
		waitConn.Add(1)
		go func() {
			for true {
				cc, err := amqp.Dial(url)
				if err != nil {
					log.Printf("dail error:%s", err)
					time.Sleep(1 * time.Second) //sleep
					continue
				} else {
					log.Printf("connect success... ...")
					conn = cc
					waitConn.Done()
					break
				}
			}
		}()
		waitConn.Wait()
	}

	go func() {
		log.Printf("mq consumer wait here for the channel to be closed,then retry connect!!!")
		// Waits here for the channel to be closed
		log.Printf("connect closing case:%s ", <-conn.NotifyClose(make(chan *amqp.Error)))
		//begin to reconnect
		_ = conn.Close()
		reconnect(url)
	}()

	//初始化消费端
	for _, consumer := range Consumers {
		err := initConsumer(conn, consumer.QueueName, consumer.ConcurrentConsumers, consumer.Handle)
		if err != nil {
			log.Fatalf("%s", *err)
			return nil, err
		}
	}
	log.Printf("init consumer successful ... ...")
	return conn, nil
}

//断线重连
func reconnect(url string) {
	log.Printf("begin to reconnect :%s", url)
	for true {
		_, err := RabbitmqConn(url)
		if err != nil {
			log.Printf("reconnect error:%s", *err)
			time.Sleep(3 * time.Second)
		} else {
			return
		}
	}
}

//初始化 consumer
func initConsumer(conn *amqp.Connection, queueName string, concurrentConsumers int, handle func(data *[]byte) ) *error {
	for i := 0; i != concurrentConsumers; i++ {
		channel, err := conn.Channel()
		if err != nil {
			return &err
		}
		queue, err := channel.QueueDeclarePassive(
			queueName, // name
			true,      // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)
		if err != nil {
			return &err
		}
		//channel qps
		_ = channel.Qos(prefetchCount, 0, false)
		deliveries, err := channel.Consume(
			queue.Name,            // queue
			queueName+"_consumer", // consumer
			true,                  // auto ack is true
			false,                 // exclusive
			false,                 // no local
			false,                 // no wait
			nil,                   // args
		)
		if err != nil {
			return &err
		}
		//
		go func() {
			for msg := range deliveries {
				// 逻辑处理
				handle(&msg.Body)
			}
		}()
	}
	return nil
}
