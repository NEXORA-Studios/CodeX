package profile

import (
	"errors"

	"CodeX/backend/common"
)

func GetAllProfile() ([]common.Profile, error) {
	profiles, err := loadProfiles()
	if err != nil {
		return nil, err
	}

	return profiles, nil
}

func GetProfile(guid string) (*common.Profile, error) {
	profiles, err := loadProfiles()
	if err != nil {
		return nil, err
	}

	for i := range profiles {
		if profiles[i].GUID == guid {
			return &profiles[i], nil
		}
	}

	return nil, errors.New("profile not found")
}
