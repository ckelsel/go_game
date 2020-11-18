package xnet

// XMessage 定义客户端、服务端的TLV消息格式
type XMessage struct {
	// 获取消息ID
	ID uint32

	// 获取消息长度
	Length uint32

	// 获取消息内容
	Data []byte
}

// NewXMessage 创建一个消息
func NewXMessage(id uint32, data []byte) *XMessage {
	return &XMessage{
		ID:     id,
		Length: uint32(len(data)),
		Data:   data,
	}
}

// GetID 获取消息ID
func (m *XMessage) GetID() uint32 {
	return m.ID
}

// GetLength 获取消息长度
func (m *XMessage) GetLength() uint32 {
	return m.Length
}

// GetData 获取消息内容
func (m *XMessage) GetData() []byte {
	return m.Data
}

// SetID 设置消息ID
func (m *XMessage) SetID(id uint32) {
	m.ID = id
}

// SetLength 设置消息长度
func (m *XMessage) SetLength(length uint32) {
	m.Length = length
}

// SetData 设置消息内容
func (m *XMessage) SetData(data []byte) {
	m.Data = data
}
