package controller

import (
	"SkyLine/dao"
	"SkyLine/data"
	"SkyLine/entity"
	"SkyLine/service"
	"SkyLine/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"sync/atomic"
)

var usersLoginInfo map[string]entity.User

var userIdSequence = int64(-1)

type UserLoginResponse struct {
	entity.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	entity.Response
	User entity.User `json:"user"`
}

func Register(c *gin.Context) {
	if userIdSequence == -1 {
		users, err := service.GetSQLUserList()
		if err != nil {
			fmt.Println("获取用户列表失败，运行终止！")
			panic(err)
		}
		fmt.Println("获取用户列表成功")
		usersLoginInfo = make(map[string]entity.User)
		userIdSequence = int64(len(users) + 1)
	}
	username := c.Query("username")
	password := c.Query("password")

	//从数据库中查询用户是否存在
	user, _ := service.GetSQLUserByName(username)
	if user.UserName != "" {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		str := util.EncryptWithMD5(password)
		sqlUser := entity.SQLUser{
			UserId:   userIdSequence,
			UserName: username,
			Password: str,
			State:    1,
		}
		// 创建数据库文件
		followDBName, err := dao.CreateDB(dao.FOLLOWS, strconv.FormatInt(userIdSequence, 10))
		if err != nil {
			data.Logger.Errorf("注册过程中创建关注者数据库失败：%s\n", err)
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: entity.Response{StatusCode: 1, StatusMsg: "Register failed due to followDB create failed"},
			})
			return
		}
		followerDBName, err := dao.CreateDB(dao.FOLLOWERS, strconv.FormatInt(userIdSequence, 10))
		if err != nil {
			_ = dao.DeleteDB(dao.FOLLOWS, followDBName)
			data.Logger.Errorf("注册过程中创建粉丝数据库失败：%s\n", err)
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: entity.Response{StatusCode: 1, StatusMsg: "Register failed due to followerDB create failed"},
			})
			return
		}
		favoriteDBName, err := dao.CreateDB(dao.FAVORITES, strconv.FormatInt(userIdSequence, 10))
		if err != nil {
			_ = dao.DeleteDB(dao.FOLLOWS, followDBName)
			_ = dao.DeleteDB(dao.FOLLOWERS, followerDBName)
			data.Logger.Errorf("注册过程中创建收藏夹数据库失败：%s\n", err)
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: entity.Response{StatusCode: 1, StatusMsg: "Register failed due to favoriteDB create failed"},
			})
			return
		}
		token, err := util.GenerateToken(sqlUser)
		if err != nil {
			token = username + password
		}
		err = dao.SetRedis(token, username)
		if err != nil {
			data.Logger.Warnf("Set Redis KV Failed: %s", err)
		}
		userDetail := entity.UserDetail{
			ID:              userIdSequence,
			Name:            username,
			Avatar:          data.DefaultAvatar,
			BackgroundImage: data.DefaultBackgroundImage,
			Signature:       data.DefaultSignature,
			FollowDB:        followDBName,
			FollowerDB:      followerDBName,
			FavoriteDB:      favoriteDBName,
		}
		err = service.CreateSQLUser(&sqlUser)
		if err != nil {
			_ = dao.DeleteDB(dao.FOLLOWS, followDBName)
			_ = dao.DeleteDB(dao.FOLLOWERS, followerDBName)
			_ = dao.DeleteDB(dao.FAVORITES, favoriteDBName)
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: entity.Response{StatusCode: 1, StatusMsg: "Register failed due to database error"},
			})
			return
		}
		err = service.CreateUserDetail(&userDetail)
		if err != nil {
			_ = dao.DeleteDB(dao.FOLLOWS, followDBName)
			_ = dao.DeleteDB(dao.FOLLOWERS, followerDBName)
			_ = dao.DeleteDB(dao.FAVORITES, favoriteDBName)
			// TODO：该处错误需要处理！！！！！！
			service.DeleteSQLUserById(int(userIdSequence))
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: entity.Response{StatusCode: 1, StatusMsg: "Register failed due to database error"},
			})
			return
		}
		atomic.AddInt64(&userIdSequence, 1)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 0, StatusMsg: "Register Success"},
			UserId:   sqlUser.UserId,
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	if userIdSequence == -1 {
		users, err := service.GetSQLUserList()
		if err != nil {
			fmt.Println("获取用户列表失败，运行终止！")
			panic(err)
		}
		fmt.Println("获取用户列表成功")
		usersLoginInfo = make(map[string]entity.User)
		userIdSequence = int64(len(users) + 1)
	}
	username := c.Query("username")
	password := c.Query("password")
	encryptedKey := util.EncryptWithMD5(password)
	user, err := service.GetSQLUserByName(username)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "Login Failed"},
		})
		return
	}
	if user.Password != encryptedKey {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "Wrong Password!"},
			UserId:   user.UserId,
		})
		return
	}
	token, err := util.GenerateToken(*user)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "Login Failed"},
			UserId:   user.UserId,
		})
		return
	}
	err = dao.SetRedis(token, username)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: entity.Response{StatusCode: 1, StatusMsg: "Login Failed"},
			UserId:   user.UserId,
			Token:    token,
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: entity.Response{StatusCode: 0, StatusMsg: "Login Success"},
		UserId:   user.UserId,
		Token:    token,
	})
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	username, err := dao.GetRedis(token)
	if err != nil {
		id, _ := strconv.Atoi(c.Query("user_id"))
		detail, err := service.GetUserDetailById(id)
		if err != nil || detail.Name == "" {
			c.JSON(http.StatusOK, UserResponse{
				Response: entity.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			})
			return
		} else {
			c.JSON(http.StatusOK, UserResponse{
				Response: entity.Response{StatusCode: 0, StatusMsg: "Nothing"},
				User: entity.User{
					Id:            detail.ID,
					Name:          detail.Name,
					FollowCount:   detail.FollowCount,
					FollowerCount: detail.FollowerCount,
				},
			})
			return
		}
	} else {
		detail, err := service.GetUserDetailByName(username)
		if err != nil {
			c.JSON(http.StatusOK, UserResponse{
				Response: entity.Response{StatusCode: 1, StatusMsg: "Get UserInfo Failed"},
			})
			return
		}
		c.JSON(http.StatusOK, UserResponse{
			Response: entity.Response{StatusCode: 0, StatusMsg: "Nothing"},
			User: entity.User{
				Id:            detail.ID,
				Name:          detail.Name,
				FollowCount:   detail.FollowCount,
				FollowerCount: detail.FollowerCount,
			},
		})
	}
}
