package iface

type IServer interface {

	// 启动服务器
	Start()
	// 停止服务器
	Stop()
	// 运行服务器
	Serve()
	// 添加路由器
	AddRouter(uint32, IRouter)

	GetConnsMgr() IConnManager
}
