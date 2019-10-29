package nsq_client

import (
	"github.com/nsqio/go-nsq"
)

type NsqProducer struct {
	Address  string
	producer *nsq.Producer
	topic    string
}

//init new producer
func InitProducer(address, topic string) (*NsqProducer, error) {
	producer, err := newProducer(address)
	if err != nil {
		return nil, err
	}
	return &NsqProducer{
		Address:  address,
		producer: producer,
		topic:    topic,
	}, nil
}

//do new producer
func newProducer(address string) (*nsq.Producer, error) {
	config := nsq.NewConfig()
	return nsq.NewProducer(address, config)
}

// Publish synchronously publishes a message body to the specified topic, returning
// an error if publish failed
func (p *NsqProducer) Publish(body []byte) error {
	return p.producer.Publish(p.topic, body)
}
