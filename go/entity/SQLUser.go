package entity

// SQLUser 数据库中的用户表
type SQLUser struct {
	UserId        int64  `gorm:"column:userid;primary_key"`
	UserName      string `gorm:"column:username"`
	Password      string `gorm:"column:passwd"`
	Avatar        string `gorm:"column:avatar"`
	Background    string `gorm:"column:background"`
	Signature     string `gorm:"column:signature"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
}

func (SQLUser) TableName() string {
	return "user"
}
