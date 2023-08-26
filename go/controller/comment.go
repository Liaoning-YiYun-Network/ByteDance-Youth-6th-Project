package controller

import (
	"SkyLine/data"
	"SkyLine/entity"
	"SkyLine/perm"
	"SkyLine/service"
	"SkyLine/util/type_conv"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CommentListResponse struct {
	entity.Response
	CommentList []entity.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	entity.Response
	Comment entity.Comment `json:"comment,omitempty"`
}

var tmpIdSequence map[int]int64

var tmpIdInitialized map[int]bool

// CommentAction add or delete comment
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	vidInt, _ := strconv.Atoi(c.Query("video_id"))
	sv, err := service.GetSQLVideoById(vidInt)

	if err != nil {
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "Failed to get video"},
		})
		return
	}
	if tmpIdInitialized[vidInt] == false {
		cs, err := service.GetAllCommentsByDBName(sv.CommentDB)
		if err != nil {
			data.Logger.Error("Failed to init comments: ", err)
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: entity.Response{StatusCode: 1, StatusMsg: "Failed to get comments"},
			})
			return
		}
		//获取列表中最大的id
		maxId := int64(0)
		for _, c := range cs {
			if c.CommentID > maxId {
				maxId = c.CommentID
			}
		}
		tmpIdSequence[vidInt] = maxId + 1
	}
	isValid, msg, user := perm.ValidateToken(token)
	actionType := c.Query("action_type")
	// 检查用户登录状态
	if isValid {
		if actionType == "1" {
			// 获取评论内容
			commentText := c.Query("comment_text")
			// 获取评论用户
			user, err := type_conv.ToUser(*user)
			if err != nil {
				data.Logger.Error("Failed to convert user: ", err)
				c.JSON(http.StatusOK, CommentActionResponse{
					Response: entity.Response{StatusCode: 1, StatusMsg: "Failed to convert user"},
				})
				return
			}
			// 创建评论对象
			newComment := entity.Comment{
				Id:         tmpIdSequence[vidInt],
				User:       user,
				Content:    commentText,
				CreateDate: strconv.FormatInt(time.Now().Unix(), 10),
			}
			// 在数据库中添加评论
			err = service.AddCommentByDBName(sv.CommentDB, entity.DBComment{
				UserID:  newComment.User.Id,
				Content: newComment.Content,
				Time:    newComment.CreateDate,
			})
			if err != nil {
				c.JSON(http.StatusOK, CommentActionResponse{
					Response: entity.Response{StatusCode: 1, StatusMsg: "Failed to add comment"},
				})
				return
			}
			// 返回新评论
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: entity.Response{StatusCode: 0},
				Comment:  newComment,
			})
			tmpIdSequence[vidInt]++
			return
		} else {
			// ！！！由于此处的数值是从客户端传过来的，所以不考虑出现异常的情况。但是在Restful API中，应该考虑到异常情况。
			commentID, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)

			// 调用删除评论的函数
			err := service.DeleteCommentByDBName(sv.CommentDB, commentID)
			if err != nil {
				c.JSON(http.StatusOK, CommentActionResponse{
					Response: entity.Response{StatusCode: 1, StatusMsg: "Failed to delete comment"},
				})
				return
			}

			c.JSON(http.StatusOK, CommentActionResponse{
				Response: entity.Response{StatusCode: 0, StatusMsg: "Delete comment success"},
			})
			return
		}
	} else {
		// 未登录用户
		c.JSON(http.StatusOK, CommentActionResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: msg},
		})
	}

}

// CommentList 获取视频的所有评论，并按发布时间倒序排列
func CommentList(c *gin.Context) {
	// ！！！由于此处的数值是从客户端传过来的，所以不考虑出现异常的情况。但是在Restful API中，应该考虑到异常情况。
	videoID, _ := strconv.Atoi(c.Query("video_id")) // 从请求中获取视频标识符
	sv, err := service.GetSQLVideoById(videoID)
	// 使用 videoID 获取与该视频相关的评论数据，按发布时间倒序排列
	dbComments, err := service.GetAllCommentsByDBName(sv.CommentDB)
	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "Failed to get comments"},
		})
		return
	}

	// 将 DBComment 转换为 Comment 类型
	comments := make([]entity.Comment, 0)
	for _, dbComment := range dbComments {
		ud, err := service.GetUserDetailById(int(dbComment.UserID))
		if err != nil {
			data.Logger.Error("Failed to get user detail by id: ", dbComment.UserID, ", Skip!")
			continue
		} else {
			comments = append(comments, entity.Comment{
				Id: dbComment.CommentID,
				User: entity.User{
					Id:            dbComment.UserID,
					Name:          ud.Name,
					FollowCount:   ud.FollowCount,
					FollowerCount: ud.FollowerCount,
				},
				Content:    dbComment.Content,
				CreateDate: dbComment.Time,
			})
		}
	}

	// 返回响应
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    entity.Response{StatusCode: 0},
		CommentList: comments,
	})
}
