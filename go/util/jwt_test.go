package util

import (
	"SkyLine/entity"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	user := new(entity.SQLUser)
	token, err := GenerateToken(*user)
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}
