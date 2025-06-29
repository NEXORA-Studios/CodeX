package utils

import (
	"os"
	"path/filepath"
)

var ensureFoldersList = []string{
	"",
	filepath.Join("Avatar"),
}

var ensureFilesList = []string{
	"Profile.toml",
	"Config.toml",
}

func ensureFolders() error {
	for _, folder := range ensureFoldersList {
		if err := os.MkdirAll(filepath.Join(GetDataPath(), folder), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func ensureFiles() error {
	for _, file := range ensureFilesList {
		var fp = filepath.Join(GetDataPath(), file)
		if _, err := os.Stat(fp); os.IsNotExist(err) {
			_, err := os.Create(fp)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func EnsurePaths() error {
	err := ensureFolders()
	if err != nil {
		return err
	}
	err = ensureFiles()
	if err != nil {
		return err
	}
	return nil
}
