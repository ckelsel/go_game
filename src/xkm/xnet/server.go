package xnet


import "xkm/xiface"
import "fmt"
import "net"

type Server struct {
    // 服务器名称
    Name string

    // 服务器ip

    IP string
    // 服务器的端口

    Port int

    // 服务器的IPv4 Ipv6
    IPVersion string
}

func callback(c *net.TCPConn, data []byte, cnt int) error {
    fmt.Println("callback called")

    _, err := c.Write(data[:cnt])

    return err
}

// 启动服务器
func (s *Server) Start() {
    fmt.Printf("Server %s Listen on IP %s, Port %d, start\n", s.Name, s.IP, s.Port)

    // 1. 获取TCP的addr
    addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
    if err != nil {
        fmt.Println("resolve tcp addr error :", err);
        return
    }

    // 2. 监听
    listen, err := net.ListenTCP(s.IPVersion, addr)
    if err != nil{
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
        
        c := NewConnection(conn, connID, callback)
        connID++

        go c.Start()
    }
}


// 停止服务器
func (s *Server) Stop() {
}


// 运行服务器
func (s *Server) Run() {
    s.Start()
}


func NewServer(name string) xiface.IServer {
    s := &Server {
        Name: name,
        IPVersion: "tcp4",
        IP:"0.0.0.0",
        Port:8889,
    }

    return s
}
