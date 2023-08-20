package entity

type SQLVideo struct {
	VideoId       int64      `gorm:"column:id;primary_key"` // 视频唯一标识
	UserDetail    UserDetail `gorm:"-"`                     // 视频作者信息
	Title         string     `gorm:"column:title"`          // 视频标题
	AuthorId      int64      `gorm:"column:userid"`         //作者id
	PlayUrl       string     `gorm:"column:play_url"`       // 视频播放地址
	CoverUrl      string     `gorm:"column:cover_url"`      // 视频封面地址
	FavoriteCount int64      `gorm:"column:favorite_count"` // 视频的点赞总数
	CommentCount  int64      `gorm:"column:comment_count"`  // 视频的评论总数
	CommentDB     string     `gorm:"column:comment_db"`     // 评论数据库
	IsFollow      bool       `gorm:"-"`                     // true-已关注，false-未关注
}

func (SQLVideo) TableName() string {
	return "video"
}
