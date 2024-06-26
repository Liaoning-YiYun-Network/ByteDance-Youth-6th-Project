package controller

import (
	"SkyLine/dao"
	"SkyLine/entity"
	"SkyLine/perm"
	"SkyLine/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const IS_FOLLOW_TURN = "2"
const IS_FOLLOW_FALSE = "1"

type UserListResponse struct {
	entity.Response
	UserList []entity.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	isUser, err, user := perm.ValidateToken(token)
	if isUser == false {
		c.JSON(http.StatusUnauthorized, entity.Response{StatusCode: 1, StatusMsg: err})
	}
	toUserId := c.Query("to_user_id")
	hisId, _ := strconv.Atoi(toUserId)
	actionType := c.Query("action_type")
	if actionType == IS_FOLLOW_FALSE {
		myFollow, _, err := service.GetFollowAndFollowerByUserid(user.UserId)
		_, hisFollower, err := service.GetFollowAndFollowerByUserid(int64(hisId))
		if err != nil {
			fmt.Println("获取失败")
		} else {
			tx := dao.SqlSession.Begin()
			if err := service.ChangeFollowCount(int(user.UserId), true); err != nil {
				// 发生错误，回滚事务
				tx.Rollback()
			}
			if err := service.ChangeFollowerCount(hisId, true); err != nil {
				// 发生错误，回滚事务
				tx.Rollback()
			}
			if err = tx.Commit().Error; err != nil {
				c.JSON(http.StatusInternalServerError, entity.Response{StatusCode: 1, StatusMsg: "进行关注失败"})
			}
			service.AddFollowByDBName(myFollow, int64(hisId))
			service.AddFollowerByDBName(hisFollower, user.UserId)
		}
		c.JSON(http.StatusOK, entity.Response{StatusCode: 0})
	}
	if actionType == IS_FOLLOW_TURN {
		myFollow, _, err := service.GetFollowAndFollowerByUserid(user.UserId)
		_, hisFollower, err := service.GetFollowAndFollowerByUserid(int64(hisId))
		if err != nil {
			fmt.Println("获取失败")
		} else {
			tx := dao.SqlSession.Begin()
			if err := service.ChangeFollowCount(int(user.UserId), false); err != nil {
				// 发生错误，回滚事务
				tx.Rollback()
			}
			if err := service.ChangeFollowerCount(hisId, false); err != nil {
				// 发生错误，回滚事务
				tx.Rollback()
			}
			if err = tx.Commit().Error; err != nil {
				c.JSON(http.StatusInternalServerError, entity.Response{StatusCode: 1, StatusMsg: "进行关注失败"})
			}
			service.DeleteFollowByDBName(myFollow, int64(hisId))
			service.DeleteFollowerByDBName(hisFollower, user.UserId)
		}
		c.JSON(http.StatusOK, entity.Response{StatusCode: 0})
	}
	c.JSON(http.StatusBadRequest, entity.Response{StatusCode: 1, StatusMsg: "参数错误"})
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: entity.Response{
			StatusCode: 0,
		},
		UserList: []entity.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: entity.Response{
			StatusCode: 0,
		},
		UserList: []entity.User{DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: entity.Response{
			StatusCode: 0,
		},
		UserList: []entity.User{DemoUser},
	})
}
