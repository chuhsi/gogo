package inet

import (
	"fmt"
	"gogogo/zinx/iface"
	"strconv"
)

type MsaManager struct {
	Apis map[uint32]iface.IRouter

	TaskQue []chan iface.IRequest

	WorkerPoolSize uint32
}

func New_MsgsHandle() *MsaManager {
	return &MsaManager{
		Apis:           make(map[uint32]iface.IRouter),
		TaskQue:        make([]chan iface.IRequest, 4),
		WorkerPoolSize: 4,
	}
}

/*

 */

func (m *MsaManager) SendMsgToTaskQue(request iface.IRequest) {
	// 将消息平均分配给Worker
	workerID := request.GetCurrentConnection().GetConnID() % m.WorkerPoolSize
	fmt.Println("Add ConnID = ", request.GetCurrentConnection().GetConnID(), "Request MsgID = ", request.GetCurrentMessage().GetMsgID(), "to WorkerID = ", workerID)
	// 将消息发送给对应的Worker的TaskQue
	m.TaskQue[workerID] <- request
}

// 启动一个Worker工作池(只能启动一个)
func (m *MsaManager) StartWorkerPool() {
	// 根据WorkerPoolSize 分别开启Worker 每个worker用一个go来开启
	var i uint32
	for i = 0; i < m.WorkerPoolSize; i++ {
		m.TaskQue[i] = make(chan iface.IRequest, 4)
		// 启动当前Worker， 阻塞消息从channel传递过来
		go m.StartOneWorker(i, m.TaskQue[i])
	}
}

// 启动一个Worker工作流程
func (m *MsaManager) StartOneWorker(workerID uint32, taskQue chan iface.IRequest) {
	fmt.Println("WorkerId = ", workerID, "is started")
	for {
		select {
		case request := <-taskQue:
			m.DoMsgHandle(request)
		}
	}
}
func (m *MsaManager) DoMsgHandle(request iface.IRequest) {
	handle, ok := m.Apis[request.GetCurrentMessage().GetMsgID()]
	if !ok {
		fmt.Println("Api msgID = ", request.GetCurrentMessage().GetMsgID(), "is not FUND! need register")
	}
	handle.PreHandle(request)
	handle.Handle(request)
	handle.PostHandle(request)
}

func (m *MsaManager) AddRouter(msgID uint32, router iface.IRouter) {
	// 1 判断 当前msg绑定的API处理方法是否存储
	if _, ok := m.Apis[msgID]; ok {
		panic("[Zinx] Repeqted API, msgID = " + strconv.Itoa(int(msgID)))
	}
	// 2 添加msgID和API的关系
	m.Apis[msgID] = router
	fmt.Println("Add API MsgID = ", msgID, "success")
}
