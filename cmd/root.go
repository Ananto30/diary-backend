package cmd

import (
	"fmt"
	"github.com/golpo/config"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the Golpo app",
	Run: func(cmd *cobra.Command, args []string) {
		config.StartServer()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
