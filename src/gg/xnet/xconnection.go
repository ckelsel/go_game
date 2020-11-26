package xnet

import (
	"errors"
	"fmt"
	"gg/xiface"
	"io"
	"net"
)

// XConnection 抽象的客户端链接
type XConnection struct {
	// 当前链接的socket
	Conn *net.TCPConn

	// 链接的ID
	ConnID uint32

	// 当前链接的状态
	IsClosed bool

	// 告知当前链接已经退出的 channel
	ExitChan chan bool

	// 当前链接的路由
	Router xiface.IXMessageRouter
}

// StartReader 循环读取数据
func (c *XConnection) StartReader() {
	fmt.Println("Enter StartReader, connID ", c.ConnID)

	// defer在函数返回前执行
	defer fmt.Println("Leave StartReader connID", c.ConnID)
	defer c.Conn.Close()

	for {
		dp := NewXDataPack()

		// 第一次读取包头
		buf := make([]byte, dp.GetHeaderLength())
		_, err := io.ReadFull(c.Conn, buf)
		if err != nil {
			fmt.Println("ReadFull err ", err)
			break
		}

		// 解包读取ID和dataLength
		msg, err := dp.UnPack(buf)

		// 第二次读取数据，长度为dataLength
		if msg.GetLength() > 0 {
			data := make([]byte, msg.GetLength())
			_, err := io.ReadFull(c.Conn, data)
			if err != nil {
				fmt.Println("ReadFull err ", err)
				break
			}
			msg.SetData(data)
		}

		// 封装Request
		req := XRequest{
			Conn: c,
			msg:  msg,
		}

		go func(request xiface.IXRequest) {
			c.Router.Handle(request)
		}(&req)

	}
}

// Start 启动链接，开始工作
func (c *XConnection) Start() {
	fmt.Println("Start, connID ", c.ConnID)

	go c.StartReader()
}

// Stop 停止链接
func (c *XConnection) Stop() {
	fmt.Println("Stop, connID ", c.ConnID)

	if c.IsClosed {
		return
	}

	c.IsClosed = true

	// 关闭sock
	c.Conn.Close()

	// 释放资源
	close(c.ExitChan)
}

// GetTCPConnection 获取当前链接绑定的socket
func (c *XConnection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取当前链接的ID
func (c *XConnection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取远程客户端的TCP状态，IP Port
func (c *XConnection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// Send 发送数据
func (c *XConnection) Send(data []byte) error {
	_, err := c.Conn.Write(data)
	return err
}

// SendMsg 发送TLV消息
func (c *XConnection) SendMsg(id uint32, data []byte) error {
	if c.IsClosed {
		return errors.New("Connection has closed. ")
	}

	dp := NewXDataPack()

	bin, err := dp.Pack(NewXMessage(id, data))
	if err != nil {
		fmt.Println("Pack failed, ", id)
		return err
	}

	_, err = c.Conn.Write(bin)
	if err != nil {
		fmt.Println("Write failed, ", id)
		return err
	}
	return err
}

// NewConnection 初始化方法
func NewConnection(conn *net.TCPConn, connID uint32, router xiface.IXMessageRouter) *XConnection {
	connection := &XConnection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		IsClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return connection
}
