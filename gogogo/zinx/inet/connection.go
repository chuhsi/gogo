package inet

import (
	"errors"
	"fmt"
	"gogogo/zinx/iface"
	"io"
	"net"
)

type Connection struct {
	//
	TCP_Server iface.IServer
	// 当前连接
	Conn *net.TCPConn
	// 当前连接ID
	ConnID uint32
	// 当前连接状态
	IsClosed bool
	// 当前退出的连接
	Exit_Chan chan bool
	// 路由
	// Router iface.IRouter
	MsgsHandler iface.IMsgManager
	// 信息通道
	Msg_Chan chan []byte
}

func New_Connection(server iface.IServer, conn *net.TCPConn, connID uint32, msgsHandler iface.IMsgManager) *Connection {
	c := &Connection{
		TCP_Server: server,
		Conn:        conn,
		ConnID:      connID,
		IsClosed:    false,
		Exit_Chan:   make(chan bool, 1),
		MsgsHandler: msgsHandler,
		Msg_Chan:    make(chan []byte),
	}
	c.TCP_Server.GetConnsMgr().Add(c)
	return c
}
func (c *Connection) StartReader() {
	fmt.Println("[Zinx] Reader Goroutine is running ... ")
	defer fmt.Println("[Zinx] Reader exit", "ConnID = ", c.ConnID, "RemoteAddr is ", c.RemoteAddr().String())
	defer c.Stop()
	for {
		// 创建封包、拆包的对象
		dp := New_DataPack()
		// 读取客户端的Msg head 二进制流 8字节
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.GetConnetion(), headData)
		if err != nil {
			fmt.Println("[Zinx] Read msg head err ", err)
			break
		}
		// 拆包 得到包头信息，ID，Len
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack err ", err)
			break
		}
		// 根据Len，读取Data
		var data []byte
		if msg.GetMsgDataLen() > 0 {
			data = make([]byte, msg.GetMsgDataLen())
			if _, err := io.ReadFull(c.GetConnetion(), data); err != nil {
				fmt.Println("[Zinx] Read msg head err ", err)
				break
			}
		}
		msg.SetMsgData(data)
		request := &Request{
			Conn: c,
			Msg:  msg,
		}
		if 4 > 0 {
			go c.MsgsHandler.SendMsgToTaskQue(request)
		} else {
			go c.MsgsHandler.DoMsgHandle(request)
		}
	}
}
func (c *Connection) StartWriter() {
	fmt.Println("[Zinx] Writer Goroutine is running ... ")
	defer fmt.Println("[Zinx] Writer exit!", c.RemoteAddr().String())
	for {
		select {
		case data := <-c.Msg_Chan:
			// 把数据写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send data err", err)
				return
			}
		case <-c.Exit_Chan:
			// 代表Reader已经退出，此时也要退出
			return
		}
	}
}

/*

 */
// 启动连接
func (c *Connection) Start() {
	fmt.Println("[Zinx] Curren Conn Start .... ConnID = ", c.ConnID)
	go c.StartReader()
	go c.StartWriter()
}

// 停止连接
func (c *Connection) Stop() {
	fmt.Println("[Zinx] Current Conn Stop ... ConnID = ", c.ConnID)
	if c.IsClosed {
		return
	}
	c.IsClosed = true
	c.Conn.Close()
	c.Exit_Chan <- true
	c.TCP_Server.GetConnsMgr().Remove(c)
	close(c.Exit_Chan)
	close(c.Msg_Chan)
}

// 获取当前的连接
func (c *Connection) GetConnetion() *net.TCPConn {
	return c.Conn
}

// 获取当前连接模块的连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端TCP状态 IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据给远程客户端
func (c *Connection) SendMsg(msdID uint32, data []byte) error {
	if c.IsClosed {
		return errors.New("[Zinx] Connection is closed when send msg")
	}
	// 将数据进行封包 Len|Id|Data
	dp := New_DataPack()
	binaryMsg, err := dp.Pack(New_Message(msdID, data))
	if err != nil {
		return errors.New("[Zinx] dp.Pack err")
	}
	// 将数据发送给MsgChan
	c.Msg_Chan <- binaryMsg
	return nil
}
