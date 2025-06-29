package config

import (
	"os"
	"path/filepath"

	"CodeX/backend/common"
	"CodeX/backend/utils"

	"github.com/pelletier/go-toml/v2"
)

func SetConfig(config *common.IConfig) error {
	path := filepath.Join(utils.GetDataPath(), "Config.toml")
	content, err := toml.Marshal(config)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, content, 0644)
	if err != nil {
		return err
	}
	return nil
}
