package type_conv

import (
	"SkyLine/entity"
	"SkyLine/service"
)

func ToUser(su entity.SQLUser) (entity.User, error) {
	ud, err := service.GetUserDetailById(int(su.UserId))
	if err != nil {
		return entity.User{}, err
	}
	return entity.User{
		Id:            ud.ID,
		Name:          ud.Name,
		FollowerCount: ud.FollowCount,
		FollowCount:   ud.FollowerCount,
	}, nil
}
