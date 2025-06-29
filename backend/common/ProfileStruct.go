package common

type IProfile struct {
	Version int       `toml:"version"`
	Profile []Profile `toml:"profile"`
}

type Profile struct {
	GUID            string          `toml:"guid"`
	Name            string          `toml:"name"`
	Avatar          bool            `toml:"avatar"`
	Settings        ProfileSettings `toml:"settings"`
	IDE             string          `toml:"ide"` // IDE 识别名称（@refer ConfigStruct.go IDEInstance.GUID)
	CurrentCourseID ProfileCCID     `toml:"current_course_id"`
}

type ProfileSettings struct {
	Theme    string `toml:"theme"`
	Language string `toml:"language"`
}

type ProfileCCID struct {
	ID    string `toml:"id"`
	SubID string `toml:"sub_id"`
}
