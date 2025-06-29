package bind

import (
	common "CodeX/backend/common"
	cpf "CodeX/backend/core/profile"
)

type ProfileBind struct{}

func (_ ProfileBind) GetAllProfile() ([]common.Profile, error) {
	return cpf.GetAllProfile()
}

func (_ ProfileBind) GetProfile(guid string) (*common.Profile, error) {
	return cpf.GetProfile(guid)
}

func (_ ProfileBind) UpsertProfile(p common.Profile) error {
	return cpf.UpsertProfile(p)
}
