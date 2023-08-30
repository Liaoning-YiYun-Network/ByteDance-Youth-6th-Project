package service

import (
	"SkyLine/data"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

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
