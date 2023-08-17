package controller

import (
	"SkyLine/entity"
	"SkyLine/service"
	"SkyLine/util"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

type VideoListResponse struct {
	entity.Response
	VideoList []entity.DouyinVideo `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	//	if _, exist := usersLoginInfo[token]; !exist {
	//		c.JSON(http.StatusOK, entity.Response{
	//StatusCode: 1,
	//			StatusMsg:  "User doesn't exist",
	//		})
	//		return
	//	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	// 获取文件大小
	fileSize := data.Size
	// 设置文件大小限制，例如 500MB
	maxFileSize := int64(500 * 1024 * 1024) // 500MB
	if fileSize > maxFileSize {
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  "File size exceeds the limit",
		})
		return
	}
	file, err := data.Open()
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  "Failed to open file",
		})
		return
	}
	defer file.Close()
	fileContent, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  "Failed to read file",
		})
		return
	}
	fileName := data.Filename
	// 调用上传函数上传视频
	//函数的第二个参数需要是 []byte 类型
	err = service.UploadFile(fileName, fileContent, service.VIDEO)
	if err != nil {
		fmt.Println("上传文件时发生错误：", err)
		// 处理错误
	} else {
		fmt.Println("文件上传成功！")
		// 处理成功
	}
	//这块是怎么回事？
	//token := c.Query("token")
	//user, err := dao.GetRedis(token)
	user := usersLoginInfo[token]
	//id, _ := strconv.Atoi(c.Query("user_id"))
	//这块是怎么回事？
	// 调用UUID函数生成带横杠的UUID
	videoUUID, err := util.UUID()
	if err != nil {
		fmt.Println("Error generating UUID:", err)
		return
	}
	newVideoName := fmt.Sprintf("%d-%s-%s.mp4", user.Id, videoUUID, fileName) //产生问题了！！！
	videoUrl := "https://tos.eyunnet.com/" + "videos/" + newVideoName
	coverBytes, _ := ReadFrameAsJpeg(videoUrl)
	// 调用上传函数上传视频封面

	err = service.UploadFile(newVideoName, coverBytes, service.VIDEO_COVER)
	if err != nil {
		fmt.Println("上传文件时发生错误：", err)
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 1,
			StatusMsg:  newVideoName + " failed in uploading",
		})
		return
	} else {
		fmt.Println("文件上传成功！")
		c.JSON(http.StatusOK, entity.Response{
			StatusCode: 0,
			StatusMsg:  newVideoName + " uploaded successfully",
		})
	}
}

// 截取视频封面
func ReadFrameAsJpeg(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), err
}

// @Summary  获取某一用户的所发布的搜游视频
// @Description  这个接口，会根据用户id去查询该用户发布的所有的视频
// @Tags         视频相关接口
// @Param        userid  query  int64  ture  "用户的userid"
// @Param        Token  query  string  ture  "该参数只有在用户登录状态下进行设置"
// @Router       /publish/list [get]
func PublishList(c *gin.Context) {
	userid, interr := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if interr != nil {
		c.JSON(http.StatusInternalServerError, entity.FeedResponse{
			Response:  entity.Response{StatusCode: 1, StatusMsg: "参数转换失败"},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
	}
	publishListRequest := &entity.PublishListRequest{userid, c.Query("token")}
	video, err := service.SelectVideoListByUserId(publishListRequest)
	if err != nil {
		fmt.Printf("视频获取出错:%v\n", err)
		c.JSON(http.StatusInternalServerError, entity.FeedResponse{
			Response:  entity.Response{StatusCode: 1, StatusMsg: "获取视频列表出现错误"},
			VideoList: nil,
			NextTime:  time.Now().Unix(),
		})
	}
	//将获取的video输出，方便测试
	//fmt.Printf("%#v", video)
	douyinVideos := make([]entity.DouyinVideo, len(video))
	for i := range video {
		author := entity.Author{
			Avatar:          video[i].UserDetail.Avatar,
			BackgroundImage: video[i].UserDetail.BackgroundImage,
			FavoriteCount:   video[i].UserDetail.FavoriteCount,
			FollowCount:     video[i].UserDetail.FollowCount,
			FollowerCount:   video[i].UserDetail.FollowerCount,
			ID:              video[i].UserDetail.ID,
			IsFollow:        video[i].UserDetail.IsFollow,
			Name:            video[i].UserDetail.Name,
			Signature:       video[i].UserDetail.Signature,
			TotalFavorited:  video[i].UserDetail.TotalFavorited,
			WorkCount:       video[i].UserDetail.WorkCount,
		}
		douyinVideo := &entity.DouyinVideo{
			Author:        author,
			CommentCount:  video[i].CommentCount,
			CoverURL:      video[i].CoverUrl,
			FavoriteCount: video[i].FavoriteCount,
			ID:            video[i].VideoId,
			IsFavorite:    video[i].IsFollow,
			PlayURL:       video[i].PlayUrl,
			Title:         video[i].Title,
		}
		douyinVideos[i] = *douyinVideo
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: entity.Response{
			StatusCode: 0,
			StatusMsg:  "查询视频列表成功",
		},
		VideoList: douyinVideos,
	})
}
