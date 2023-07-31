package cmd

import (
	"SkyLine/config"
	"SkyLine/router"
)

// 项目启动初始化各种配置
func Start() {
	//初始化读取配置文件
	config.InitConfig()
	//初始化路由
	router.InitRouter()

}
