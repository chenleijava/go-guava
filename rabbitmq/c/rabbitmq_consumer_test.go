package c

import (
	"github.com/robfig/cron"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"testing"
)

//测试rabbitmq
func TestInitRabbitmqConn(t *testing.T) {
	var forever sync.WaitGroup
	forever.Add(1)
	c := &Consumer{
		QueueName:           "delivery_queue",
		ConcurrentConsumers: 1,
		PrefetchCount:       32,
		AutoAck:             false,
		Handle: func(data *amqp.Delivery) error {
			e := data.Ack(true)
			if e != nil {
				log.Printf("ack faild:%v", e)
			}
			return nil
		}}
	Consumers = append(Consumers, c)
	RabbitmqConn("amqp://chenlei:123@localhost:5672/")
	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	forever.Wait()
}

func TestCron(t *testing.T) {
	c := cron.New()
	c.AddFunc("*/1 * * * * ?", func() { log.Println("every 3 seconds run ... ...") })
	c.Run()
	c.Stop() // Stop the scheduler (does not stop any jobs already running).
}
