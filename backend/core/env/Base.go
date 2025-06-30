package env

import (
	"os/exec"
	"runtime"
	"strings"
)

func GetDevEnv(envName string, envArgs ...string) (string, error) {
	var cmd *exec.Cmd
	var path string
	var hasErr bool = false

	// 先尝试查找运行时路径
	if runtime.GOOS == "windows" {
		_p, err := exec.Command("where", envName).Output()
		if err != nil {
			hasErr = true
			return "", err
		}
		// Windows下按换行符分割结果，取第一个
		paths := strings.Split(strings.TrimSpace(string(_p)), "\n")
		path = string(paths[0])
	} else {
		_p, err := exec.Command("command", "-v", envName).Output()
		if err != nil {
			hasErr = true
			return "", err
		}
		// Unix下跳过目标字符串长度+2后按空格分割
		output := string(_p)
		prefixLen := len(envName) + 2
		if len(output) > prefixLen {
			paths := strings.Split(strings.TrimSpace(output[prefixLen:]), " ")
			path = string(paths[0])
		} else {
			path = string(_p)
		}
	}

	// 如果找到运行时路径则使用，否则使用原始命令
	if hasErr == false && len(path) > 0 {
		if runtime.GOOS == "windows" {
			// Windows 使用 cmd /C
			full := append([]string{"/C", strings.TrimSpace(string(path))}, envArgs...)
			cmd = exec.Command("cmd", full...)
		} else {
			// Unix 使用 bash -c
			full := strings.TrimSpace(string(path)) + " " + strings.Join(envArgs, " ")
			cmd = exec.Command("bash", "-c", full)
		}
	} else {
		// 回退到原始命令
		if runtime.GOOS == "windows" {
			// Windows 使用 cmd /C
			full := append([]string{"/C", envName}, envArgs...)
			cmd = exec.Command("cmd", full...)
		} else {
			// Unix 使用 bash -c
			full := envName + " " + strings.Join(envArgs, " ")
			cmd = exec.Command("bash", "-c", full)
		}
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
