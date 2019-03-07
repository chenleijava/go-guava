package p

import (
	"log"
	"sync"
	"testing"
	"time"
)

//tracking.delivery.routing.key=delivery_routing_key
//tracking.delivery.exchange=delivery_exchange
//tracking.delivery.queue=delivery_queue
//测试rabbitmq
func TestInitRabbitmqConn(t *testing.T) {
	var forever sync.WaitGroup
	forever.Add(1)
	var i = 0
	p := &Producer{ExchangeName: "delivery_exchange", QueueName: "delivery_queue", RouteKey: "delivery_routing_key"}
	for i != 1 {
		Producers = append(Producers, p)
		i++
	}

	RabbitmqConn("amqp://chenlei:123@localhost:5672/")

	for true {
		d := []byte("123")
		e := p.Send(&d)
		//log.Printf("send>>>")
		if e != nil {
			log.Printf("send err:%s", e.Error())
		}
		time.Sleep(10*time.Microsecond)
	}
	forever.Wait()
}
