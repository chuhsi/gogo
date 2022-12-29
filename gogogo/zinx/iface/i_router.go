package iface

type IRouter interface {
	//处理Conn业务之前的方法
	PreHandle(request IRequest)
	//处理Conn业务的主方法
	Handle(request IRequest)
	//处理Conn业务之的方法
	PostHandle(request IRequest)
}
