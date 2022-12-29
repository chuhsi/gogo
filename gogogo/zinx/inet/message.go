package inet

type Message struct {
	// 消息ID
	ID uint32
	// 消息长度
	DataLen uint32
	// 消息内容
	Data []byte
}

func New_Message(id uint32, data []byte) *Message {
	m := &Message{
		ID:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
	return m
}

/*

 */
// getter
func (m *Message) GetMsgID() uint32 {
	return m.ID
}
func (m *Message) GetMsgDataLen() uint32 {
	return m.DataLen
}
func (m *Message) GetMsgData() []byte {
	return m.Data
}

// setter
func (m *Message) SetMsgID(id uint32) {
	m.ID = id
}
func (m *Message) SetMsgDataLen(data_len uint32) {
	m.DataLen = data_len
}
func (m *Message) SetMsgData(data []byte) {
	m.Data = data
}
