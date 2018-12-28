# go-guava
Common tools based on golang

# use ss proxy!!!
# dep ensure -v -update

### gin request info get 
```
//test get request
func TestGinRequestInfo(t *testing.T) {
	r := gin.New()
	//register middle ware 
	//get request info ，save logs or send to mq ?
	r.Use(GinRequestInfo(func(req *RequestInfo) {
		d, _ := json.Marshal(req)
		logger.Debug(string(d))
	}))
	r.GET("/v1/load", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"code": 0})
	})
	r.Run(":7777")
}
```

```$xslt
//mq test for producer
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
```

```$xslt
//mq test for consumer
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
```


![Build Status](https://circleci.com/gh/rmulley/go-fast-sql.svg?style=shield)
[![Test Coverage](https://codeclimate.com/github/rmulley/go-fast-sql/badges/coverage.svg)](https://codeclimate.com/github/rmulley/go-fast-sql/coverage)
[![Code Climate](https://codeclimate.com/github/rmulley/go-fast-sql/badges/gpa.svg)](https://codeclimate.com/github/rmulley/go-fast-sql)
[![GoDoc](https://godoc.org/github.com/rmulley/go-fast-sql?status.svg)](https://godoc.org/github.com/rmulley/go-fast-sql)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/rmulley/go-fast-sql/master/LICENSE)
# go-fast-sql
Package fastsql is a library which extends Go's standard [database/sql](https://golang.org/pkg/database/sql/) library.  It provides performance that's easy to take advantage of.

Even better, the fastsql.DB object embeds the standard sql.DB object meaning access to all the standard database/sql library functionality is preserved.  It also means that integrating fastsql into existing codebases is a breeze.

Additional functionality inclues:
  1. Easy, readable, and performant batch insert queries using the BatchInsert method.
  2. Automatic creation and re-use of prepared statements.
  3. A convenient holder for manually used prepared statements.

## Example usage

```go
package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/rmulley/go-fast-sql"
	"log"
	"net/url"
)

func main() {
	var (
		err error
		i   uint = 1
		dbh *fastsql.DB
	)

	// Create new FastSQL DB object with a flush-interval of 100 rows
	if dbh, err = fastsql.Open("mysql", "user:pass@tcp(localhost:3306)/db_name?"+url.QueryEscape("charset=utf8mb4,utf8&loc=America/New_York"), 100); err != nil {
		log.Fatalln(err)
	}
	defer dbh.Close()

	// Some loop performing SQL INSERTs
	var base="INSERT INTO test_table(id, id2, id3) VALUES(?, ?, ?);"
	for i <= 250 {
		if err = dbh.BatchInsert(base, i, i + 1, i + 2); err != nil {
			log.Fatalln(err)
		}

		i++
	}
	//flush last data into Db
	dbh.LastBatchInsert(base)
}
```



### common date format utils

* Please refer to the source code for more details. 
