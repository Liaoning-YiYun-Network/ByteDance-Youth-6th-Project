# ！！提示！！
该项目随着第六届字节青训营结束，不再维护，已被 Archive

### 运行该代码所必需的环境：

### 一、安装ffmpeg

因为本项目用到了抽帧制作图片，所以需要安装ffmpeg。

安装参考该地址：http://t.csdn.cn/WVYWj

###  二、配置swagger使用文档

接口文档地址：[Swagger UI](http://localhost:9090/swagger/index.html)

新增的接口，都加上以下注释将该接口添加到swagger接口文档

```go
// @Summary 接口名称
// @Description 接口详情描述
// @Param 参数名称 以什么类型传过来的 参数类型 是否必须 "描述"   （有多个参数需要加多个@Param）
// @Tags 标签分类
// @Router 路由地址 [请求类型]
```

示例：`swagger_example.go` 文件

```go
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
```

接口写好后，第一次生成swagger文档需要在控制台执行如下命令，只用执行一次即可，目的是给你的电脑加上swagger的环境变量。

```shell
go install github.com/swaggo/swag/cmd/swag@latest
```

使用如下命令查看swagger是否安装好,若出现版本号则安装成功

```shell
swag -v
```

接下来进行如下操作来生成接口文档

```shell
swag init
```

启动项目，然后就可以访问上面提到的接口文档地址来查看接口文档了



### 三、启动项目

1、首先转移go项目目录

```shell
cd go
```

2、安装所需要的依赖

```shell
go mod download
```

3、启动项目
