package inet

import "gogogo/zinx/iface"

type BaseRouter struct{}

/*

 */
//处理Conn业务之前的方法
func (br *BaseRouter) PreHandle(request iface.IRequest) {

}

//处理Conn业务的主方法
func (br *BaseRouter) Handle(request iface.IRequest) {

}

//处理Conn业务之的方法
func (br *BaseRouter) PostHandle(request iface.IRequest) {

}
