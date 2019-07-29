package nsq_client

import (
	"github.com/nsqio/go-nsq"
)

type NsqProducer struct {
	Address  string
	producer *nsq.Producer
}

//init new producer
func InitProducer(address string) (*NsqProducer, error) {
	producer, err := newProducer(address)
	if err != nil {
		return nil, err
	}
	return &NsqProducer{
		Address:  address,
		producer: producer,
	}, nil
}

//do new producer
func newProducer(address string) (*nsq.Producer, error) {
	config := nsq.NewConfig()
	return nsq.NewProducer(address, config)
}

// Publish synchronously publishes a message body to the specified topic, returning
// an error if publish failed
func (p *NsqProducer) Publish(topic string, body []byte) error{
	return p.producer.Publish(topic, body)
}
