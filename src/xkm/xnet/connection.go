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
	isClosed bool

	// 当前链接的回调函数
	handleAPI xiface.HandleFunc

	// 告知当前链接已经退出的 channel
	ExitChan chan bool
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
			break;
		}

		err = c.handleAPI(c.Conn, buf, cnt)
		if err != nil {
			fmt.Println("handleAPI err ", err);
			continue;
		}
	}
}

// 启动链接，开始工作
func (c *Connection) Start(){
	fmt.Println("Start, connID ", c.ConnID)

	go c.StartReader()
}

// 停止链接
func (c *Connection) Stop() {
	fmt.Println("Stop, connID ", c.ConnID)

	if c.isClosed {
		return
	}

	// 关闭sock
	c.Conn.Close()

	// 释放资源
	close(c.ExitChan)
}


// 获取当前链接绑定的socket
func (c *Connection)  GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前链接的ID
func (c *Connection)  GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的TCP状态，IP Port
func (c *Connection)  RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// 发送数据
func (c *Connection) Send(data []byte) error {
	_, err := c.Conn.Write(data)
	return err;
}

// 初始化方法

func NewConnection(conn *net.TCPConn, connID uint32, callback xiface.HandleFunc) *Connection {
	connection := &Connection {
		Conn: conn,
		ConnID: connID,
		handleAPI:callback,
		isClosed:false,
		ExitChan:make(chan bool, 1),
	}
	return connection
}