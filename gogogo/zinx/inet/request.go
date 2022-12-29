package inet

import "gogogo/zinx/iface"

type Request struct {
	Conn iface.IConnection
	Msg  iface.IMessage
}

/*

 */
// 获得当前连接
func (r *Request) GetCurrentConnection() iface.IConnection{
	return r.Conn
}

// 获得当前信息
func (r *Request) GetCurrentMessage() iface.IMessage{
	return r.Msg
}
