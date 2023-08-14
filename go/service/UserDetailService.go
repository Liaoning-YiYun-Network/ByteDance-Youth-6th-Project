package service

import (
	"SkyLine/dao"
	"SkyLine/entity"
)

// CreateUserDetail 创建用户详情
func CreateUserDetail(userDetail *entity.UserDetail) error {
	return dao.SqlSession.Create(userDetail).Error
}

// GetUserDetailById 根据id获取用户详情
func GetUserDetailById(id int) (*entity.UserDetail, error) {
	userDetail := new(entity.UserDetail)
	err := dao.SqlSession.Where("userid = ?", id).First(userDetail).Error
	return userDetail, err
}

// GetUserDetailByName 根据用户名获取用户详情
func GetUserDetailByName(name string) (*entity.UserDetail, error) {
	userDetail := new(entity.UserDetail)
	err := dao.SqlSession.Where("name = ?", name).First(userDetail).Error
	return userDetail, err
}

// UpdateUserDetail 更新用户详情
func UpdateUserDetail(userDetail *entity.UserDetail) error {
	return dao.SqlSession.Save(userDetail).Error
}

// DeleteUserDetail 删除用户详情
func DeleteUserDetail(userDetail *entity.UserDetail) error {
	return dao.SqlSession.Delete(userDetail).Error
}

// GetUserDetailList 获取用户详情列表
func GetUserDetailList() ([]*entity.UserDetail, error) {
	var userDetail []*entity.UserDetail
	err := dao.SqlSession.Find(&userDetail).Error
	return userDetail, err
}
