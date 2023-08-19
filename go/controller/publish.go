package controller

import (
	"SkyLine/dao"
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
	username, err := dao.GetRedis(token)
	if err != nil {
		c.JSON(http.StatusForbidden, entity.Response{
			StatusCode: 1,
			StatusMsg:  "token invalid or expired",
		})
		return
	}
	user, err := service.GetSQLUserByName(username)
	if err != nil {
		c.JSON(http.StatusForbidden, entity.Response{
			StatusCode: 1,
			StatusMsg:  "user not exist",
		})
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{
			StatusCode: 1,
			StatusMsg:  "Failed to get file",
		})
		return
	}
	// 获取文件大小
	fileSize := data.Size
	// 设置文件大小限制，例如 500MB
	maxFileSize := int64(500 * 1024 * 1024) // 500MB
	if fileSize > maxFileSize {
		c.JSON(http.StatusRequestEntityTooLarge, entity.Response{
			StatusCode: 1,
			StatusMsg:  "File size exceeds the limit",
		})
		return
	}
	file, err := data.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{
			StatusCode: 1,
			StatusMsg:  "Failed to open file",
		})
		return
	}
	defer file.Close()
	fileContent, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{
			StatusCode: 1,
			StatusMsg:  "Failed to read file",
		})
		return
	}
	fileName := data.Filename
	videoUUID, err := util.UUIDWithoutHyphen()
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{
			StatusCode: 1,
			StatusMsg:  "Failed to process file name",
		})
		return
	}
	//截取后面.MP4，防止最后拼接成.MP4.MP4
	fileName = fileName[0 : len(fileName)-4]

	//生成视频、视频封面地址
	newVideoName := fmt.Sprintf("%d-%s-%s.mp4", user.UserId, videoUUID, fileName)
	videoUrl := "https://tos.eyunnet.com/videos/" + newVideoName
	coverName := fmt.Sprintf("%d-%s-%s.jpg", user.UserId, videoUUID, fileName)
	coverUrl := "https://tos.eyunnet.com/video_covers/" + coverName

	//在数据库中加入视频信息
	err = service.CreateSQLVideo(&entity.SQLVideo{
		AuthorId: user.UserId,
		Title:    c.PostForm("title"),
		PlayUrl:  videoUrl,
		CoverUrl: coverUrl,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, entity.Response{
			StatusCode: 1,
			StatusMsg:  "Failed to save video info",
		})
		return
	}

	// 调用上传函数上传视频
	//函数的第二个参数需要是 []byte 类型
	err = service.UploadFile(newVideoName, fileContent, service.VIDEO)
	if err != nil {
		service.DeleteSQLVideo(&entity.SQLVideo{
			AuthorId: user.UserId,
			Title:    c.PostForm("title"),
			PlayUrl:  videoUrl,
			CoverUrl: coverUrl,
		})
		// 处理错误
		c.JSON(http.StatusInternalServerError, entity.Response{
			StatusCode: 1,
			StatusMsg:  "Failed to upload video file",
		})
		return
	}

	//截取视频封面
	coverBytes, err := ReadFrameAsJpeg(videoUrl)
	if err != nil {
		fmt.Printf("视频封面截取失败%v\n", err)
		c.JSON(http.StatusInternalServerError, entity.Response{
			StatusCode: 1,
			StatusMsg:  "截取视频封面出现错误",
		})
		return
	}

	// 调用上传函数上传视频封面
	err = service.UploadFile(coverName, coverBytes, service.VIDEO_COVER)
	if err != nil {
		fmt.Println("上传文件时发生错误：", err)
		service.DeleteFile(newVideoName, service.VIDEO)
		service.DeleteSQLVideo(&entity.SQLVideo{
			AuthorId: user.UserId,
			Title:    c.PostForm("title"),
			PlayUrl:  videoUrl,
			CoverUrl: coverUrl,
		})
		c.JSON(http.StatusInternalServerError, entity.Response{
			StatusCode: 1,
			StatusMsg:  coverName + " failed in uploading",
		})
		return
	} else {
		fmt.Printf("文件%s上传成功！\n", newVideoName)
	}

	c.JSON(http.StatusOK, entity.Response{
		StatusCode: 0,
		StatusMsg:  "Publish success",
	})
}

// ReadFrameAsJpeg 截取视频封面
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
		return
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
