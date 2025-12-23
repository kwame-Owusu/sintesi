package sysinfo

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getFile(filepath string) []byte {
	data, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error occurred reading filepath, %v", err)
	}
	return data
}

// parses the os-release file and returns PRETTY_NAME value
func getLinuxOSName() string {
	osFile := getFile("/etc/os-release")
	lines := strings.Split(string(osFile), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := strings.Trim(parts[1], `"`)

		if key == "PRETTY_NAME" {
			return value
		}
	}
	return "unknown OS"
}

func processCpuInfo(lines []string, cpuInfo *Cpu) {
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if key == "model name" && cpuInfo.Model == "" {
			cpuInfo.Model = value
		}

		if key == "cpu MHz" && cpuInfo.MHZ == "" {
			cpuInfo.MHZ = convertMhzToGhz(value)
		}

		// Stop early once we have both
		if cpuInfo.Model != "" && cpuInfo.MHZ != "" {
			break
		}
		cpuInfo.Model = strings.Trim(cpuInfo.Model, "Processor")
	}
}

func convertMhzToGhz(mhz string) string {
	floatVal, err := strconv.ParseFloat(mhz, 32)
	if err != nil {
		return ""
	}
	ghzValue := floatVal / 1000
	return fmt.Sprintf("%.2f", ghzValue)
}

func getLinuxCpuInfo() *Cpu {
	cpuData := getFile("/proc/cpuinfo")
	lines := strings.Split(string(cpuData), "\n")
	cpuInfo := &Cpu{}
	processCpuInfo(lines, cpuInfo)
	return cpuInfo
}
