package xnet

import (
	"fmt"
	"gg/utils"
	"gg/xiface"
	"net"
)

// XServer 服务端
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
	Router xiface.IXMessageRouter

	// 连接管理器
	ConnectionManager xiface.IXConnectionManager

	// 连接建立后的回调
	OnConnectionStartCallBack func(conn xiface.IXConnection)

	// 连接断开前的回调
	OnConnectionStopCallBack func(conn xiface.IXConnection)
}

// AddOnConnectionStartCallBack 设置连接建立后的回调
func (s *XServer) AddOnConnectionStartCallBack(onstart func(conn xiface.IXConnection)) {
	s.OnConnectionStartCallBack = onstart
}

// AddOnConnectionStopCallBack 设置连接断开前的回调
func (s *XServer) AddOnConnectionStopCallBack(onstop func(conn xiface.IXConnection)) {
	s.OnConnectionStopCallBack = onstop
}

// OnConnectionStart 连接建立后的回调
func (s *XServer) OnConnectionStart(conn xiface.IXConnection) {
	if s.OnConnectionStartCallBack != nil {
		s.OnConnectionStartCallBack(conn)
	}
}

// OnConnectionStop 连接断开前的回调
func (s *XServer) OnConnectionStop(conn xiface.IXConnection) {
	if s.OnConnectionStopCallBack != nil {
		s.OnConnectionStopCallBack(conn)
	}
}

// GetConnectionManager 获取连接管理器
func (s *XServer) GetConnectionManager() (cm xiface.IXConnectionManager) {
	return s.ConnectionManager
}

// Start 启动服务器
func (s *XServer) Start() {
	fmt.Printf("XServer %s, Version %s.%s.%s, MaxConn %d, MaxPacketSize %d\n",
		utils.GlobalObject.XServerName,
		utils.GlobalObject.MajorVersion,
		utils.GlobalObject.MinorVersion,
		utils.GlobalObject.PatchVersion,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPacketSize)

	fmt.Printf("Listen on IP %s, Port %d, start\n", s.IP, s.Port)

	s.Router.StartWorkerPool()

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

		if s.ConnectionManager.Length() >= utils.GlobalObject.MaxConn {
			fmt.Println("too many connections. MaxConn ", utils.GlobalObject.MaxConn)
			// TODO: 发送超出最大连接数的包
			conn.Close()
			continue
		}

		fmt.Println("player incoming")

		c := NewConnection(s, conn, connID, s.Router)

		connID++

		go c.Start()
	}
}

// Stop 停止服务器
func (s *XServer) Stop() {
	s.ConnectionManager.Clear()
}

// Run 运行服务器
func (s *XServer) Run() {
	s.Start()
}

// AddRouter 添加消息处理路由
func (s *XServer) AddRouter(msgid uint32, router xiface.IXRouter) {
	fmt.Println("AddRouter success")
	s.Router.AddRouter(msgid, router)
}

// NewXServer 初始化
func NewXServer() xiface.IXServer {
	utils.Init()

	s := &XServer{
		Name:                      utils.GlobalObject.XServerName,
		IPVersion:                 "tcp4",
		IP:                        utils.GlobalObject.Host,
		Port:                      utils.GlobalObject.Port,
		Router:                    NewXMessageRouter(),
		ConnectionManager:         NewXConnectionManager(),
		OnConnectionStartCallBack: nil,
		OnConnectionStopCallBack:  nil,
	}

	utils.GlobalObject.TCPXServer = s

	return s
}
