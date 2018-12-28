package c

import (
	"dsp-tencentcloud-vpc-ip-security/com/pb"
	"github.com/golang/protobuf/proto"
	"github.com/robfig/cron"
	"log"
	"sync"
	"testing"
)

//测试rabbitmq
func TestInitRabbitmqConn(t *testing.T) {
	var forever sync.WaitGroup
	forever.Add(1)
	Consumers = append(Consumers, &Consumer{QueueName: "dau_queue", Handle: func(data *[]byte) {
		dau := pb.DauReqeust{}
		dau.XXX_Unmarshal(*data)
		tt := proto.MarshalTextString(&dau)
		log.Printf("%s  tt:%s", dau.String(), tt)
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

