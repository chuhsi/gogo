package main

import (
	"fmt"
	"gogogo/zinx/iface"
	"gogogo/zinx/inet"
)

type PingRouter struct {
	inet.BaseRouter
}
func (*PingRouter) Handle(request iface.IRequest) {
	fmt.Println("[Zinx] Call Router Handle")
	// request.GetCurrentConnection().GetConnetion().Write([]byte("handle ping ...\n"))
	request.GetCurrentConnection().SendMsg(200, []byte("ping...ping...ping"))
}

func main() {
	server := inet.New_Server()
	server.AddRouter(1, &PingRouter{})
	server.Serve()
}