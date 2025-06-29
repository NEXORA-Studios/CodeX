package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"CodeX/backend/common"

	"github.com/pelletier/go-toml/v2"
)

func validateConfigToml() error {
	path := filepath.Join(GetDataPath(), "Config.toml")

	var invalid bool = false

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var cfg common.IConfig
	if err := toml.Unmarshal(data, &cfg); err != nil {
		invalid = true
	}

	// 检查字段是否为空或无效
	// if len(data) == 0 || cfg.Workspace == "" || len(cfg.AvailableIDE) == 0 { // 从未初始化时可以是空
	if len(data) == 0 {
		invalid = true
	}

	for _, ide := range cfg.AvailableIDE {
		if ide.Name == "" || ide.LaunchTarget == "" {
			invalid = true
		}
	}

	if invalid {
		fmt.Println("Config.toml 格式错误，备份并重置文件")
		backupPath := filepath.Join(GetDataPath(), "Config.toml.bak")
		_ = os.Rename(path, backupPath)

		// 初始化空结构体
		newConfig := common.IConfig{
			AvailableIDE: []common.IDEInstance{},
			Workspace:    "",
		}
		newData, err := toml.Marshal(newConfig)
		if err != nil {
			return err
		}
		_ = os.WriteFile(path, newData, 0644)
	}

	return nil
}
