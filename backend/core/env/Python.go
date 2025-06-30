package env

import "runtime"

func GetPythonVersion() (string, error) {
	var name string
	if runtime.GOOS == "windows" {
		name = "python"
	} else {
		name = "python3"
	}
	return GetDevEnv(name, "--version")
}
