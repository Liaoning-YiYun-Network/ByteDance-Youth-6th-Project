package service

import (
	"SkyLine/data"
	"SkyLine/entity"
	"database/sql"
	"fmt"
)

func AddMessageByDBName(dbName string, message entity.DBMessage) error {
	//如果已经打开过这个数据库，就直接从缓存中取出
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		_, err := db.Exec("INSERT INTO messages (user_id, content, create_time) VALUES (?, ?, ?)", message.UserID, message.Content, message.CreateTime)
		if err != nil {
			return fmt.Errorf("尝试插入SQLite数据库时发生错误：%s", err)
		}
		return nil
	} else {
		db, err := sql.Open("sqlite3", "./dbs/messages/"+dbName)
		if err != nil {
			return fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		data.TempSQLiteConnects[dbName] = db
		_, err = db.Exec("INSERT INTO messages (user_id, content, create_time) VALUES (?, ?, ?)", message.UserID, message.Content, message.CreateTime)
		if err != nil {
			return fmt.Errorf("尝试插入SQLite数据库时发生错误：%s", err)
		}
		return nil
	}
}

func GetAllMessagesByDBName(dbName string) (messages []entity.DBMessage, err error) {
	//如果已经打开过这个数据库，就直接从缓存中取出
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		rows, err := db.Query("SELECT * FROM messages")
		if err != nil {
			return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
		}
		for rows.Next() {
			var message entity.DBMessage
			err = rows.Scan(&message.MessageID, &message.UserID, &message.Content, &message.CreateTime)
			if err != nil {
				return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
			}
			messages = append(messages, message)
		}
		return messages, nil
	} else {
		db, err := sql.Open("sqlite3", "./dbs/messages/"+dbName)
		if err != nil {
			return nil, fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		data.TempSQLiteConnects[dbName] = db
		rows, err := db.Query("SELECT * FROM messages")
		if err != nil {
			return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
		}
		for rows.Next() {
			var message entity.DBMessage
			err = rows.Scan(&message.MessageID, &message.UserID, &message.Content, &message.CreateTime)
			if err != nil {
				return nil, fmt.Errorf("尝试查询SQLite数据库时发生错误：%s", err)
			}
			messages = append(messages, message)
		}
		return messages, nil
	}
}

func DeleteMessageByDBName(dbName string, messageId int64) error {
	//如果已经打开过这个数据库，就直接从缓存中取出
	if db, ok := data.TempSQLiteConnects[dbName]; ok {
		_, err := db.Exec("DELETE FROM messages WHERE id = ?", messageId)
		if err != nil {
			return fmt.Errorf("尝试删除SQLite数据库时发生错误：%s", err)
		}
		return nil
	} else {
		db, err := sql.Open("sqlite3", "./dbs/messages/"+dbName)
		if err != nil {
			return fmt.Errorf("尝试打开SQLite数据库时发生错误：%s", err)
		}
		data.TempSQLiteConnects[dbName] = db
		_, err = db.Exec("DELETE FROM messages WHERE id = ?", messageId)
		if err != nil {
			return fmt.Errorf("尝试删除SQLite数据库时发生错误：%s", err)
		}
		return nil
	}
}
