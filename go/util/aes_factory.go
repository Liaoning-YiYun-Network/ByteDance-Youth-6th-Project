package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

// EncryptWithAES 使用AES加密数据
//
// key: 密钥，长度必须为16、24或32字节以分别选择AES-128、AES-192或AES-256
//
// data: 待加密的数据
//
// 返回值1: 加密后的数据，十六进制字符串
//
// 返回值2: 错误信息
func EncryptWithAES(key []byte, data string) (string, error) {

	// 待加密的数据
	plaintext := []byte(data)

	// 创建AES加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 创建一个伪随机数生成器，用于生成初始向量（IV）
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// 使用CBC模式加密
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	// 将IV与密文连接起来
	copy(ciphertext[:aes.BlockSize], iv)

	// 将密文转换为十六进制字符串
	ciphertextHex := hex.EncodeToString(ciphertext)

	return ciphertextHex, nil
}

// DecryptWithAES 使用AES解密数据
//
// key: 密钥，长度必须为16、24或32字节以分别选择AES-128、AES-192或AES-256
//
// data: 待解密的数据，十六进制字符串
//
// 返回值1: 解密后的数据
//
// 返回值2: 错误信息
func DecryptWithAES(key []byte, data string) (string, error) {

	// 将十六进制字符串转换为字节切片
	ciphertext, err := hex.DecodeString(data)
	if err != nil {
		return "", err
	}

	// 创建AES加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 获取初始向量（IV）
	iv := ciphertext[:aes.BlockSize]

	// 获取密文
	ciphertext = ciphertext[aes.BlockSize:]

	// 创建一个使用CBC模式解密的流
	mode := cipher.NewCBCDecrypter(block, iv)

	// 解密
	mode.CryptBlocks(ciphertext, ciphertext)

	// 返回解密后的数据
	return string(ciphertext), nil
}
