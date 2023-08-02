package controller

import (
	"SkyLine/data"
	"SkyLine/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	entity.Response
	VideoList []entity.DouyinVideo `json:"video_list,omitempty"`
	NextTime  int64                `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  entity.Response{StatusCode: 0, StatusMsg: "无事发生"},
		VideoList: data.Videos,
		NextTime:  time.Now().Unix(),
	})
}
