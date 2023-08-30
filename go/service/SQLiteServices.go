package service

import (
	"SkyLine/data"
	"SkyLine/entity"
	"database/sql"
	"fmt"
)

// GetAllCommentsByDBName 根据数据库名称获取所有评论
func GetAllCommentsByDBName(dbName string) (comments []entity.DBComment, err error) {
	//如果已经打开过这个数据库，就直接从缓存中取出
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		rows, err := db.Query("SELECT * FROM comments ORDER BY time DESC")
		if err != nil {
			return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
		}
		for rows.Next() {
			var comment entity.DBComment
			err = rows.Scan(&comment.CommentID, &comment.UserID, &comment.Content, &comment.Time)
			if err != nil {
				return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
			}
			comments = append(comments, comment)
		}
		return comments, nil
	} else {
		db, err := sql.Open("sqlite3", "./dbs/comments/"+dbName)
		if err != nil {
			return nil, fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		data.TempSQLiteConnects[dbName] = db
		rows, err := db.Query("SELECT * FROM comments ORDER BY time DESC")
		if err != nil {
			return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
		}
		for rows.Next() {
			var comment entity.DBComment
			err = rows.Scan(&comment.CommentID, &comment.UserID, &comment.Content, &comment.Time)
			if err != nil {
				return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
			}
			comments = append(comments, comment)
		}
		return comments, nil
	}
}

// AddCommentByDBName 根据数据库名称添加评论
func AddCommentByDBName(dbName string, comment entity.DBComment) error {
	//如果已经打开过这个数据库，就直接从缓存中取出
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		_, err := db.Exec("INSERT INTO comments (userid, content, time) VALUES (?, ?, ?)", comment.UserID, comment.Content, comment.Time)
		if err != nil {
			return fmt.Errorf("尝试插入SQLite数据库时发生错误：%s", err)
		}
		return nil
	} else {
		db, err := sql.Open("sqlite3", "./dbs/comments/"+dbName)
		if err != nil {
			return fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		data.TempSQLiteConnects[dbName] = db
		_, err = db.Exec("INSERT INTO comments (userid, content, time) VALUES (?, ?, ?)", comment.UserID, comment.Content, comment.Time)
		if err != nil {
			return fmt.Errorf("尝试插入SQLite数据库时发生错误：%s", err)
		}
		return nil
	}
}

// DeleteCommentByDBName 根据数据库名称删除评论
func DeleteCommentByDBName(dbName string, commentID int64) error {
	//如果已经打开过这个数据库，就直接从缓存中取出
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		_, err := db.Exec("DELETE FROM comments WHERE comment_id = ?", commentID)
		if err != nil {
			return fmt.Errorf("尝试删除SQLite数据库时发生错误：%s", err)
		}
		return nil
	} else {
		db, err := sql.Open("sqlite3", "./dbs/comments/"+dbName)
		if err != nil {
			return fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		data.TempSQLiteConnects[dbName] = db
		_, err = db.Exec("DELETE FROM comments WHERE comment_id = ?", commentID)
		if err != nil {
			return fmt.Errorf("尝试删除SQLite数据库时发生错误：%s", err)
		}
		return nil
	}
}

// GetAllFollowsByDBName 根据数据库名称获取所有关注
func GetAllFollowsByDBName(dbName string) ([]int64, error) {
	//如果已经打开过这个数据库，就直接从缓存中取出
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		rows, err := db.Query("SELECT * FROM follows")
		if err != nil {
			return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
		}
		var follows []int64
		for rows.Next() {
			var follow int64
			err = rows.Scan(&follow)
			if err != nil {
				return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
			}
			follows = append(follows, follow)
		}
		return follows, nil
	} else {
		db, err := sql.Open("sqlite3", "./dbs/follows/"+dbName)
		if err != nil {
			return nil, fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		data.TempSQLiteConnects[dbName] = db
		rows, err := db.Query("SELECT * FROM follows")
		if err != nil {
			return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
		}
		var follows []int64
		for rows.Next() {
			var follow int64
			err = rows.Scan(&follow)
			if err != nil {
				return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
			}
			follows = append(follows, follow)
		}
		return follows, nil
	}
}

