package entity

type SQLUser struct {
	UserId     int    `gorm:"column:userid;primary_key"`
	UserName   string `gorm:"column:username"`
	Password   string `gorm:"column:passwd"`
	Avatar     string `gorm:"column:avatar"`
	Background string `gorm:"column:background"`
	Signature  string `gorm:"column:signature"`
}

func (SQLUser) TableName() string {
	return "user"
}
