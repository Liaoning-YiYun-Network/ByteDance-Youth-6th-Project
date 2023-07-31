package main

import (
	"SkyLine/cmd"
	"SkyLine/service"
)

// @title 简易版抖音
// @version 0.0.1
// @description 简易版抖音开发接口文档目录
func main() {
	go service.RunMessageServer()
	//进行一系列的初始化操作
	cmd.Start()

}
