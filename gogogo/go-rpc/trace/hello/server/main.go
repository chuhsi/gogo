package main

import (
	"context"
	"fmt"
	pb "gogogo/go-rpc/trace/proto/hello"
	"log"
	"net"
	"net/http"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
)

const (
	Address = "127.0.0.1:9090"
)

type helloService struct {}

var HelloService = helloService{}

func (helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	hr := new(pb.HelloResponse)
	hr.Msg = fmt.Sprintf("Hello %s.", in.Name)
	return hr, nil
}

func main() {
	l, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, HelloService)
	go startTrace()

	fmt.Println("Listen on " + Address)
	s.Serve(l)
}

func startTrace()  {
	trace.AuthRequest = func(req *http.Request) (any bool, sensitive bool) {
		return true, true
	}
	go http.ListenAndServe(":9999",nil)
	fmt.Println("Trace listen on 9999")
}