package iface

type IMessage interface {
	// getter
	GetMsgID() uint32
	GetMsgDataLen() uint32
	GetMsgData() []byte
	// setter
	SetMsgID(uint32)
	SetMsgDataLen(uint32)
	SetMsgData([]byte)
}
