package service

import (
	"SkyLine/data"
	"SkyLine/entity"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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
