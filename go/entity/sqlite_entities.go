package entity

type DBComment struct {
	CommentID int64  `gorm:"primary_key;column:comment_id;type:INTEGER;not null"`
	UserID    int64  `gorm:"column:user_id;type:INTEGER;not null"`
	Content   string `gorm:"column:content;type:TEXT;not null"`
	Time      string `gorm:"column:time;type:TEXT;not null"`
}

func (DBComment) TableName() string {
	return "comments"
}

type DBMessage struct {
	MessageID  int64  `gorm:"primary_key;column:message_id;type:INTEGER;not null"`
	UserID     int64  `gorm:"column:user_id;type:INTEGER;not null"`
	Content    string `gorm:"column:content;type:TEXT;not null"`
	CreateTime string `gorm:"column:create_time;type:TEXT;not null"`
}

func (DBMessage) TableName() string {
	return "messages"
}
