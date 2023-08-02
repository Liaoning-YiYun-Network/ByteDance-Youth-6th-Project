package data

import "SkyLine/entity"

var Videos = []entity.DouyinVideo{
	{
		ID:            1,
		Author:        User,
		PlayURL:       "https://www.w3schools.com/html/movie.mp4",
		CoverURL:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         "我的老天爷",
	},
}

var User = entity.Author{
	Avatar:          "",         // 用户头像
	BackgroundImage: "",         // 用户个人页顶部大图
	FavoriteCount:   50,         // 喜欢数
	FollowCount:     0,          // 关注总数
	FollowerCount:   0,          // 粉丝总数
	ID:              1,          // 用户id
	IsFollow:        false,      // true-已关注，false-未关注
	Name:            "TestUser", // 用户名称
	Signature:       "wang",     // 个人简介
	TotalFavorited:  "10",       // 获赞数量
	WorkCount:       1,          // 作品数
}
