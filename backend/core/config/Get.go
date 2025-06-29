package config

import (
	"CodeX/backend/common"
	"CodeX/backend/utils"
)

func GetConfig() (*common.IConfig, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	return config, nil
}

func GetConfigWithKey(key string) string {
	config, err := loadConfig()
	if err != nil {
		return ""
	}
	val, err := utils.GetConfigFieldValue(config, key)
	if err != nil {
		return ""
	}
	return val.(string)
}
