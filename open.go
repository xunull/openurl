package openurl

import (
	"bytes"
	"errors"
	"os/exec"
	"runtime"
)

const (
	OsWindow = "windows"
	OsDarwin = "darwin"
	OsLinux  = "linux"
)

func Open(url string) (error, string) {
	platform := runtime.GOOS

	var command string

	var args []string

	switch platform {
	case OsWindow:
		command = "c:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell"
		args = append(args, "Start")
		args = append(args, url)
	case OsDarwin:
		command = "open"
		args = append(args, url)
	case OsLinux:
		command = "xdg-open"
		args = append(args, url)
	default:
		return errors.New("unknown platform"), ""
	}

	cmd := exec.Command(command, args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stdout
	err := cmd.Run()
	return err, stdout.String()
}
