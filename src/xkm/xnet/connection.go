package xnet

import (
	"fmt"
	"net"
	"xkm/xiface"
)

type Connection struct {
	// 当前链接的socket
	Conn *net.TCPConn

	// 链接的ID
	ConnID uint32

	// 当前链接的状态
	IsClosed bool

	// 告知当前链接已经退出的 channel
	ExitChan chan bool

	// 当前链接的路由
	Router xiface.IRouter
}

func (c *Connection) StartReader() {
	fmt.Println("Enter StartReader, connID ", c.ConnID)

	// defer在函数返回前执行
	defer fmt.Println("Leave StartReader connID", c.ConnID)
	defer c.Conn.Close()

	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("Read err ", err)
			break
		}

		// 封装Request
		req := XRequest{
			Conn: c,
			Data: buf[:cnt],
		}

		go func(request xiface.IXRequest) {
			c.Router.PreHandle(request)

			c.Router.Handle(request)

			c.Router.PostHandle(request)
		}(&req)

	}
}

// 启动链接，开始工作
func (c *Connection) Start() {
	fmt.Println("Start, connID ", c.ConnID)

	go c.StartReader()
}

// 停止链接
func (c *Connection) Stop() {
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

// 获取当前链接绑定的socket
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前链接的ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的TCP状态，IP Port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据
func (c *Connection) Send(data []byte) error {
	_, err := c.Conn.Write(data)
	return err
}

// 初始化方法

func NewConnection(conn *net.TCPConn, connID uint32, router xiface.IRouter) *Connection {
	connection := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		IsClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return connection
}
