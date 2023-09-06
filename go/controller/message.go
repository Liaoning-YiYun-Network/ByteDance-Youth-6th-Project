package controller

import (
	"SkyLine/dao"
	"SkyLine/data"
	"SkyLine/entity"
	"SkyLine/perm"
	"SkyLine/service"
	"SkyLine/util"
	"SkyLine/util/type_conv"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type ChatResponse struct {
	entity.Response
	MessageList []entity.Message `json:"message_list"`
}

// MessageAction 向指定用户ID发送消息
func MessageAction(c *gin.Context) {
	//!!! 请注意，这里的token是指用户A的token，而不是用户B的token
	isValid, msg, user := perm.ValidateToken(c.Query("token"))

	if isValid {
		//!!! 由于本项目中API均为客户端调用，故不考虑为空的情况，但在REST API中，应该考虑为空的情况
		userIdB, _ := strconv.Atoi(c.Query("to_user_id"))
		content := c.Query("content")
		chatKey := genChatKey(user.UserId, int64(userIdB))

		dbName := fmt.Sprintf("messages-%v.sqlite", chatKey)
		if !util.IsFileExist("./dbs/messages/" + dbName) {
			data.Logger.Info("Database not exist: " + dbName + ", try to create it now")
			_, err := dao.CreateDB(dao.MESSAGES, chatKey)
			if err != nil {
				c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "Message sent failed due to database error"})
				return
			}
		}

		err := service.AddMessageByDBName(dbName, entity.DBMessage{
			UserID:     user.UserId,
			Content:    content,
			CreateTime: time.Now().Format(time.Kitchen),
		})
		if err != nil {
			c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: "Message sent failed due to database error"})
			return
		}

		c.JSON(http.StatusOK, entity.Response{StatusCode: 0, StatusMsg: "Message sent successfully"})
	} else {
		c.JSON(http.StatusOK, entity.Response{StatusCode: 1, StatusMsg: msg})
	}
}

// MessageChat 查询与指定用户ID的聊天记录
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusOK, ChatResponse{
			Response:    entity.Response{StatusCode: 1, StatusMsg: "Query chat history failed due to empty token"},
			MessageList: nil,
		})
		return
	}

	isValid, msg, user := perm.ValidateToken(token)

	if isValid {
		//!!! 由于本项目中API均为客户端调用，故不考虑为空的情况，但在REST API中，应该考虑为空的情况
		userIdB, _ := strconv.Atoi(c.Query("to_user_id"))
		chatKey := genChatKey(user.UserId, int64(userIdB))

		dbName := fmt.Sprintf("messages-%v.sqlite", chatKey)

		if !util.IsFileExist("./dbs/messages/" + dbName) {
			c.JSON(http.StatusOK, ChatResponse{
				Response:    entity.Response{StatusCode: 0, StatusMsg: "Query chat history successfully"},
				MessageList: nil,
			})
			return
		}

		messages, err := service.GetAllMessagesByDBName(dbName)
		if err != nil {
			c.JSON(http.StatusOK, ChatResponse{
				Response:    entity.Response{StatusCode: 1, StatusMsg: "Query chat history failed due to database error"},
				MessageList: nil,
			})
			return
		}

		list := type_conv.ToMessageList(messages)
		c.JSON(http.StatusOK, ChatResponse{
			Response:    entity.Response{StatusCode: 0, StatusMsg: "Query chat history successfully"},
			MessageList: list,
		})
	} else {
		c.JSON(http.StatusOK, ChatResponse{
			Response:    entity.Response{StatusCode: 1, StatusMsg: msg},
			MessageList: nil,
		})
	}
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
