package cmd

import (
	"SkyLine/config"
	"SkyLine/dao"
	"SkyLine/router"
	"fmt"
)

// Start 项目启动初始化各种配置
func Start() {
	//初始化读取配置文件
	config.InitConfig()

	defer dao.CloseMySql()
	//defer dao.CloseRedis()

	//初始化数据库
	err := dao.InitMySql()
	if err != nil {
		fmt.Println("数据库初始化失败，请检查数据库配置是否正确，运行终止！")
		panic(err)
	}
	fmt.Print("初始化数据库成功")

	//初始化redis
	//err = dao.InitRedis()
	//if err != nil {
	//	fmt.Println("redis初始化失败，请检查redis配置是否正确，运行终止！")
	//	panic(err)
	//}

	//将初始化路由放入最后，否则初始化路由后面的代码都不会执行
	//初始化路由
	router.InitRouter()
}
