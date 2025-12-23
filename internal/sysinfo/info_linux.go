package sysinfo

import (
	"os"
	"strings"
)

func OS() string {
	return getLinuxOSName()
}

func Kernel() string {
	// /proc/version should always exist on linux
	procver, _ := os.ReadFile("/proc/version")

	// /proc/version has the same format with "Linux version <kern-version>" as the 3rd
	// word, `procver` is []byte so has to be converted
	return strings.Split(string(procver), " ")[2]
}
