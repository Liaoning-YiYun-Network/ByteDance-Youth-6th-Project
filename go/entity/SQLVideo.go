package entity

type SQLVideo struct {
	VideoId       int64  `gorm:"column:id;primary_key"`
	AuthorId      int64  `gorm:"column:user_id"`
	PlayUrl       string `gorm:"column:play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	CreateTime    string `gorm:"column:create_time"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	CommentDB     string `gorm:"column:comment_db"`
}

func (SQLVideo) TableName() string {
	return "video"
}
