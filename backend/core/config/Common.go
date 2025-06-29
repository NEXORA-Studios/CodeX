package config

import (
	"os"
	"path/filepath"

	"CodeX/backend/common"
	"CodeX/backend/utils"

	"github.com/pelletier/go-toml/v2"
)

func loadConfig() (*common.IConfig, error) {
	path := filepath.Join(utils.GetDataPath(), "config.toml")
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := &common.IConfig{}
	err = toml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
