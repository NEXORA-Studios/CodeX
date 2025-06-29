package env

import ()

func GetGoVersion() (string, error) {
	return GetDevEnv("go", "version")
}
