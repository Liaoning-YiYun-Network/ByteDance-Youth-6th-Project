package controller

import (
	"SkyLine/dao"
	"SkyLine/entity"
	"SkyLine/perm"
	"SkyLine/service"
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

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	username, err := dao.GetRedis(token)
	user, err := service.GetSQLUserByName(username)
	isValid, _, _ := perm.ValidateToken(token)
	actionType := c.Query("action_type")
	id := c.Query("comment_id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return
	}
	// 检查用户登录状态
	if isValid {
		if actionType == "1" {
			// 获取评论内容
			commentText := c.Query("comment_text")

			// 创建评论对象
			newComment := entity.Comment{
				Id:         idInt,
				User:       user.ToUser(), // 使用 ToUser 方法进行转换
				Content:    commentText,
				CreateDate: time.Now().Format("05-15"),
			}
			// 在数据库中添加评论
			err := service.AddCommentByDBName("your_database_name", entity.DBComment{
				CommentID: newComment.Id,
				UserID:    newComment.User.Id,
				Content:   newComment.Content,
				Time:      newComment.CreateDate,
			})
			if err != nil {
				c.JSON(http.StatusInternalServerError, entity.Response{
					StatusCode: 1,
					StatusMsg:  "Failed to add comment to the database",
				})
				return
			}
			// 返回新评论
			c.JSON(http.StatusOK, CommentActionResponse{
				Response: entity.Response{StatusCode: 0},
				Comment:  newComment,
			})
			return
		}
		if actionType == "2" {
			commentIDStr := c.Query("comment_id")
			commentID, err := strconv.ParseInt(commentIDStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, entity.Response{StatusCode: 2, StatusMsg: "Invalid comment ID"})
				return
			}

			// 调用删除评论的函数
			err = service.DeleteCommentByDBName("your_database_name", commentID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, entity.Response{StatusCode: 3, StatusMsg: "Failed to delete comment"})
				return
			}

			c.JSON(http.StatusOK, entity.Response{StatusCode: 0})
			return
		}
		c.JSON(http.StatusOK, entity.Response{StatusCode: 0})
	} else {
		// 未登录用户
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}

}

// CommentList all videos have same demo comment list
// CommentList 获取视频的所有评论，并按发布时间倒序排列
func CommentList(c *gin.Context) {
	videoID := c.Query("video_id") // 从请求中获取视频标识符

	// 使用 videoID 获取与该视频相关的评论数据，按发布时间倒序排列
	dbComments, err := service.GetAllCommentsByDBName(videoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 将 DBComment 转换为 Comment 类型
	var comments []entity.Comment
	for _, dbComment := range dbComments {
		comment := entity.Comment{
			Id:         dbComment.CommentID,
			User:       entity.User{Id: dbComment.UserID}, // 假设 User 有一个 Id 字段
			Content:    dbComment.Content,
			CreateDate: dbComment.Time,
		}
		comments = append(comments, comment)
	}

	// 构造响应结构体
	response := CommentListResponse{
		Response:    entity.Response{StatusCode: 0},
		CommentList: comments,
	}

	// 返回响应
	c.JSON(http.StatusOK, response)
}
