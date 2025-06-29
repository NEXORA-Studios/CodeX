package profile

import (
	"os"
	"path/filepath"

	"CodeX/backend/common"
	"CodeX/backend/utils"

	"github.com/pelletier/go-toml/v2"
)

func UpsertProfile(p common.Profile) error {
	path := filepath.Join(utils.GetDataPath(), "Profile.toml")

	data, err := os.ReadFile(path)
	if err != nil {
		// 文件不存在时，初始化一个空结构体
		if os.IsNotExist(err) {
			pftoml := common.IProfile{
				Profile: []common.Profile{p},
			}
			newData, errr := toml.Marshal(pftoml)
			if errr != nil {
				return errr
			}
			return os.WriteFile(path, newData, 0644)
		}
		return err
	}

	var pftoml common.IProfile
	if err = toml.Unmarshal(data, &pftoml); err != nil {
		return err
	}

	found := false
	for i := range pftoml.Profile {
		if pftoml.Profile[i].GUID == p.GUID {
			pftoml.Profile[i] = p
			found = true
			break
		}
	}

	if !found {
		pftoml.Profile = append(pftoml.Profile, p)
	}

	newData, err := toml.Marshal(pftoml)
	if err != nil {
		return err
	}

	return os.WriteFile(path, newData, 0644)
}
