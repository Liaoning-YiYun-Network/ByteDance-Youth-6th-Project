package entity

type UserDetail struct {
	ID              int64  `gorm:"column:userid"`
	Avatar          string `gorm:"column:avatar"`           // 用户头像
	BackgroundImage string `gorm:"column:background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `gorm:"column:favorite_count"`   // 喜欢数
	FollowCount     int64  `gorm:"column:follow_count"`     // 关注总数
	FollowerCount   int64  `gorm:"column:follower_count"`   // 粉丝总数
	Name            string `gorm:"column:name"`             // 用户名称
	Signature       string `gorm:"column:signature"`        // 个人简介
	TotalFavorited  string `gorm:"column:total_favorited"`  // 获赞数量
	WorkCount       int64  `gorm:"column:work_count"`       // 作品数
}

// 定义Author对应的数据库表名
func (UserDetail) TableName() string {
	return "userdetail"
}
