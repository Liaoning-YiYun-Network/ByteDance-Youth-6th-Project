package controller

import (
	"SkyLine/data"
	"SkyLine/entity"
	"SkyLine/perm"
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
	token := c.Query("token")
	feedRequest := &entity.FeedRequest{}
	video, err := service.SelectVideo(feedRequest)
	var isValid bool
	var user *entity.SQLUser
	if token != "" {
		isValid, _, user = perm.ValidateToken(token)

	}
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
		var tag bool
		ls, err := service.GetUserDetailById(int(user.UserId))
		if err != nil {
			data.Logger.Errorf("Try to get user detail failed: %v", err)
			tag = false
		} else {
			tag = true
		}
		var isFav = false
		if tag {
			lss, err := service.GetAllFavoritesByDBName(ls.FavoriteDB)
			if err != nil {
				data.Logger.Errorf("Try to get user detail failed: %v", err)
			} else {
				for _, value := range lss {
					if value == video[i].VideoId {
						isFav = true
					}
				}
			}
		}

		douyinVideo := &entity.DouyinVideo{
			Author:        author,
			CommentCount:  video[i].CommentCount,
			CoverURL:      video[i].CoverUrl,
			FavoriteCount: video[i].FavoriteCount,
			ID:            video[i].VideoId,
			IsFavorite:    isFav,
			PlayURL:       video[i].PlayUrl,
			Title:         video[i].Title,
		}
		douyinVideos[i] = *douyinVideo
	}
	if isValid {
		douyinVideos = IsFollow(douyinVideos, user.UserId)
	}
	//待根据业务逻辑，将查询到的东西返回前端
	c.JSON(http.StatusOK, entity.FeedResponse{
		Response: entity.Response{StatusCode: 0, StatusMsg: "获取视频成功"},
		//真实数据
		VideoList: douyinVideos,
		NextTime:  time.Now().Unix(),
	})
}

func IsFollow(douyinVideos []entity.DouyinVideo, id int64) []entity.DouyinVideo {
	userDetail, err := service.GetUserDetailById(int(id))
	if err != nil {
		fmt.Println(err)
	}
	isFollow := make(map[int64]bool)
	for i := 0; i < len(douyinVideos); i++ {
		if _, ok := isFollow[douyinVideos[i].Author.ID]; !ok {
			isFollow[douyinVideos[i].Author.ID] = false
		}
	}
	var ids []int64
	for key := range isFollow {
		ids = append(ids, key)
	}
	userIds, err := service.GetFollowByUserIds(userDetail.FollowDB, ids)
	if err != nil {
		return nil
	}
	for i := 0; i < len(douyinVideos); i++ {
		if contains(userIds, douyinVideos[i].Author.ID) {
			douyinVideos[i].Author.IsFollow = true
		} else {
			douyinVideos[i].Author.IsFollow = false
		}
	}

	return douyinVideos
}

func contains(slice []int, target int64) bool {
	id := int(target)
	for _, element := range slice {
		if element == id {
			return true
		}
	}
	return false
}
