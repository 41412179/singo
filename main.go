package main

import (
	"accelerator/conf"
	"accelerator/server"
)

func main() {
	// 从配置文件读取配置
	new(conf.Conf).Init()

	// 装载路由
	r := server.NewRouter()
	r.Run(":3000")
}
