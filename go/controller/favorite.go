package controller

import (
	"SkyLine/data"
	"SkyLine/entity"
	"SkyLine/perm"
	"SkyLine/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	isValid, msg, user := perm.ValidateToken(c.Query("token"))

	if isValid {
		vid := c.Query("video_id")
		vidInt, _ := strconv.Atoi(vid)
		actionType := c.Query("action_type")
		ud, err := service.GetUserDetailById(int(user.UserId))
		if err != nil {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "Failed to get user detail"})
			return
		}
		sv, err := service.GetSQLVideoById(vidInt)
		if err != nil {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "Failed to get video detail"})
			return
		}
		if actionType == "1" {
			sv.FavoriteCount++
			err = service.UpdateSQLVideo(sv)
			if err != nil {
				c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "Failed to update video"})
				return
			}
			err = service.AddFavoriteByDBName(ud.FavoriteDB, int64(vidInt))
			if err != nil {
				data.Logger.Errorf("???:%v", err)
				c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "Failed to add favorite"})
				sv.FavoriteCount--
				//TODO：此处的错误需要处理！！！
				_ = service.UpdateSQLVideo(sv)
				return
			}

			c.JSON(http.StatusOK, entity.Response{StatusCode: 0, StatusMsg: "Add favorite success"})
		} else {
			sv.FavoriteCount--
			err = service.UpdateSQLVideo(sv)
			if err != nil {
				c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "Failed to update video"})
				return
			}
			err := service.DeleteFavoriteByDBName(ud.FavoriteDB, int64(vidInt))
			if err != nil {
				c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "Failed to delete favorite"})
				sv.FavoriteCount++
				//TODO：此处的错误需要处理！！！
				_ = service.UpdateSQLVideo(sv)
				return
			}
			c.JSON(http.StatusOK, entity.Response{StatusCode: 0, StatusMsg: "Delete favorite success"})
		}
	} else {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: msg})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	isValid, msg, user := perm.ValidateToken(c.Query("token"))

	if isValid {
		ud, err := service.GetUserDetailById(int(user.UserId))
		if err != nil {
			c.JSON(http.StatusOK, VideoListResponse{
				Response: entity.Response{
					StatusCode: 1,
					StatusMsg:  "Failed to get user detail",
				},
				VideoList: nil,
			})
			return
		}
		favorites, err := service.GetAllFavoritesByDBName(ud.FavoriteDB)
		if err != nil {
			c.JSON(http.StatusOK, VideoListResponse{
				Response: entity.Response{
					StatusCode: 1,
					StatusMsg:  "Failed to get favorites",
				},
				VideoList: nil,
			})
			return
		}
		var videos []entity.DouyinVideo
		for _, favorite := range favorites {
			v, err := service.GetSQLVideoById(int(favorite))
			if err != nil {
				data.Logger.Errorf("Failed to get video detail by id: %d, Skip! \nError:%v", favorite, err)
				continue
			}
			var ar entity.Author
			ud, err := service.GetUserDetailById(int(v.AuthorId))
			if err != nil {
				data.Logger.Errorf("Failed to get user detail by id: %d, Skip! \nError:%v", v.AuthorId, err)
				continue
			}
			ar.ID = ud.ID
			ar.Name = ud.Name
			ar.FollowCount = ud.FollowCount
			ar.FollowerCount = ud.FollowerCount
			ar.Avatar = ud.Avatar
			ar.BackgroundImage = ud.BackgroundImage
			ar.Signature = ud.Signature
			list, err := service.GetAllFollowersByDBName(ud.FollowerDB)
			if err != nil {
				data.Logger.Errorf("Failed to get followers by db name: %s, Skip! \nError:%v", ud.FollowerDB, err)
				continue
			}
			for _, follower := range list {
				if follower == user.UserId {
					ar.IsFollow = true
					break
				}
			}
			ar.WorkCount = ud.WorkCount
			//获取所有作品的总点赞数
			var totalLikeCount int64
			lt, err := service.GetSQLVideosByAuthorId(int(v.AuthorId))
			if err != nil {
				data.Logger.Errorf("Failed to get videos by author id: %d, Skip! \nError:%v", v.AuthorId, err)
				continue
			}
			for _, v := range lt {
				totalLikeCount += v.FavoriteCount
			}
			ar.TotalFavorited = strconv.FormatInt(totalLikeCount, 10)
			var video entity.DouyinVideo
			video.ID = v.VideoId
			video.Author = ar
			video.PlayURL = v.PlayUrl
			video.CoverURL = v.CoverUrl
			video.FavoriteCount = v.FavoriteCount
			video.CommentCount = v.CommentCount
			video.IsFavorite = true
			videos = append(videos, video)
		}
		c.JSON(http.StatusOK, VideoListResponse{
			Response: entity.Response{
				StatusCode: 0,
				StatusMsg:  "Get favorite list success",
			},
			VideoList: videos,
		})
	} else {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: entity.Response{
				StatusCode: 1,
				StatusMsg:  msg,
			},
			VideoList: nil,
		})
	}
}
