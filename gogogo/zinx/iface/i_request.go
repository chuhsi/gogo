package iface

type IRequest interface {
	// 获得当前连接
	GetCurrentConnection() IConnection
	// 获得当前信息
	GetCurrentMessage() IMessage
}