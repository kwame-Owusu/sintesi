//go:build darwin

package sysinfo

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"time"
)

func OS() string {
	name, err := exec.Command("sw_vers", "-productName").CombinedOutput()
	if err != nil {
		return "Unknown"
	}

	name = bytes.TrimSuffix(name, []byte{'\n'})

	version, err := exec.Command("sw_vers", "-productVersion").CombinedOutput()
	if err != nil {
		return string(name)
	}

	version = bytes.TrimSuffix(version, []byte{'\n'})
	return string(name) + " " + string(version)
}

func Kernel() string {
	out, err := exec.Command("uname", "-r").CombinedOutput()
	if err != nil {
		return "Unknown"
	}
	return strings.TrimSuffix(string(out), "\n")
}

func Shell() string {
	shellenv := strings.Split(os.Getenv("SHELL"), "/")
	return shellenv[len(shellenv)-1]
}

func Terminal() string {
	return os.Getenv("TERM")
}

func Uptime() (time.Duration, error) {
	// Simple version: just return 0 for now
	// macOS uptime parsing is complex, we'll skip it
	// TODO parse the return from "uptime" command in macOS
	return 0, nil
}
