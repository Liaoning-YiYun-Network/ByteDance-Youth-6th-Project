package service

import (
	"SkyLine/dao"
	"SkyLine/entity"
)

func CreateSQLVideo(video *entity.SQLVideo) error {
	return dao.SqlSession.Create(video).Error
}

func GetSQLVideoById(id int) (*entity.SQLVideo, error) {
	video := new(entity.SQLVideo)
	err := dao.SqlSession.Where("id = ?", id).First(video).Error
	return video, err
}

func GetSQLVideosByAuthorId(id int) ([]*entity.SQLVideo, error) {
	var videos []*entity.SQLVideo
	err := dao.SqlSession.Where("user_id = ?", id).Find(&videos).Error
	return videos, err
}

func DeleteSQLVideo(video *entity.SQLVideo) error {
	return dao.SqlSession.Delete(video).Error
}

// SelectVideo 视频查询
//
// # FeedRequest
//
// return: SQLVideo,错误
func SelectVideo(feedRequest *entity.FeedRequest) ([]entity.SQLVideo, error) {
	var video []entity.SQLVideo
	err := dao.SqlSession.
		Order("create_time DESC").
		Limit(30).
		Find(&video).Error
	for i := range video {
		err = dao.SqlSession.Where("userid = ?", video[i].AuthorId).Find(&video[i].UserDetail).Error
	}
	return video, err
}