// AddFollowByDBName 根据数据库名称添加关注
func AddFollowByDBName(dbName string, followId int64) error {
	//如果已经打开过这个数据库，就直接从缓存中取出
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		_, err := db.Exec("INSERT INTO follows (userid) VALUES (?)", followId)
		if err != nil {
			return fmt.Errorf("尝试插入SQLite数据库时发生错误：%s", err)
		}
		return nil
	} else {
		db, err := sql.Open("sqlite3", "./dbs/follows/"+dbName)
		if err != nil {
			return fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		data.TempSQLiteConnects[dbName] = db
		_, err = db.Exec("INSERT INTO follows (userid) VALUES (?)", followId)
		if err != nil {
			return fmt.Errorf("尝试插入SQLite数据库时发生错误：%s", err)
		}
		return nil
	}
}

// DeleteFollowByDBName 根据数据库名称删除关注
func DeleteFollowByDBName(dbName string, followId int64) error {
	//如果已经打开过这个数据库，就直接从缓存中取出
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		_, err := db.Exec("DELETE FROM follows WHERE userid = ?", followId)
		if err != nil {
			return fmt.Errorf("尝试删除SQLite数据库时发生错误：%s", err)
		}
		return nil
	} else {
		db, err := sql.Open("sqlite3", "./dbs/follows/"+dbName)
		if err != nil {
			return fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		data.TempSQLiteConnects[dbName] = db
		_, err = db.Exec("DELETE FROM follows WHERE userid = ?", followId)
		if err != nil {
			return fmt.Errorf("尝试删除SQLite数据库时发生错误：%s", err)
		}
		return nil
	}
}

// 根据用户id查看作者列表是其粉丝
func GetFollowByUserId(dbName string, userId int64) (bool, error) {
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		res, err := db.Exec("SELECT FROM follows WHERE userid = ?", userId)
		if err != nil {
			return false, fmt.Errorf("尝试删除SQLite数据库时发生错误：%s", err)
		}
		fmt.Println(res)
		return false, nil
	} else {
		db, err := sql.Open("sqlite3", "./dbs/follows/"+dbName)
		if err != nil {
			return false, fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		_, err = db.Exec("SELECT FROM follows WHERE userid = ?", userId)
	}
	return false, nil
}

// GetAllFollowersByDBName 根据数据库名称获取所有粉丝
func GetAllFollowersByDBName(dbName string) ([]int64, error) {
	//如果已经打开过这个数据库，就直接从缓存中取出
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		rows, err := db.Query("SELECT * FROM followers")
		if err != nil {
			return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
		}
		var followers []int64
		for rows.Next() {
			var follower int64
			err = rows.Scan(&follower)
			if err != nil {
				return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
			}
			followers = append(followers, follower)
		}
		return followers, nil
	} else {
		db, err := sql.Open("sqlite3", "./dbs/followers/"+dbName)
		if err != nil {
			return nil, fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		data.TempSQLiteConnects[dbName] = db
		rows, err := db.Query("SELECT * FROM followers")
		if err != nil {
			return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
		}
		var followers []int64
		for rows.Next() {
			var follower int64
			err = rows.Scan(&follower)
			if err != nil {
				return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
			}
			followers = append(followers, follower)
		}
		return followers, nil
	}
}

// AddFollowerByDBName 根据数据库名称添加粉丝
func AddFollowerByDBName(dbName string, followerId int64) error {
	//如果已经打开过这个数据库，就直接从缓存中取出
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		_, err := db.Exec("INSERT INTO followers (userid) VALUES (?)", followerId)
		if err != nil {
			return fmt.Errorf("尝试插入SQLite数据库时发生错误：%s", err)
		}
		return nil
	} else {
		db, err := sql.Open("sqlite3", "./dbs/followers/"+dbName)
		if err != nil {
			return fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		data.TempSQLiteConnects[dbName] = db
		_, err = db.Exec("INSERT INTO followers (userid) VALUES (?)", followerId)
		if err != nil {
			return fmt.Errorf("尝试插入SQLite数据库时发生错误：%s", err)
		}
		return nil
	}
}

