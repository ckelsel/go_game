package xnet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"gg/utils"
	"gg/xiface"
)

// XDataPack 封包、拆包
type XDataPack struct {
}

// NewXDataPack 构造函数
func NewXDataPack() *XDataPack {
	return &XDataPack{}
}

// GetHeaderLength 获取包头的长度
func (dp *XDataPack) GetHeaderLength() uint32 {
	// L uint32
	// T uint32
	// V []byte
	return 8
}

// Pack 封包，格式化的IXMessage->二进制流
func (dp *XDataPack) Pack(m xiface.IXMessage) ([]byte, error) {
	// 创建一个缓冲区
	buf := bytes.NewBuffer([]byte{})

	// 写入Length
	err := binary.Write(buf, binary.LittleEndian, m.GetLength())
	if err != nil {
		return nil, err
	}

	// 写入ID
	err = binary.Write(buf, binary.LittleEndian, m.GetID())
	if err != nil {
		return nil, err
	}

	// 写入Data
	err = binary.Write(buf, binary.LittleEndian, m.GetData())
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// UnPack 拆包，二进制流->IXMessage
// 只读取Header数据（L，T），data需要再次读取
func (dp *XDataPack) UnPack(data []byte) (xiface.IXMessage, error) {
	msg := &XMessage{}

	buf := bytes.NewReader(data)

	// 读取Length
	err := binary.Read(buf, binary.LittleEndian, &msg.Length)
	if err != nil {
		return nil, err
	}

	if msg.Length > utils.GlobalObject.MaxPacketSize {
		return nil, errors.New("msg length > MaxPacketSize")
	}

	// 读取ID
	err = binary.Read(buf, binary.LittleEndian, &msg.ID)
	if err != nil {
		return nil, err
	}

	// 未读取Data

	return msg, nil
}
