package dao

import "testing"

func TestCreateDB(t *testing.T) {
	db, err := CreateDB(FOLLOWS, 1)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(db)
	}
}
