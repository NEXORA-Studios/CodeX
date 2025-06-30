package env

import (
	"os/exec"
	"runtime"
	"strings"
)

func GetDevEnv(envName string, envArgs ...string) (string, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		// Windows 使用 cmd /C
		full := append([]string{"/C", envName}, envArgs...)
		cmd = exec.Command("cmd", full...)
	} else {
		// Unix 使用 bash -c
		full := envName + " " + strings.Join(envArgs, " ")
		cmd = exec.Command("bash", "-c", full)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
