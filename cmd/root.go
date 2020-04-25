package cmd

import (
	"fmt"
	"github.com/golpo/config"
	"github.com/golpo/server"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the Golpo app",
	Run: func(cmd *cobra.Command, args []string) {

		if err := config.ConnectDB(); err != nil {
			log.Fatal(err)
		}

		if err := server.StartServer(); err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
