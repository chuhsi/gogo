package main

import (
	"fmt"
	"gogogo/zinx/inet"
	"net"
	"time"
)

var (
	network = "tcp"
	address = "127.0.0.1:9090"
)

func main() {
	fmt.Println("[Zinx] Client Starting ...")
	time.Sleep(500 * time.Millisecond)
	conn, err := net.Dial(network, address)
	if err != nil {
		fmt.Println("net.Dial err", err)
		return
	}
	fmt.Println("[Zinx] Client Started ...")
	for {
		dp := inet.New_DataPack()
		binary, err := dp.Pack(inet.New_Message(1, []byte("Hello Zinx")))
		if err != nil {
			fmt.Println("dp.Pack err", err)
			return
		}
		_, err = conn.Write(binary)
		if err != nil {
			fmt.Println("conn.Write err", err)
			return
		}
		headData := make([]byte, dp.GetHeadLen())
		// buf := make([]byte, 512)
		_, err = conn.Read(headData)
		if err != nil {
			fmt.Println("conn.Read err", err)
			return
		}
		data, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("dp.Unpack err", err)
			return
		}
		if data.GetMsgDataLen() > 0 {
			msg := data.(*inet.Message)
			msg.Data = make([]byte, msg.GetMsgDataLen())
			_, err := conn.Read(msg.Data)
			if err != nil {
				fmt.Println("conn.Read err", err)
				return
			}
			fmt.Println("[Zinx] ------>Recv Server Msg: ID=", msg.ID, "Len=", msg.DataLen, "data=", string(msg.Data))
		}
		time.Sleep(500 * time.Millisecond)
	}
}
