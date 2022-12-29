package inet

import (
	"bytes"
	"encoding/binary"
	// "errors"
	"gogogo/zinx/iface"
)

// 封包、拆包具体模块
type DataPack struct{}

func New_DataPack() *DataPack {
	return &DataPack{}
}

// 获取包头的长度
func (dp *DataPack) GetHeadLen() uint32 {
	// Id uint32 + DateLen uint32 = 8字节
	return 8
}
// Pack 封包方法(压缩数据)
func (dp *DataPack) Pack(msg iface.IMessage) ([]byte, error) {
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})
	//写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgDataLen()); err != nil {
		return nil, err
	}
	//写msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}
//Unpack 拆包方法(解压数据)
func (dp *DataPack) Unpack(binaryData []byte) (iface.IMessage, error) {
	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)
	//只解压head的信息，得到dataLen和msgID
	msg := &Message{}
	//读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}
	//读msgID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}
	//判断dataLen的长度是否超出我们允许的最大包长度
	// if utils.GlobalObject.MaxPackageSize > 0 && msg.DateLen > utils.GlobalObject.MaxPackageSize {
	// 	return nil, errors.New("too large msg data received")
	// }
	//这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
	return msg, nil
}
