package type_conv

import (
	"SkyLine/entity"
	"SkyLine/service"
)

func ToUser(su entity.SQLUser) entity.User {
	ud, err := service.GetUserDetailById(int(su.UserId))
	if err != nil {
		return entity.User{}
	}
	return entity.User{
		Id:            ud.ID,
		Name:          ud.Name,
		FollowerCount: ud.FollowCount,
		FollowCount:   ud.FollowerCount,
	}
}
