package xiface

// IXDataPack 对IXMessage进行封包、拆包的
type IXDataPack interface {
	// GetHeaderLength 获取包头的长度
	GetHeaderLength() uint32

	// Pack 封包，IXMessage->数据流
	Pack(m IXMessage) ([]byte, error)

	// UnPack 拆包，数据流->IXMessage
	UnPack([]byte) (IXMessage, error)
}
