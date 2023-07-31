package main

import (
	"SkyLine/service"
	"github.com/gin-gonic/gin"
)

// @title 简易版抖音
// @version 0.0.1
// @description 简易版抖音开发接口文档目录
func main() {
	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
