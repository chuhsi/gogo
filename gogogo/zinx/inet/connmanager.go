package inet

import (
	"errors"
	"fmt"
	"gogogo/zinx/iface"
	"log"
	"sync"
)

type ConnManager struct {
	Conns    map[uint32]iface.IConnection
	ConnLock sync.RWMutex
}

func New_ConnManager() *ConnManager {
	return &ConnManager{
		Conns: make(map[uint32]iface.IConnection),
	}
}

/*

 */
func (cm *ConnManager) Add(conn iface.IConnection) {
	cm.ConnLock.Lock()
	defer cm.ConnLock.Unlock()
	// 将conn加入到ConnManager中
	cm.Conns[conn.GetConnID()] = conn
	log.Println("connection add to ConnManager successfully: conn num = ", cm.Len())
}

func (cm *ConnManager) Remove(conn iface.IConnection) {
	cm.ConnLock.Lock()
	defer cm.ConnLock.Unlock()

	delete(cm.Conns, conn.GetConnID())
	log.Println("connID = ", conn.GetConnID(), "remove from ConnManager successfully: conn num = ", cm.Len())
}

func (cm *ConnManager) Get(connID uint32) (iface.IConnection, error) {
	cm.ConnLock.RLock()
	defer cm.ConnLock.RUnlock()
	if conn, ok := cm.Conns[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("[Zinx] Connection don't FOUND!!!")
	}
}

func (cm *ConnManager) Len() int {
	return len(cm.Conns)
}

func (cm *ConnManager) ClearConn() {
	cm.ConnLock.Lock()
	defer cm.ConnLock.Unlock()
	for connID, conn := range cm.Conns {
		conn.Stop()
		delete(cm.Conns, connID)
	}
	fmt.Printf("[Zinx] Clear All Connections success! Conn num = %d", cm.Len())
}
