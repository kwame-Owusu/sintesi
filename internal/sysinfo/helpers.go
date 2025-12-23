package sysinfo

import (
	"fmt"
	"os"
	"strings"
)

func getFile() []byte {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		fmt.Printf("Error occurred reading os-release, %v", err)
	}
	return data
}

// parses the os-release file and returns PRETTY_NAME value
func getLinuxOSName() string {
	osFile := getFile()
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
