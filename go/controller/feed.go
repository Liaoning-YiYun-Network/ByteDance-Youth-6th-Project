package controller

import (
	"SkyLine/entity"
	"SkyLine/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	entity.Response
	VideoList []entity.DouyinVideo `json:"video_list,omitempty"`
	NextTime  int64                `json:"next_time,omitempty"`
}

// @Summary  获取视频流
// @Description  这个接口，在用户刚进入抖音之后就会被调用，并将视频以及作者的信息推送给用户
// @Tags         视频相关接口
// @Param        LatestTime  query  string  false  "可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间"
// @Param        Token  query  string  false  "该参数只有在用户登录状态下进行设置"
// @Router       /douyin/feed [get]
func Feed(c *gin.Context) {
	feedRequest := &entity.FeedRequest{nil, nil}
	video, err := service.SelectVideo(feedRequest)
	if err != nil {
		fmt.Printf("视频获取出错:%v\n", err)
		c.JSON(http.StatusInternalServerError, entity.FeedResponse{
			Response:  entity.Response{StatusCode: 1, StatusMsg: "获取视频出现错误"},
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
	//待根据业务逻辑，将查询到的东西返回前端
	c.JSON(http.StatusOK, entity.FeedResponse{
		Response: entity.Response{StatusCode: 0, StatusMsg: "获取视频成功"},
		//真实数据
		VideoList: douyinVideos,
		//为方便测试伪造的数据
		//VideoList: data.Videos,
		NextTime: time.Now().Unix(),
	})
}
