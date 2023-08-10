package util

import (
	"io"
	"os"
)

// CopyFile 将文件拷贝到指定目录
func CopyFile(src, dst string) error {
	exist := isDirExist(dst)
	if !exist {
		err := os.MkdirAll(dst, os.ModePerm)
		if err != nil {
			return err
		}
	}
	// 打开源文件
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	// 创建目标文件
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	// 拷贝文件
	buf := make([]byte, 1024)
	for {
		// 读取字节到buf
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}
		// 写入字节到目标文件
		if _, err := dstFile.Write(buf[:n]); err != nil {
			return err
		}
	}
	return nil
}

// isDirExist 判断目录是否存在
func isDirExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
