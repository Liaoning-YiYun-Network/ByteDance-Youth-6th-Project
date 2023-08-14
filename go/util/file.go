package util

import (
	"io"
	"os"
)

// CopyFile 将文件拷贝到指定目录
func CopyFile(src, dst string) error {
	//匹配dst最后一个/，并移除其最后一个/后的内容
	//例如：将./dbs/follows/1-1-1-1-1-1-1-1-1-1-1-1-1-1-1.sqlite
	//转换为./dbs/follows/
	var parentDir string
	for i := len(dst) - 1; i >= 0; i-- {
		if dst[i] == '/' {
			parentDir = dst[0 : i+1]
			break
		}
	}
	// 判断目标目录是否存在，不存在则创建
	exist := isDirExist(parentDir)
	if !exist {
		err := os.MkdirAll(parentDir, os.ModePerm)
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

// RenameFile 将文件重命名
func RenameFile(src, dst string) error {
	return os.Rename(src, dst)
}

// isDirExist 判断目录是否存在
func isDirExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
