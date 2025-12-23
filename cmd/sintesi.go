package cmd

import (
	"fmt"
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
}
