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

func getLinuxCpuInfo() *Cpu {
	data, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		return &Cpu{}
	}

	cpu := &Cpu{}
	lines := strings.Split(string(data), "\n")

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

		if key == "model name" && cpu.Model == "" {
			cpu.Model = value
		}

		if key == "cpu MHz" && cpu.MHZ == 0 {
			f, err := strconv.ParseFloat(value, 64)
			if err == nil {
				cpu.MHZ = f
			}
		}

		if cpu.Model != "" && cpu.MHZ != 0 {
			break
		}
		cpu.Model = strings.Trim(cpu.Model, "Processor")
	}

	return cpu
}

func getLinuxRamInfo() *Ram {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return &Ram{}
	}

	ram := &Ram{}
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Fields(line) // splits by any whitespace
		if len(parts) < 2 {
			continue
		}

		key := strings.TrimSuffix(parts[0], ":")
		valueStr := parts[1]

		val, err := strconv.ParseUint(valueStr, 10, 64)
		if err != nil {
			continue
		}

		switch key {
		case "MemTotal":
			ram.Total = val
		case "MemAvailable":
			ram.Available = val
		}

		if ram.Total != 0 && ram.Available != 0 {
			break
		}
	}

	return ram
}

func FormatMemory(totalKB, availableKB uint64) string {
	usedKB := totalKB - availableKB
	return fmt.Sprintf("%.2f GiB / %.2f GiB",
		float64(usedKB)/1024/1024,
		float64(totalKB)/1024/1024)
}

func FormatCPU(cpu *Cpu) string {
	return fmt.Sprintf("%s @ %.2f GHz", cpu.Model, cpu.MHZ/1000)
}

func GetHardware() Hardware {
	return LinuxHardware()
}
