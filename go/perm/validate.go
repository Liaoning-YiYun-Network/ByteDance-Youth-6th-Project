package perm

import (
	"SkyLine/dao"
	"SkyLine/service"
	"strconv"
)

// ValidateToken 验证token是否有效
//
// token: token
//
// return: bool 是否有效, string true: 用户id/false: 错误信息
func ValidateToken(token string) (bool, string) {
	username, err := dao.GetRedis(token)
	if err != nil {
		return false, "Token Invalid  or expired"
	}
	user, err := service.GetSQLUserByName(username)
	if err != nil {
		return false, "User doesn't exist"
	}
	return true, strconv.FormatInt(user.UserId, 10)
}
