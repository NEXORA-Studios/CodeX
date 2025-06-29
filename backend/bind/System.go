package bind

import (
	"strings"

	"CodeX/backend/utils"
)

type SystemBind struct{}

func (_ SystemBind) ShellExec(command string) (string, error) {
	output, err := utils.ShellExec(command)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(output), nil
}
