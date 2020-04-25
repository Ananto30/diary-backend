package cmd

import (
	"github.com/golpo/config"
	"github.com/golpo/model"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrates the database",
	Run: func(cmd *cobra.Command, args []string) {
		if err := config.ConnectDB(); err != nil {
			log.Fatal(err)
		}
		config.DB.AutoMigrate(&model.User{}, &model.Diary{})
	},
}
