package nsq_client

import (
	"testing"
	"time"
)

const testTopic = "test"

func TestInitProducer(t *testing.T) {
	strIP1 := "127.0.0.1:4150"
	p, _ := InitProducer(strIP1, testTopic)
	for true {
		time.Sleep(time.Second)
		command := []byte("123")
		err := p.Publish(command)
		if err != nil {
			//retry send message
		}
	}
}
