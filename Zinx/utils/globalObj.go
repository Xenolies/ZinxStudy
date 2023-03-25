package utils

import (
	"ZinxStudy/Zinx/ziface"
	"encoding/json"
	"fmt"
	"os"
)

/*
全局配置模块
在 "服务器程序/conf/zinx.json"中写入配置
将框架中的硬代码.使用 globalObj进行替换
*/

type GlobalObj struct {
	/*
		Server Info
	*/
	TcpServer ziface.IServer // 当前 Zinx Server 全局对选哪个
	Host      string         // 当前服务器监听IP
	TcpPort   int            // 当前服务器端口
	Name      string         // 按当前服务器名称

	/*
		Zinx Info
	*/
	Version          string // 当前 Zinx 版本号
	MaxConn          int    // 最大链接数量
	MaxPackageSize   uint32 // 当前 Zinx 数据包最大值
	WorkerPoolSize   uint32 // 业务工作数量 worker 数量
	MaxWorkerTaskLen uint32 // worker消息队列的任务数量最大值

}

var GlobalObject *GlobalObj

// 提供一个 init 方法,提供一个初始的默认值
func init() {
	GlobalObject = &GlobalObj{
		TcpServer:        nil,
		Host:             "127.0.0.1",
		TcpPort:          8899,
		Name:             "Zinx Server",
		Version:          "v0.8",
		MaxConn:          2,
		MaxPackageSize:   512,
		WorkerPoolSize:   5,
		MaxWorkerTaskLen: 10,
	}

	// 从 conf/zinx.json 加载用户自定义参数
	GlobalObject.Reload()

}

// Reload 从 zinx.json中加载自定义参数
func (g *GlobalObj) Reload() {
	data, err := os.ReadFile("conf/zinx.json")
	if err != nil {
		//panic(err)
		fmt.Println(err)
		return
	}

	// 将 JSON 文件解析到 GlobalObj
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		return
	}

}
