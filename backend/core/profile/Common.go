package profile

import (
	"os"
	"path/filepath"

	"CodeX/backend/common"
	"CodeX/backend/utils"

	toml "github.com/pelletier/go-toml/v2"
)

func loadProfiles() ([]common.Profile, error) {
	path := filepath.Join(utils.GetDataPath(), "Profile.toml")

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var pftoml common.IProfile
	if err := toml.Unmarshal(data, &pftoml); err != nil {
		return nil, err
	}

	return pftoml.Profile, nil
}
