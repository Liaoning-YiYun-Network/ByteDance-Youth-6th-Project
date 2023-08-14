package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncryptWithMD5 使用MD5加密字符串
func EncryptWithMD5(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}
