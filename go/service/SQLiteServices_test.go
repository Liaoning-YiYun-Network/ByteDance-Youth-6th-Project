package service

import (
	"SkyLine/dao"
	"SkyLine/entity"
	"testing"
	"time"
)

func TestAddCommentByDBName(t *testing.T) {
	dbName, err := dao.CreateDB(dao.COMMENTS, 114514)
	if err != nil {
		t.Error(err)
		return
	}
	dateStr := time.Now().Format("2006-01-02 15:04:05")
	err = AddCommentByDBName(dbName, entity.DBComment{
		UserID:  114514,
		Content: "test",
		Time:    dateStr,
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log("Add comment success")
	}
}
