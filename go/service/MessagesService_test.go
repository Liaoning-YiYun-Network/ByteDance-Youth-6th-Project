package service

import (
	"SkyLine/dao"
	"SkyLine/entity"
	"testing"
)

func TestAddMessageByDBName(t *testing.T) {
	type args struct {
		dbName string
		msg    entity.DBMessage
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				dbName: "test.db",
				msg: entity.DBMessage{
					UserID:     114514,
					Content:    "test",
					CreateTime: "2021-01-01 00:00:00",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := dao.CreateDB(dao.MESSAGES, tt.args.dbName)
			if err != nil {
				t.Errorf("AddMessageByDBName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err := AddMessageByDBName(db, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("AddMessageByDBName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
