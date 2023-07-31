package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// @Summary  GET请求的例子
// @Description  GET请求的例子描述
// @Tags         Swagger请求示例
// @Param        id  query  int  true  "Account ID"
// @Router       /swaggereget [get]
func Swaggerget(c *gin.Context) {
	fmt.Println("swagger注释示例")
}

// @Summary  POST请求的例子
// @Description  POST请求的例子描述
// @Tags         Swagger请求示例
// @Param        id  query  int  true  "Account ID"
// @Router       /swaggerpost [post]
func Swaggerpost(c *gin.Context) {
	fmt.Println("swagger注释示例")
}
