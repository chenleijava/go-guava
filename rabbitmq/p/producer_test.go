package p

import (
	"dsp-tencentcloud-vpc-ip-security/com/pb"
	"github.com/golang/protobuf/proto"
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
	for i != 1 {
		Producers = append(Producers, &Producer{ExchangeName: "delivery_exchange", QueueName: "delivery_queue", RouteKey: "delivery_routing_key"})
		i++
	}

	RabbitmqConn("amqp://chenlei:123@localhost:5672/")

	for true {
		time.Sleep(10 * time.Millisecond)
		data := &pb.DeliveryData{DataType: pb.DataType_PLATFORM_AD, DeviceID: "123qwe",
			ChannelName: "test_002", TimeStamp: uint64(time.Now().Unix())}
		if body, e := proto.Marshal(data); e == nil {
			for _,p:=range Producers{
				sendErr := p.Send(&body)
				if sendErr != nil {
					log.Printf("send err:%s", sendErr)
				}
			}
		} else {
			log.Printf("proto marshl error%s", e.Error())
		}
	}
	forever.Wait()
}
