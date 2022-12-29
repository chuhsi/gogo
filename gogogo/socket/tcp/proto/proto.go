package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Encode(msg string) ([]byte, error) {
	
	var length = int32(len(msg))
	var pkg = new(bytes.Buffer)

	// 写入消息头
	err := binary.Write(pkg, binary.LittleEndian, length)

	if err != nil {
		return nil, err
	}

	// 写入消息体
	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	lengthByte, _ := reader.Peek(4)
	lengthBuff := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length)
	if err != nil {
		return " ", err
	}
	if int32(reader.Buffered()) < length + 4 {
		return " ", nil
	}

	//
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return " ", nil
	}	
	return string(pack[4:]), nil
}