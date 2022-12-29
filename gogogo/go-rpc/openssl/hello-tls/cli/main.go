package main

import (
	"context"
	"fmt"
	pb "gogogo/go-rpc/openssl/proto/hello"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	Address = "127.0.0.1:9090"
)

func main(){
	creds, err := credentials.NewClientTLSFromFile("../../keys/server.pem", "")
	if err != nil {
		log.Fatalln(err)
	}
	cc, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalln(err)
	}

	defer cc.Close()
	// 初始化客户端
	c := pb.NewHelloClient(cc)

	req := &pb.HelloRequest{Name: "grpc-TLS"}
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.Msg)
}