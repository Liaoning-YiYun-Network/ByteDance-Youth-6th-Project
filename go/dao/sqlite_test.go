package dao

import (
	"strconv"
	"testing"
)

func TestCreateDB(t *testing.T) {
	db, err := CreateDB(FOLLOWS, strconv.Itoa(1))
	if err != nil {
		t.Error(err)
	} else {
		t.Log(db)
	}
}
