package p

import (
	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

//https://github.com/rabbitmq/rabbitmq-tutorials/blob/master/go/emit_log_direct.go

type Producer struct {
	RouteKey     string        // 路由key
	QueueName    string        // 队列名
	ExchangeName string        //交换机
	ch           *amqp.Channel // send  channel
}

var Producers []*Producer

//初始化mq链接
func RabbitmqConn(url string) (*amqp.Connection, *error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		//尝试重新链接，直到成功
		var waitConn sync.WaitGroup
		waitConn.Add(1)
		go func() {
			for true {
				cc, err := amqp.Dial(url)
				if err != nil {
					log.Printf("dail error:%s", err)
					time.Sleep(1 * time.Second) //sleep
					continue
				} else {
					log.Printf("kepp re_connect success ,exit current goruntue... ...")
					conn = cc
					waitConn.Done()
					break
				}
			}
		}()
		waitConn.Wait()
	} else {
		log.Printf("first time connect success >>>")
	}

	go func() {
		log.Printf("mq producer wait here for the channel to be closed,then retry connect!!!")
		// Waits here for the channel to be closed
		log.Printf("connect closing case:%s ", <-conn.NotifyClose(make(chan *amqp.Error)))
		log.Printf("close producers channel,conn!")
		for _, p := range Producers {
			_ = p.ch.Close()
			p.ch = nil// next retry connect
		}
		//close conn,begin to reconnect
		_ = conn.Close()
		//try to re connect!
		reconnect(url)
	}()

	//init producers
	for _, p := range Producers {
		if p.ch == nil {
			_ = p.initProducer(conn)
		}
	}
	log.Printf("init producer done")
	return conn, nil
}

func (p *Producer) initProducer(conn *amqp.Connection) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	//exchange
	err = ch.ExchangeDeclare(
		p.ExchangeName, // name
		"direct",       // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)

	//queue
	q, err := ch.QueueDeclare(
		p.QueueName, // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	if err != nil {
		return err
	}

	//bind
	log.Printf("Binding queue %s to exchange %s with routing key %s", q.Name, p.ExchangeName, p.RouteKey)
	err = ch.QueueBind(
		q.Name,         // queue name
		p.RouteKey,     // routing key
		p.ExchangeName, // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")
	p.ch = ch
	return nil
}

//Send msg to mq
func (p *Producer) Send(data *[]byte) error {
	if p.ch!=nil{
		return p.ch.Publish(p.ExchangeName, p.RouteKey, false, false, amqp.Publishing{Body: *data})
	}else {
		return errors.New("reconnecting!")
	}
}

//断线重连
func reconnect(url string) {
	log.Printf("begin to reconnect :%s", url)
	for true {
		_, err := RabbitmqConn(url)
		if err != nil {
			log.Printf("reconnect error:%s", *err)
			time.Sleep(3 * time.Second)
		} else {
			log.Printf("re-connect success>>>>>>>>>")
			return
		}
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
