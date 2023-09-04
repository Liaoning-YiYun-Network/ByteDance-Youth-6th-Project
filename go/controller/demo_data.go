package controller

import (
	"SkyLine/entity"
)

var DemoUser = entity.User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
