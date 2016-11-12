// +build darwin freebsd linux netbsd openbsd

package system

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
)

var homedir_cache string = ""

func homeDir() (string, error) {
	if homedir_cache != "" {
		return homedir_cache, nil
	}
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		homedir_cache = home
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output")
	}
	homedir_cache = result
	return result, nil
}
