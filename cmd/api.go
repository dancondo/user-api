package cmd

import (
	"github.com/dancondo/users-api/api"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Initialize api",
	Long:  `Initialize api`,
	RunE:  apiExecute,
}

func apiExecute(cmd *cobra.Command, args []string) error {
	return api.StartHTTP()
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
