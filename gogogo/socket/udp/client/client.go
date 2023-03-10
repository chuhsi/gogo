package main

import (
	"fmt"
	"log"
	"net"
)

func main()  {
	socket, err := net.DialUDP("udp",nil, &net.UDPAddr{
		IP: net.IPv4(0,0,0,0),
		Port: 30000,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer socket.Close()

	sendD := []byte("hello updfjadsfjldsjflsdakjfl")
	_ , err = socket.Write(sendD)
	if err != nil {
		log.Fatal(err)
		return
	}
	data := make([]byte, 4096)
	n, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("recv: %v addr: %v count: %v\n",string(data[:n]), remoteAddr, n)
}