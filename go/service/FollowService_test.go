package service

import (
	"SkyLine/dao"
	"strconv"
	"testing"
)

func TestAddFollowByDBName(t *testing.T) {
	dbName, err := dao.CreateDB(dao.FOLLOWS, strconv.Itoa(114514))
	if err != nil {
		t.Errorf("CreateDB() error = %v", err)
		return
	}
	t.Log("CreateDB() PASS, dbName = " + dbName)
	err = AddFollowByDBName(dbName, 1)
	if err != nil {
		t.Errorf("AddFollowByDBName() error = %v", err)
		return
	}
	t.Log("AddFollowByDBName() PASS")
}