// DeleteFollowerByDBName 根据数据库名称删除粉丝
func DeleteFollowerByDBName(dbName string, followerId int64) error {
	//如果已经打开过这个数据库，就直接从缓存中取出
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		_, err := db.Exec("DELETE FROM followers WHERE userid = ?", followerId)
		if err != nil {
			return fmt.Errorf("尝试删除SQLite数据库时发生错误：%s", err)
		}
		return nil
	} else {
		db, err := sql.Open("sqlite3", "./dbs/followers/"+dbName)
		if err != nil {
			return fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		data.TempSQLiteConnects[dbName] = db
		_, err = db.Exec("DELETE FROM followers WHERE userid = ?", followerId)
		if err != nil {
			return fmt.Errorf("尝试删除SQLite数据库时发生错误：%s", err)
		}
		return nil
	}
}

// GetAllFavoritesByDBName 根据数据库名称获取所有喜欢
func GetAllFavoritesByDBName(dbName string) ([]int64, error) {
	//如果已经打开过这个数据库，就直接从缓存中取出
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		rows, err := db.Query("SELECT * FROM favorites")
		if err != nil {
			return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
		}
		var favorites []int64
		for rows.Next() {
			var favorite int64
			err = rows.Scan(&favorite)
			if err != nil {
				return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
			}
			favorites = append(favorites, favorite)
		}
		return favorites, nil
	} else {
		db, err := sql.Open("sqlite3", "./dbs/favorites/"+dbName)
		if err != nil {
			return nil, fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		data.TempSQLiteConnects[dbName] = db
		rows, err := db.Query("SELECT * FROM favorites")
		if err != nil {
			return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
		}
		var favorites []int64
		for rows.Next() {
			var favorite int64
			err = rows.Scan(&favorite)
			if err != nil {
				return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
			}
			favorites = append(favorites, favorite)
		}
		return favorites, nil
	}
}

// AddFavoriteByDBName 根据数据库名称添加喜欢
func AddFavoriteByDBName(dbName string, favoriteId int64) error {
	//如果已经打开过这个数据库，就直接从缓存中取出
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		_, err := db.Exec("INSERT INTO favorites (id) VALUES (?)", favoriteId)
		if err != nil {
			return fmt.Errorf("尝试插入SQLite数据库时发生错误：%s", err)
		}
		return nil
	} else {
		db, err := sql.Open("sqlite3", "./dbs/favorites/"+dbName)
		if err != nil {
			return fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		data.TempSQLiteConnects[dbName] = db
		_, err = db.Exec("INSERT INTO favorites (id) VALUES (?)", favoriteId)
		if err != nil {
			return fmt.Errorf("尝试插入SQLite数据库时发生错误：%s", err)
		}
		return nil
	}
}

// DeleteFavoriteByDBName 根据数据库名称删除喜欢
func DeleteFavoriteByDBName(dbName string, favoriteId int64) error {
	//如果已经打开过这个数据库，就直接从缓存中取出
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		_, err := db.Exec("DELETE FROM favorites WHERE id = ?", favoriteId)
		if err != nil {
			return fmt.Errorf("尝试删除SQLite数据库时发生错误：%s", err)
		}
		return nil
	} else {
		db, err := sql.Open("sqlite3", "./dbs/favorites/"+dbName)
		if err != nil {
			return fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		data.TempSQLiteConnects[dbName] = db
		_, err = db.Exec("DELETE FROM favorites WHERE id = ?", favoriteId)
		if err != nil {
			return fmt.Errorf("尝试删除SQLite数据库时发生错误：%s", err)
		}
		return nil
	}
}
