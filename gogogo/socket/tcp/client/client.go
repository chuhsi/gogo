package main

import (
	// "bufio"
	// "fmt"
	"fmt"
	"gogogo/socket/tcp/proto"
	"log"
	"net"
	// "os"
	// "strings"
)

func main() {
	c, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer c.Close()
	for i := 0; i < 10; i++ {
		msg := "hello, hello. hao are you"
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err", err)
			return
		}
		c.Write(data)
	}
	// r := bufio.NewReader(os.Stdin)
	// for {
	// 	input, _ := r.ReadString('\n')
	// 	s := strings.Trim(input, "\r\n")
	// 	if strings.ToUpper(s) == "Q" {
	// 		return
	// 	}
	// 	_, err := c.Write([]byte(s))
	// 	if err != nil {
	// 		return
	// 	}
	// 	buf := [512]byte{}
	// 	n, err := c.Read(buf[:])
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		return
	// 	}
	// 	fmt.Println("服务端: ",string(buf[:n]))
	// }
}
