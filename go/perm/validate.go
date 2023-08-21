package perm

import (
	"SkyLine/dao"
	"SkyLine/entity"
	"SkyLine/service"
)

// ValidateToken 验证token是否有效
//
// token: token
//
// return: bool 是否有效, string false: 错误信息, entity.SQLUser true: 用户信息
func ValidateToken(token string) (bool, string, *entity.SQLUser) {
	username, err := dao.GetRedis(token)
	if err != nil {
		return false, "Token Invalid  or expired", new(entity.SQLUser)
	}
	user, err := service.GetSQLUserByName(username)
	if err != nil {
		return false, "User doesn't exist", new(entity.SQLUser)
	}
	return true, "", user
}
