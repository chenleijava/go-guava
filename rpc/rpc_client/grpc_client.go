package main

import (
	"github.com/chenleijava/go-guava/rpc"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

func main() {

	//dial conn
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//
	defer conn.Close()
	//
	greeterClient := pb.NewGreeterClient(conn)
	//http://lanlingzi.cn/post/technical/2016/0802_go_context/
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	//init wait group ,wait  goroutine done !
	wait := &sync.WaitGroup{}
	initialSize := 3
	wait.Add(initialSize)
	var j = 0
	for ; j < initialSize; j++ {
		tmp := j
		go func() {
			time.Sleep(time.Second)
			//ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
			//pass root context
			response, e := greeterClient.SayHello(context.Background(),
				&pb.HelloRequest{Name: *proto.String("客户端发送数据go")})
			if e != nil {
				log.Println(e)
			} else {
				log.Printf("order: %d 响应数据: %s", tmp, response.GetMessage())
			}
			wait.Done() // current send done!
		}()
	}
	wait.Wait()

}
