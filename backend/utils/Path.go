package utils

import (
	"os"
	"path/filepath"
)

func GetBasePath() string {
	// 返回目前 App 所在目录
	exePath, _ := os.Executable()
	basePath := filepath.Dir(exePath)
	return basePath
}

func GetDataPath() string {
	dataPath := filepath.Join(GetBasePath(), "UserData")
	return dataPath
}

