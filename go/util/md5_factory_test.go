package util

import "testing"

func TestEncryptWithMD5(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"123456"}, "e10adc3949ba59abbe56e057f20f883e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EncryptWithMD5(tt.args.str); got != tt.want {
				t.Errorf("EncryptWithMD5() = %v, want %v", got, tt.want)
			}
		})
	}
}
