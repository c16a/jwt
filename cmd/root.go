package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use: "jwt",
	Short: "jwt is a tiny JWT cli",
	Long: `jwt is a tool to quickly validate and generate JWT on the command line`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}