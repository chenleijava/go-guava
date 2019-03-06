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
	Consumers = append(Consumers,
		&Consumer{
			QueueName:           "delivery_queue",
			ConcurrentConsumers: 1,
			PrefetchCount:       0,
			AutoAck:             false,
			Handle: func(data *amqp.Delivery) error {
				return nil
			}})
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
