package cmd

import (
	"SkyLine/config"
	"SkyLine/dao"
	"SkyLine/router"
	"fmt"
	"tencent.com/mmkv"
)

// Start 项目启动初始化各种配置
func Start() {
	//初始化读取配置文件
	config.InitConfig()
	//初始化路由
	router.InitRouter()
	//初始化数据库
	err := dao.InitMySql()
	if err != nil {
		fmt.Println("数据库初始化失败，请检查数据库配置是否正确，运行终止！")
		panic(err)
	}
	//初始化redis
	err = dao.InitRedis()
	if err != nil {
		fmt.Println("redis初始化失败，请检查redis配置是否正确，运行终止！")
		panic(err)
	}
	//初始化MMKV
	mmkv.InitializeMMKV("./data/mmkv")
}
