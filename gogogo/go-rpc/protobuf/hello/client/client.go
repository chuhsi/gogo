package main

import (
	"context"
	"fmt"
	"log"
	pb "gogogo/go-rpc/protobuf/proto/hello" // 引入proto包
	"google.golang.org/grpc"
)

// 服务地址
const (
	Address = "127.0.0.1:9090"
)

func main() {
	cc, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	c := pb.NewHelloClient(cc)

	req := &pb.HelloRequest{Name: "Go-rpc"}

	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res.Msg)
}
