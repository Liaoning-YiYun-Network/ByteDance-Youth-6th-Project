package service

import (
	"SkyLine/data"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"strings"
)

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

// GetFollowByUserId 根据用户id查看作者列表是其粉丝
func GetFollowByUserId(dbName string, userId int64) (bool, error) {
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		res, err := db.Exec("SELECT * FROM follows WHERE userid = ?", userId)
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
		_, err = db.Exec("SELECT * FROM follows WHERE userid = ?", userId)
	}
	return false, nil
}

func GetFollowByUserIds(dbName string, ids []int64) ([]int, error) {
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		args := make([]interface{}, len(ids))
		for i, id := range ids {
			args[i] = strconv.FormatInt(id, 10)
		}
		var sql strings.Builder
		sql.WriteString("SELECT * FROM follows WHERE userid IN(")
		sql.WriteString("")
		for i := 0; i < len(ids); i++ {
			sql.WriteString("?")
			if i == len(ids)-1 {
				break
			}
			sql.WriteString(",")
		}
		sql.WriteString(")")
		query, err := db.Query(sql.String(), args...)
		if err != nil {
			fmt.Printf("读取错误:%s", err)
		}
		var userid []int
		for query.Next() {
			var column1 int
			// 将查询结果映射到变量
			err3 := query.Scan(&column1)
			if err3 != nil {
				fmt.Println("扫描结果错误:", err3)
			}
			userid = append(userid, column1)
		}
		return userid, nil

	} else {
		db, err := sql.Open("sqlite3", "./dbs/follows/"+dbName)
		if err != nil {
			return nil, fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		args := make([]interface{}, len(ids))
		for i, id := range ids {
			args[i] = strconv.FormatInt(id, 10)
		}
		var sql strings.Builder
		sql.WriteString("SELECT * FROM follows WHERE userid IN(")
		sql.WriteString("")
		for i := 0; i < len(ids); i++ {
			sql.WriteString("?")
			if i == len(ids)-1 {
				break
			}
			sql.WriteString(",")
		}
		sql.WriteString(")")
		query, err := db.Query(sql.String(), args...)
		fmt.Printf("%#v", query)
		if err != nil {
			fmt.Printf("读取错误:%s", err)
		}
		var userid []int
		for query.Next() {
			var column1 int
			// 将查询结果映射到变量
			err3 := query.Scan(&column1)
			if err3 != nil {
				fmt.Println("扫描结果错误:", err3)
			}
			userid = append(userid, column1)
		}
		return userid, nil

	}
}
