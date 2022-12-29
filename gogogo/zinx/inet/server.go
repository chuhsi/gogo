package inet

import (
	"fmt"
	"gogogo/zinx/iface"
	"net"
)

type Server struct {
	// 服务器名称
	Name string
	// IP版本
	IPVersion string
	// IP
	IP string
	// 端口
	Port int32
	// 路由
	// Router iface.IRouter
	MsgsHandler iface.IMsgManager
	// 连接管理
	ConnsManager iface.IConnManager
}

// 初始化Server模块
func New_Server() *Server {
	s := &Server{
		Name:         "zinx",
		IPVersion:    "tcp4",
		IP:           "127.0.0.1",
		Port:         9090,
		MsgsHandler:  New_MsgsHandle(),
		ConnsManager: New_ConnManager(),
	}
	return s
}

/*
	IServer实现
*/
func (s *Server) GetConnsMgr() iface.IConnManager {
	return s.ConnsManager
}

// 添加路由器
func (s *Server) AddRouter(msgID uint32, router iface.IRouter) {
	fmt.Printf("[Zinx] Add %d Router Success\n", msgID)
	s.MsgsHandler.AddRouter(msgID, router)
}

// 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name: %s, Listen At IP: %s, Port: %d\n", s.Name, s.IP, s.Port)
	go func() {
		s.MsgsHandler.StartWorkerPool()
		tcpAddr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println(err)
			return
		}
		listen, err := net.ListenTCP(s.IPVersion, tcpAddr)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("[Zinx] Server Success: %s, Succeed Listenning\n", s.Name)
		// 连接标识
		var connID uint32 = 0
		for {
			conn, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println("listen.Accept err", err)
				continue
			}
			dealConn := New_Connection(s, conn, connID, s.MsgsHandler)
			connID++
			dealConn.Start()
			// buf := make([]byte, 512)
			// n, err := conn.Read(buf)
			// if err != nil {
			// 	fmt.Println("conn.Read err", err)
			// 	continue
			// }
			// fmt.Printf("[Zinx] Client Send data: %s", buf[:n])
			// _, err = conn.Write(buf[:n])
			// if err != nil {
			// 	fmt.Println("conn.Write err", err)
			// 	continue
			// }
		}
	}()
}

// 停止服务器
func (s *Server) Stop() {
	fmt.Println("[Zinx] STOP Server ...")
	s.ConnsManager.ClearConn()
}

// 运行服务器
func (s *Server) Serve() {
	s.Start()

	select {}
}
