package env

import ()

func GetRustVersion() (string, error) {
	return GetDevEnv("rustc", "--version")
}

func GetCargoVersion() (string, error) {
	return GetDevEnv("cargo", "--version")
}

func GetRustupVersion() (string, error) {
	return GetDevEnv("rustup", "--version")
}
