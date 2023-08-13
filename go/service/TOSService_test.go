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
		return
	}
	defer CloseTOS()
	err = UploadFile("test", []byte("test"), VIDEO_COVER)
	if err != nil {
		t.Error(err)
		return
	}
}
