package env

import ()

func GetPythonVersion() (string, error) {
	return GetDevEnv("python", "--version")
}
