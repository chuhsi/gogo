package main

import (
	"bufio"
	"fmt"
	"gogogo/socket/tcp/proto"
	"io"

	// "io"
	"log"
	"net"

)

func Process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		s, err := proto.Decode(reader)
		if err != io.EOF {
			return
		}
		if err != nil {
			log.Fatalln("decode msg failed, err ", err)
			return
		}
		fmt.Println("收到客户端的数据: ",s)
	}
	// for {
	// 	r := bufio.NewReader(conn)
	// 	var buf [1024]byte
	// 	n, err := r.Read(buf[:])
	// 	if err != nil {
	// 		log.Fatal("r.Read ",err)
	// 		break
	// 	}
	// 	recvStr := string(buf[:n])
	// 	fmt.Println("接受客户端的数据: ",recvStr)
	// 	conn.Write([]byte(recvStr))
	
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go Process(c)
	}
}