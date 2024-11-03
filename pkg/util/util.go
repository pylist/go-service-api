package util

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

type FileSave interface {
	Save(data []byte) (string, error)
}

type AvatarSave struct {
	path string
}

func NewAvatarSave(path string) *AvatarSave {
	return &AvatarSave{path: path}
}

func (a *AvatarSave) Save(data []byte) (string, error) {
	// 计算图片哈希
	hash := sha256.Sum256(data)
	hashStr := hex.EncodeToString(hash[:])
	if a.path == "" {
		a.path = "./static/image/avatar/"
	}
	filePath := a.path + hashStr
	url := "/api/static/image/avatar/" + hashStr
	// 确认文件是否存在， 如果存在就返回path
	if FileExists(filePath) {
		return url, nil
	}
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return "", err
	}
	return url, nil
}

func SaveImage(data []byte) (string, error) {
	// 计算图片哈希
	hash := sha256.Sum256(data)
	hashStr := hex.EncodeToString(hash[:])
	// 确认文件是否存在， 如果存在就返回path
	if FileExists(hashStr) {
		return hashStr, nil
	}
	err := os.WriteFile(hashStr, data, 0644)
	if err != nil {
		return "", err
	}
	return hashStr, nil
}

// fileExists 检查文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil { // 文件存在
		return true
	}
	if os.IsNotExist(err) { // 文件不存在
		return false
	}
	// 其他错误
	return false
}
