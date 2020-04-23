package cmd

import (
	"github.com/golpo/db"
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
		if err := db.Connect(); err != nil {
			log.Fatal(err)
		}
		db.DB.AutoMigrate(&model.User{})
	},
}
