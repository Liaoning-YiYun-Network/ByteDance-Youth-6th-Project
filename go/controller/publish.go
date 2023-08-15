package controller

import (
	"SkyLine/entity"
	"SkyLine/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	entity.Response
	VideoList []entity.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	fileName := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, fileName)
	saveFile := filepath.Join("./public/", finalName)
	// 使用os.Stat函数检查文件是否存在
	//_, err := os.Stat(saveFile)
	//if err == nil {
	//	fmt.Println("文件存在")
	//} else if os.IsNotExist(err) {
	//	fmt.Println("文件不存在")
	//} else {
	//	fmt.Println("发生了其他错误：", err)
	//}
	fileContent, err := ioutil.ReadFile(saveFile)
	if err != nil {
		fmt.Println("读取文件内容时发生错误：", err)
		return
	}
	// 调用上传函数
	err = service.UploadFile(fileName, fileContent, 2)
	if err != nil {
		fmt.Println("上传文件时发生错误：", err)
	} else {
		fmt.Println("文件上传成功！")
	}
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: entity.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
