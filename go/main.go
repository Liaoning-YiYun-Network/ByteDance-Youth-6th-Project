package main

import (
	"SkyLine/cmd"
	"SkyLine/data"
	"runtime"
)

// @title 简易版抖音
// @version 0.0.1
// @description 简易版抖音开发接口文档目录
func main() {
	//判断宿主机的操作系统
	if runtime.GOOS == "windows" {
		data.OS = "windows"
	} else if runtime.GOOS == "linux" {
		data.OS = "linux"
	} else {
		data.OS = "mac"
	}
	data.Logger.Info("当前操作系统为：" + data.OS)
	//进行一系列的初始化操作
	cmd.Start()

}
