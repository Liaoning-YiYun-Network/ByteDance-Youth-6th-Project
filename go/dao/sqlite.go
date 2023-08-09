package dao

import "fmt"

type DBType string

// FOLLOWS 正在关注的人
const FOLLOWS DBType = "follows"

// FOLLOWERS 关注我的人
const FOLLOWERS DBType = "followers"

// FAVORITES 我的收藏
const FAVORITES DBType = "favorites"

// COMMENTS 视频评论
const COMMENTS DBType = "comments"

// CreateDB 根据给定的数据库类型创建一个sqlite数据库并返回数据库名称和是否错误
func CreateDB(dbType DBType, id int) (string, error) {
	switch dbType {
	case FOLLOWS:
		//从resources/default_dbs目录下复制follows.sqlite到dbs目录下，并重命名为随机UUID
	case FOLLOWERS:
		//从resources/default_dbs目录下复制followers.sqlite到dbs目录下，并重命名为随机UUID
	case COMMENTS:
		//从resources/default_dbs目录下复制comments.sqlite到dbs目录下，并重命名为随机UUID
	case FAVORITES:
		//从resources/default_dbs目录下复制favorites.sqlite到dbs目录下，并重命名为随机UUID
	}
	return "Not yet Implemented", fmt.Errorf("not yet implemented")
}
