package bind

import (
	"CodeX/backend/common"
	"CodeX/backend/core/config"
)

type ConfigBind struct{}

func (_ ConfigBind) GetConfig() (*common.IConfig, error) {
	config, err := config.GetConfig()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (_ ConfigBind) GetConfigWithKey(key string) string {
	return config.GetConfigWithKey(key)
}

func (_ ConfigBind) SetConfig(cfg *common.IConfig) error {
	return config.SetConfig(cfg)
}
