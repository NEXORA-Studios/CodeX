package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"CodeX/backend/common"

	"github.com/pelletier/go-toml/v2"
)

func validateProfileToml() error {
	path := filepath.Join(GetDataPath(), "Profile.toml")

	var invaild bool = false

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var cfg common.IProfile
	if err := toml.Unmarshal(data, &cfg); err != nil {
		invaild = true
	}

	if len(data) == 0 || cfg.Version == 0 || len(cfg.Profile) == 0 {
		invaild = true
	}

	for _, p := range cfg.Profile {
		if p.GUID == "" || p.Name == "" {
			invaild = true
		}
		if p.Settings.Theme == "" || p.Settings.Language == "" {
			invaild = true
		}
	}

	if invaild {
		fmt.Println("Profile.toml 格式错误，备份并重置文件")
		backupPath := filepath.Join(GetDataPath(), "Profile.toml.bak")
		os.Rename(path, backupPath)
		// 初始化一个空结构体
		pftoml := common.IProfile{
			Version: 1,
			Profile: []common.Profile{},
		}
		newData, errr := toml.Marshal(pftoml)
		if errr != nil {
			return errr
		}
		os.WriteFile(path, newData, 0644)
	}

	return nil
}
