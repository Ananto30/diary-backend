package cmd

import (
	"fmt"
	"github.com/golpo/server"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the Golpo app",
	Run: func(cmd *cobra.Command, args []string) {
		server.StartServer()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
