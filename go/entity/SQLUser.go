package entity

// SQLUser 数据库中的用户表
type SQLUser struct {
	UserId   int64  `gorm:"column:userid;primary_key"` //用户唯一标识
	UserName string `gorm:"column:username"`           //用户名
	Password string `gorm:"column:password"`           //密码
	State    int8   `gorm:"column:state"`              //用户状态
}

func (SQLUser) TableName() string {
	return "user"
}

func (su *SQLUser) ToUser() User {
	return User{
		// 将 SQLUser 的字段赋值给 User 的字段
	}
}
