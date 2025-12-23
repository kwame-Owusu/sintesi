package cmd

import (
	"fmt"
	"github/kwame-Owusu/sintesi/internal/sysinfo"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "sintesi",
	Short: "Simple system info fetch utility",
	Long:  "System info fetch utility, showcases important info about a system",
	Run: func(cmd *cobra.Command, args []string) {
		sintesi()
	},
}

func sintesi() {
	fmt.Println("printing system info...")
	fmt.Printf("Title: %s\n", sysinfo.Title())
	fmt.Printf("OS: %s\n", sysinfo.OS())
	fmt.Printf("Kernel: %s\n", sysinfo.Kernel())
	fmt.Printf("Shell: %s\n", sysinfo.Shell())
	hardware := sysinfo.LinuxHardware()
	fmt.Printf("CPU: %s @ %sGHZ\n", hardware.Cpu.Model, hardware.Cpu.MHZ)
	uptime, err := sysinfo.Uptime()
	if err != nil {
		fmt.Printf("Error getting uptime through syscall: %v", err)
	}
	fmt.Printf("Uptime: %d hours, %d mins\n", int(uptime.Hours())%24, int(uptime.Minutes())%60)
	fmt.Printf("Ram: %s\n", hardware.Ram.Total)
}
