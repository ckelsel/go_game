package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"xkm/xiface"
)

// 全局配置文件

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
	MaxConn uint32

	// 数据包的最大值
	MaxPacketSize uint32
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/gg.conf")
	if err != nil {
		fmt.Println("config/gg.conf not found")
		return
	}

	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func Init() {
	GlobalObject = &GlobalObj{
		TCPXServer:  nil,
		Host:        "0.0.0.0",
		Port:        8889,
		XServerName: "Good Game",

		MajorVersion:  "0",
		MinorVersion:  "4",
		PatchVersion:  "0",
		MaxConn:       1000,
		MaxPacketSize: 512,
	}

	GlobalObject.Reload()
}
