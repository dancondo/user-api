package cmd

import (
	"os"

	"github.com/dancondo/users-api/common"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "users-api",
	Short: "users-api",
	Long:  "users-api",
}

func Execute() {
	common.LoadEnv()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
