package env

import ()

func GetNodeVersion() (string, error) {
	return GetDevEnv("node", "-v")
}

func GetNPMVersion() (string, error) {
	return GetDevEnv("npm", "-v")
}

func GetYarnVersion() (string, error) {
	return GetDevEnv("yarn", "-v")
}

func GetPNPMVersion() (string, error) {
	return GetDevEnv("pnpm", "-v")
}

func GetNVMVersion() (string, error) {
	return GetDevEnv("nvm", "-v")
}
