package sysinfo

import (
	"os"
	"os/user"
)

// returns user + host
func Title() string {
	currUser, err := user.Current()
	user := "Unknown"
	if err == nil {
		user = currUser.Username
	}

	hostname, err := os.Hostname()
	host := "Unknown"
	if err == nil {
		host = hostname
	}

	return user + "@" + host
}
