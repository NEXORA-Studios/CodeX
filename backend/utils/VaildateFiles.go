package utils

func ValidateFiles() error {
	err := validateProfileToml()
	if err != nil {
		return err
	}
	err = validateConfigToml()
	if err != nil {
		return err
	}
	return nil
}
