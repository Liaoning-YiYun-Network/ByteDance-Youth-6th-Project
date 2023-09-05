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

// UpdateUserById 根据用户id详情列表
func UpdateUserById(userDetail *entity.UserDetail) error {
	//err := dao.SqlSession.Where("userid = ?", userDetail.ID).Updates(userDetail).Error
	err := dao.SqlSession.Updates(userDetail).Error
	return err
}

// 根据用户id改变关注总数
func ChangeFollowCount(id int, isAdd bool) error {
	var userDetail []*entity.UserDetail
	//userDetailById, err := GetUserDetailById(id)
	err := dao.SqlSession.Where("userid = ?", id).Find(&userDetail).Error
	if err != nil {
		return err
	}
	if isAdd {
		userDetail[0].FollowCount++
	} else {
		userDetail[0].FollowCount--
	}
	//err = UpdateUserDetail(userDetailById)
	dao.SqlSession.Model(&userDetail[0]).Where("userid = ?", id).Update("follow_count", userDetail[0].FollowCount)
	if err != nil {
		return err
	}
	return nil
}

// 根据用户id改变粉丝数量
func ChangeFollowerCount(id int, isAdd bool) error {
	//userDetailById, err := GetUserDetailById(id)
	var userDetail []*entity.UserDetail
	err := dao.SqlSession.Where("userid = ?", id).Find(&userDetail).Error
	if err != nil {
		return err
	}
	if isAdd {
		userDetail[0].FollowerCount++
	} else {
		userDetail[0].FollowerCount--
	}
	dao.SqlSession.Model(&userDetail[0]).Where("userid = ?", id).Update("follower_count", userDetail[0].FollowerCount)
	if err != nil {
		return err
	}
	return nil
}
