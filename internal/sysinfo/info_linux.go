package sysinfo

import (
	"os"
	"strings"
	"syscall"
	"time"
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

func Shell() string {
	shellEnv := strings.Split(os.Getenv("SHELL"), "/")
	return shellEnv[len(shellEnv)-1]
}

func Uptime() (time.Duration, error) {
	var info syscall.Sysinfo_t
	err := syscall.Sysinfo(&info)
	if err != nil {
		return 0, err
	}

	return time.Duration(info.Uptime) * time.Second, nil
}

func Terminal() string {
	terminalEnv := os.Getenv("TERM")
	return terminalEnv
}

func LinuxHardware() Hardware {
	cpu := getLinuxCpuInfo()
	ram := getLinuxRamInfo()
	return Hardware{Cpu: cpu, Ram: ram}
}
