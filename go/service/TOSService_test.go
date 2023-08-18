package service

import (
	"SkyLine/config"
	"testing"
)

func TestUploadFile(t *testing.T) {
	config.InitConfig()
	err := InitTOS()
	if err != nil {
		t.Error(err)
	}
	err = UploadFile("test", []byte("test"), VIDEO)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("success")
	}
}
