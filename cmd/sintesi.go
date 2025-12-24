package cmd

import (
	"fmt"
	"github/kwame-Owusu/sintesi/internal/sysinfo"

	"github.com/spf13/cobra"
)

const (
	reset     = "\033[0m"
	red       = "\033[31m"
	green     = "\033[32m"
	yellow    = "\033[33m"
	blue      = "\033[34m"
	magenta   = "\033[35m"
	cyan      = "\033[36m"
	lightPink = "\033[38;5;218m" // soft pastel pink
)

var cat = []string{
	"      ／＞　 フ",
	"     | 　^　^|  ",
	"    ／` ミ＿xノ",
	"   /　　　　 | ",
	"  /　 ヽ　　 ﾉ ",
	" │　　|　|　| ",
	" ／￣  |　|　| ",
	" (￣ヽ＿_ヽ_)__)",
	"＼二)",
}

var RootCmd = &cobra.Command{
	Use:   "sintesi",
	Short: "Simple system info fetch utility",
	Long:  "System info fetch utility, showcases important info about a system",
	Run: func(cmd *cobra.Command, args []string) {
		sintesi()
	},
}

func sintesi() {
	info := infoLines()

	maxLines := len(cat)
	maxLines = max(maxLines, len(info))

	fmt.Println()
	fmt.Println()
	for i := 0; i < maxLines; i++ {
		// Left (bunny)
		if i < len(cat) {
			fmt.Printf("%-12s", cat[i])
		} else {
			fmt.Printf("%-12s", "")
		}

		// Right (info)
		if i < len(info) {
			fmt.Println("  " + info[i])
		} else {
			fmt.Println()
		}
	}

	// Color palette preview
	fmt.Printf(reset + "\n                  ")
	for i := range 8 {
		fmt.Printf("\033[10%dm   ", i)
	}
	fmt.Println(reset)
}

func infoLines() []string {
	hardware := sysinfo.GetHardware()

	lines := []string{
		lightPink + sysinfo.Title() + reset,
		green + "OS:      " + reset + sysinfo.OS(),
		yellow + "Kernel:  " + reset + sysinfo.Kernel(),
		blue + "Shell:   " + reset + sysinfo.Shell(),
		magenta + "CPU:     " + reset + sysinfo.FormatCPU(hardware.Cpu),
		green + "Memory:  " + reset +
			sysinfo.FormatMemory(hardware.Ram.Total, hardware.Ram.Available),
		red + "Terminal:" + reset + " " + sysinfo.Terminal(),
	}

	uptime, err := sysinfo.Uptime()
	if err == nil {
		lines = append(
			lines,
			cyan+"Uptime:  "+reset+
				fmt.Sprintf("%d hours, %d mins",
					int(uptime.Hours())%24,
					int(uptime.Minutes())%60),
		)
	}

	return lines
}
