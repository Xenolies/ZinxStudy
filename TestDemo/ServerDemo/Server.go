package main

import "ZinxDemo01/Zinx/znet"

/**
基于 Zinx开发的服务端应用
*/

func main() {
	s := znet.NewServer("[Zinx]")

	s.Serve()
}
