package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/volcengine/ve-tos-golang-sdk/v2/tos"
)

// TOSFileType 上传文件类型
type TOSFileType int

// VIDEO_COVER 视频封面
const VIDEO_COVER TOSFileType = 1

// VIDEO 视频
const VIDEO TOSFileType = 2

// AVATAR 头像
const AVATAR TOSFileType = 3

// BACKGROUND 背景
const BACKGROUND TOSFileType = 4

var AccessKey = viper.GetString("tos.accessKey")
var SecretKey = viper.GetString("tos.secretKey")
var BucketName = viper.GetString("tos.bucketName")

// client TOS客户端
var client *tos.ClientV2

// InitTOS 初始化TOS
func InitTOS() error {
	var endpoint string
	if viper.GetBool("tos.useInner") {
		endpoint = viper.GetString("tos.endpoint-inside")
	} else {
		endpoint = viper.GetString("tos.endpoint-outside")
	}
	credentials := tos.NewStaticCredentials(AccessKey, SecretKey)
	var err error
	client, err = tos.NewClientV2(
		endpoint,
		tos.WithCredentials(credentials),
		tos.WithRegion(viper.GetString("tos.region")),
		tos.WithMaxRetryCount(3),
		tos.WithEnableVerifySSL(true))
	return err
}

// CloseTOS 关闭TOS
func CloseTOS() {
	//关闭tos连接
	client.Close()
}

// UploadFile 上传文件
// fileName 文件名
// fileContent 文件内容
// fileType 文件类型
// 返回requestID和错误信息
func UploadFile(fileName string, fileContent []byte, fileType TOSFileType) (requestID string, err error) {
	var pathPrefix string
	switch fileType {
	case VIDEO_COVER:
		pathPrefix = "video_covers/"
	case VIDEO:
		pathPrefix = "videos/"
	case AVATAR:
		pathPrefix = "avatars/"
	case BACKGROUND:
		pathPrefix = "backgrounds/"
	}
	var key = pathPrefix + fileName
	var ctx = context.Background()
	var reader = bytes.NewReader(fileContent)
	output, err := client.PutObjectV2(ctx, &tos.PutObjectV2Input{
		PutObjectBasicInput: tos.PutObjectBasicInput{
			Bucket: BucketName,
			Key:    key,
		},
		Content: reader,
	})
	if output.StatusCode != 200 {
		fmt.Println("上传文件失败，错误码：", output.StatusCode, "，错误信息：", err)
		return output.RequestID, err
	}
	return output.RequestID, err
}
