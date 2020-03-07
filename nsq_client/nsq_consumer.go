package nsq_client

import (
	"github.com/nsqio/go-nsq"
	"time"
)

//init a consumer
//topic: subjects of interest to consumers
//channel: consumer subscription logical grouping
//handler:  data handle
//concurrency: the number of goroutines to spawn for message handling.
//maxInFlight: batch
//address: nsqlookupd addresses (cluster at least 3 nodes)
func InitConsumer(topic string, channel string, handler nsq.Handler,
	concurrency, maxInFlight int, address ...string) {
	cfg := nsq.NewConfig()
	cfg.MaxAttempts = 65535                    //max attempts
	cfg.LookupdPollInterval = time.Second * 60 //lookupd pool interval
	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		panic(err)
	}
	c.SetLogger(nil, 0)
	c.ChangeMaxInFlight(maxInFlight)
	c.AddConcurrentHandlers(handler, concurrency)
	if err := c.ConnectToNSQLookupds(address); err != nil {
		panic(err)
	}
}
