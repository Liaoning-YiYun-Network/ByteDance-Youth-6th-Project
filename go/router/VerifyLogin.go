package router

import (
	"github.com/gin-gonic/gin"
)

// 定义一个中间件函数来进行用户认证
func verifyLogin(c *gin.Context) {
	c.Next()
}
