package xiface

type IXMessage interface {
	// 获取消息ID
	GetID() uint32
	// 获取消息长度
	GetLength() uint32
	// 获取消息内容
	GetData() []byte

	// 设置消息ID
	SetID(uint32)
	// 设置消息长度
	SetLength(uint32)
	// 设置消息内容
	SetData([]byte)
}