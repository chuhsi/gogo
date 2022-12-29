package main

import (
	"fmt"
	pb "gogogo/go-rpc/protobuf/proto/hello" // 引入编译生成的包
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	// "google.golang.org/grpc/grpclog"
)

const (
	Address = "127.0.0.1:9090"
)

type helloService struct{}

var HelloService helloService

// SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
func (helloService) SayHello(c context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Msg = fmt.Sprintf("Hello %s", in.Name)
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalln("net.Listen: ", err)
	}
	server := grpc.NewServer()
	pb.RegisterHelloServer(server, HelloService)
	fmt.Println("listen on ", Address)
	server.Serve(listen)
}
