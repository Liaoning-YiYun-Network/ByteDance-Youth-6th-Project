package service

import (
	"SkyLine/dao"
	"SkyLine/entity"
)

// CreateSQLUser 创建用户
//
// user: 用户实体
//
// return: 错误
func CreateSQLUser(user *entity.SQLUser) error {
	return dao.SqlSession.Create(user).Error
}

// GetSQLUserById 根据id获取用户
//
// id: 用户id
//
// return: 用户实体/错误
func GetSQLUserById(id int) (*entity.SQLUser, error) {
	user := new(entity.SQLUser)
	err := dao.SqlSession.Where("`userid` = ?", id).First(user).Error
	return user, err
}

// GetSQLUserByName 根据用户名获取用户
//
// name: 用户名
//
// return: 用户实体/错误
func GetSQLUserByName(name string) (*entity.SQLUser, error) {
	user := new(entity.SQLUser)
	err := dao.SqlSession.Where("`username` = ?", name).First(user).Error
	return user, err
}

// UpdateSQLUser 更新用户
//
// user: 用户实体
//
// return: 错误
func UpdateSQLUser(user *entity.SQLUser) error {
	return dao.SqlSession.Save(user).Error
}

// DeleteSQLUser 删除用户
//
// user: 用户实体
//
// return: 错误
func DeleteSQLUser(user *entity.SQLUser) error {
	return dao.SqlSession.Delete(user).Error
}

// GetSQLUserList 获取用户列表
//
// return: 用户列表/错误
func GetSQLUserList() ([]*entity.SQLUser, error) {
	var users []*entity.SQLUser
	err := dao.SqlSession.Find(&users).Error
	return users, err
}

//GetFollowAndFollowerByUserid 根据用户id获取关注和粉丝所存储的数据库名

//userid 用户id

//return 关注和粉丝所存储的数据库名

func GetFollowAndFollowerByUserid(Userid int64) (follow, follower string, err error) {
	var users entity.UserDetail
	err = dao.SqlSession.Where("userid = ?", Userid).First(&users).Error
	return users.FollowDB, users.FollowerDB, err
}
