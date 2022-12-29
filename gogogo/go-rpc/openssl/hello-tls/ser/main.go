package main

import (
	"context"
	"fmt"
	pb "gogogo/go-rpc/openssl/proto/hello"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	Address = "127.0.0.1:9090"
)

type helloService struct{}
var HelloService = helloService{}

func (helloService) SayHello(c context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error ) {
	resp := new(pb.HelloResponse)
	resp.Msg = fmt.Sprintf("Hello %s", in.Name)
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalln(err)
	}
	// TLS认证
	creds, err := credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
	if err != nil {
		log.Fatalln(err)
	}
	// 实例化grpc server， 并开启TLS认证
	s := grpc.NewServer(grpc.Creds(creds))

	//
	pb.RegisterHelloServer(s, HelloService)

	fmt.Println("Listen on " + Address + " with TSL")

	s.Serve(listen)
}