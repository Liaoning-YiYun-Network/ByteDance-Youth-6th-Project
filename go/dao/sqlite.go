package dao

import (
	"SkyLine/util"
	"fmt"
)

type DBType string

// FOLLOWS 正在关注的人
const FOLLOWS DBType = "follows"

// FOLLOWERS 关注我的人
const FOLLOWERS DBType = "followers"

// FAVORITES 我的收藏
const FAVORITES DBType = "favorites"

// COMMENTS 视频评论
const COMMENTS DBType = "comments"

const MESSAGES DBType = "messages"

// CreateDB 根据给定的数据库类型创建一个sqlite数据库并返回数据库名称和是否错误
func CreateDB(dbType DBType, id string) (string, error) {

	var namePrefix string
	switch dbType {
	case FOLLOWS:
		randomStr, err := util.UUIDWithoutHyphen()
		if err != nil {
			return "", fmt.Errorf("尝试创建SQLite数据库时发生错误：%s", err)
		}
		//从resources/default_dbs目录下复制follows.sqlite到dbs目录下，并重命名为随机UUID
		namePrefix = string(FOLLOWS) + "-" + id + "-" + randomStr + ".sqlite"
		err = util.CopyFile("./resources/default_dbs/follows.sqlite", "./dbs/follows/"+namePrefix)
		if err != nil {
			return "", fmt.Errorf("尝试创建SQLite数据库时发生错误：%s", err)
		}
		return namePrefix, nil
	case FOLLOWERS:
		randomStr, err := util.UUIDWithoutHyphen()
		if err != nil {
			return "", fmt.Errorf("尝试创建SQLite数据库时发生错误：%s", err)
		}
		//从resources/default_dbs目录下复制followers.sqlite到dbs目录下，并重命名为随机UUID
		namePrefix = string(FOLLOWERS) + "-" + id + "-" + randomStr + ".sqlite"
		err = util.CopyFile("./resources/default_dbs/followers.sqlite", "./dbs/followers/"+namePrefix)
		if err != nil {
			return "", fmt.Errorf("尝试创建SQLite数据库时发生错误：%s", err)
		}
		return namePrefix, nil
	case COMMENTS:
		randomStr, err := util.UUIDWithoutHyphen()
		if err != nil {
			return "", fmt.Errorf("尝试创建SQLite数据库时发生错误：%s", err)
		}
		//从resources/default_dbs目录下复制comments.sqlite到dbs目录下，并重命名为随机UUID
		namePrefix = string(COMMENTS) + "-" + id + "-" + randomStr + ".sqlite"
		err = util.CopyFile("./resources/default_dbs/comments.sqlite", "./dbs/comments/"+namePrefix)
		if err != nil {
			return "", fmt.Errorf("尝试创建SQLite数据库时发生错误：%s", err)
		}
		return namePrefix, nil
	case FAVORITES:
		randomStr, err := util.UUIDWithoutHyphen()
		if err != nil {
			return "", fmt.Errorf("尝试创建SQLite数据库时发生错误：%s", err)
		}
		//从resources/default_dbs目录下复制favorites.sqlite到dbs目录下，并重命名为随机UUID
		namePrefix = string(FAVORITES) + "-" + id + "-" + randomStr + ".sqlite"
		err = util.CopyFile("./resources/default_dbs/favorites.sqlite", "./dbs/favorites/"+namePrefix)
		if err != nil {
			return "", fmt.Errorf("尝试创建SQLite数据库时发生错误：%s", err)
		}
		return namePrefix, nil
	case MESSAGES:
		namePrefix = string(MESSAGES) + "-" + id + ".sqlite"
		err := util.CopyFile("./resources/default_dbs/messages.sqlite", "./dbs/messages/"+namePrefix)
		if err != nil {
			return "", fmt.Errorf("尝试创建SQLite数据库时发生错误：%s", err)
		}
		return namePrefix, nil
	}
	return "", fmt.Errorf("尝试创建SQLite数据库时发生错误：未知的数据库类型")
}

// DeleteDB 根据给定的数据库名称删除一个sqlite数据库并返回是否错误
func DeleteDB(dbType DBType, dbName string) error {
	var err error
	switch dbType {
	case FOLLOWS:
		err = util.DeleteFile("./dbs/follows/" + dbName)
	case FOLLOWERS:
		err = util.DeleteFile("./dbs/followers/" + dbName)
	case COMMENTS:
		err = util.DeleteFile("./dbs/comments/" + dbName)
	case FAVORITES:
		err = util.DeleteFile("./dbs/favorites/" + dbName)
	case MESSAGES:
		err = util.DeleteFile("./dbs/messages/" + dbName)
	}
	if err != nil {
		return fmt.Errorf("尝试删除SQLite数据库时发生错误：%s", err)
	}
	return nil
}
