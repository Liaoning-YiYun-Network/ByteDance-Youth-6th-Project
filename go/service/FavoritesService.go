package service

import (
	"SkyLine/data"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

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
