package xnet

import (
	"fmt"
	"gg/utils"
	"gg/xiface"
	"net"
)

type XServer struct {
	// 服务器名称
	Name string

	// 服务器ip

	IP string
	// 服务器的端口

	Port int

	// 服务器的IPv4 Ipv6
	IPVersion string

	// 路由，处理所有的connection
	Router xiface.IXRouter
}

// 启动服务器
func (s *XServer) Start() {
	fmt.Printf("XServer %s, Version %s.%s.%s, MaxConn %d, MaxPacketSize %d\n",
		utils.GlobalObject.XServerName,
		utils.GlobalObject.MajorVersion,
		utils.GlobalObject.MinorVersion,
		utils.GlobalObject.PatchVersion,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPacketSize)

	fmt.Printf("Listen on IP %s, Port %d, start\n", s.IP, s.Port)

	// 1. 获取TCP的addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error :", err)
		return
	}

	// 2. 监听
	listen, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("Listen failed, ", err)
		return
	}

	fmt.Printf("Listen on IP %s, Port %d, success\n", s.IP, s.Port)

	var connID uint32
	connID = 0

	// 3. 等待客户端连接
	for {
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err", err)
			continue
		}

		fmt.Println("player incoming")

		c := NewConnection(conn, connID, s.Router)
		connID++

		go c.Start()
	}
}

// 停止服务器
func (s *XServer) Stop() {
}

// 运行服务器
func (s *XServer) Run() {
	s.Start()
}

func (s *XServer) AddRouter(router xiface.IXRouter) {
	fmt.Println("AddRouter success")
	s.Router = router
}

func NewXServer() xiface.IXServer {
	utils.Init()

	s := &XServer{
		Name:      utils.GlobalObject.XServerName,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.Port,
		Router:    nil,
	}

	utils.GlobalObject.TCPXServer = s

	return s
}
