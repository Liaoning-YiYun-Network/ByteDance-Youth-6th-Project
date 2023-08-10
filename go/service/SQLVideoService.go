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
