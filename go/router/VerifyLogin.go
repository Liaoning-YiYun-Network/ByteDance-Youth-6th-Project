package router

import (
	"SkyLine/dao"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 定义一个中间件函数来进行用户认证
func verifyLogin(c *gin.Context) {

	fmt.Println("验证登录信息的请求地址:", c.Request.URL.String())
	token := c.Query("token")
	if token == "" {
		c.Next()
	}
	loginUser, err := dao.IsKeyExist(token)
	if err != nil || loginUser == false {
		fmt.Println("用户验证失败")
		c.Abort()
		return
	}
	// 用户已通过验证，继续路由处理
	c.Next()
}
