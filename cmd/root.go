package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "golang-clean-arc",
	Short: "Golang Clean Architecture",
	Long:  "Golang Clean Architecture is a project that implements the clean architecture pattern in Golang",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := os.Setenv("TZ", "UTC"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(serverCmd)
	rootCmd.Execute()
}
