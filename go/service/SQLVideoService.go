package service

import (
	"SkyLine/dao"
	"SkyLine/entity"
	"fmt"
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
// user: 用户实体
//
// return: 错误
func SelectVideo(feedRequest *entity.FeedRequest) ([]entity.SQLVideo, error) {
	var video []entity.SQLVideo
	query := `
        SELECT * 
        FROM video AS v LEFT JOIN userdetail AS u ON v.userid = u.userid
		order by create_time DESC
		limit 2
    `
	result := dao.SqlSession.Raw(query).Scan(&video)
	if result.Error != nil {
		return nil, fmt.Errorf("Failed to execute SQL query: %v", result.Error)
	}
	return video, result.Error
}
