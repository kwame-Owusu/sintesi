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
}
