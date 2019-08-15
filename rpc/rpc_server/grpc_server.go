package main

import (
	"github.com/chenleijava/go-guava/rpc"
	"github.com/chenleijava/go-guava/zlog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

const (
	port = ":8888"
)

//https://github.com/grpc-ecosystem/go-grpc-middleware

// server is used to implement helloworld.GreeterServer.
type GreeterServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *GreeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "form server 收到客户端消息: " + in.Name}, nil
}

//register service
func RegisterServer(server *grpc.Server) {
	pb.RegisterGreeterServer(server, &GreeterServer{})
	//todo  register service
}

func main() {
	_log := zlog.NewLog2Console()
	//create new grpc server
	server := grpc.NewServer()

	//register service
	RegisterServer(server)

	//last start grpc server
	listener, er := net.Listen("tcp", port)
	if er != nil {
		_log.Sugar().Fatal(er)
	}
	_log.Sugar().Info("grpc server listener:", listener.Addr())

	//start service
	_ = server.Serve(listener)
}
