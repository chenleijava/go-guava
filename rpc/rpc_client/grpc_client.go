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
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		{
			_conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			conn = _conn
		}
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
	wait.Add(1)
	var i = 0
	var j = 0
	for ; j <= 10; j++ {
		go func() {
			for i < 1 {
				time.Sleep(100 * time.Millisecond)
				//ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
				//pass root context
				response, e := greeterClient.SayHello(context.Background(),
					&pb.HelloRequest{Name: *proto.String("客户端发送数据go")})
				if e != nil {
					log.Println(e)
				} else {
					log.Println("响应数据\n", response.GetMessage())
				}
			}
		}()
		log.Printf(">>>>>>>>")
	}
	wait.Wait()

}
