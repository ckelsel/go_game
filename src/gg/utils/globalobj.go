package utils

import (
	"encoding/json"
	"fmt"
	"gg/xiface"
	"io/ioutil"
)

// GlobalObj 全局配置文件
type GlobalObj struct {
	//
	// 基础配置
	//

	// 全局XServer对象
	TCPXServer xiface.IXServer

	// 监听的IP
	Host string

	// 监听的TCP端口
	Port int

	// 服务器名称
	XServerName string

	//
	// 高级配置
	//

	// 主版本号
	MajorVersion string

	// 次版本号
	MinorVersion string

	// 补丁版本
	PatchVersion string

	// 允许的最大连接数
	MaxConn int

	// 数据包的最大值
	MaxPacketSize uint32

	// 工作池的大小
	WorkerPoolSize uint32

	// 消息队列的大小（框架限制)
	MaxTaskQueueSize uint32
}

// GlobalObject 全局配置实例
var GlobalObject *GlobalObj

// Reload 读取配置文件
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/gg.conf")
	if err != nil {
		fmt.Println("conf/gg.conf not found")
		return
	}

	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

// Init 全局配置初始化
func Init() {
	GlobalObject = &GlobalObj{
		TCPXServer:  nil,
		Host:        "0.0.0.0",
		Port:        8889,
		XServerName: "Good Game",

		MajorVersion:     "0",
		MinorVersion:     "4",
		PatchVersion:     "0",
		MaxConn:          1024,
		MaxPacketSize:    512,
		WorkerPoolSize:   10,
		MaxTaskQueueSize: 1024,
	}

	GlobalObject.Reload()
}
