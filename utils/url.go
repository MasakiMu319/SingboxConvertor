package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"net/url"
)

const Key = "EjErrWr9g6B63Svi"

func Encrypt(stringToEncrypt, keyString string) (string, error) {
	key := []byte(keyString)
	plaintext := []byte(stringToEncrypt)

	// 创建一个新的AES块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 创建加密器
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// 将加密结果转换为Base64
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(encryptedString string, keyString string) (string, error) {
	key := []byte(keyString)

	// 解码Base64
	encryptedData, err := base64.URLEncoding.DecodeString(encryptedString)
	if err != nil {
		return "", err
	}

	// 创建一个新的AES块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 检查数据长度
	if len(encryptedData) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	// 解密数据
	iv := encryptedData[:aes.BlockSize]
	encryptedData = encryptedData[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encryptedData, encryptedData)

	return string(encryptedData), nil
}

func IsValidURL(urlStr string) bool {
	_, err := url.ParseRequestURI(urlStr)

	if err != nil {
		return false
	}

	u, err := url.Parse(urlStr)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
