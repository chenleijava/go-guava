package nsq_client

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"sync"
	"testing"
)

const testChannel = "test-channel"

type ConsumerT struct{}
type ConsumerQ struct{}

//handle message
func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Printf("T address:%s; receive message:%s timeStamp:%d\n", msg.NSQDAddress, msg.Body,
		msg.Timestamp)
	return nil
}

//handle message
func (*ConsumerQ) HandleMessage(msg *nsq.Message) error {
	fmt.Printf("Q address:%s; receive message:%s timeStamp:%d\n", msg.NSQDAddress, msg.Body,
		msg.Timestamp)
	msg.Finish()
	return nil
}
func TestInitConsumer(t *testing.T) {
	var wait sync.WaitGroup
	wait.Add(1)
	InitConsumer(testTopic, testChannel, &ConsumerQ{}, 3, 16, "127.0.0.1:4161")
	wait.Wait()
}
