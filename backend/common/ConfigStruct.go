package common

type IConfig struct {
	AvailableIDE []IDEInstance `toml:"available_ide"`
	Workspace    string        `toml:"workspace"`
}

type IDEInstance struct {
	GUID         string `toml:"guid"`
	Name         string `toml:"name"`
	LaunchTarget string `toml:"launch_target"`
}
