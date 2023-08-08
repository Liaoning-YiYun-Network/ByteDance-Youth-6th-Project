package cmd

import (
	"SkyLine/config"
	"SkyLine/dao"
	"SkyLine/router"
	"SkyLine/service"
	"fmt"
)

// Start 项目启动初始化各种配置
func Start() {
	//初始化读取配置文件
	config.InitConfig()

	defer dao.CloseMySql()
	defer dao.CloseRedis()
	defer service.CloseTOS()

	//初始化数据库
	err := dao.InitMySql()
	if err != nil {
		fmt.Println("数据库初始化失败，请检查数据库配置是否正确，运行终止！")
		panic(err)
	}
	fmt.Println("初始化数据库成功")

	//初始化redis
	err = dao.InitRedis()
	if err != nil {
		fmt.Println("Redis初始化失败，请检查Redis配置是否正确，运行终止！")
		panic(err)
	}
	fmt.Println("初始化Redis成功")

	err = service.InitTOS()
	if err != nil {
		fmt.Println("TOS初始化失败，请检查TOS配置是否正确，运行终止！")
		panic(err)
	}
	fmt.Println("初始化TOS成功")

	//将初始化路由放入最后，否则初始化路由后面的代码都不会执行
	//初始化路由
	router.InitRouter()
}
