package iface

import "net"

type IConnection interface {
	// 启动连接
	Start()
	// 停止连接
	Stop()
	// 获取当前的连接
	GetConnetion() *net.TCPConn
	// 获取当前连接模块的连接ID
	GetConnID() uint32
	// 获取远程客户端TCP状态 IP port
	RemoteAddr() net.Addr
	// 发送数据给远程客户端
	SendMsg(uint32, []byte) error
}