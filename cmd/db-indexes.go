package cmd

import (
	userRepository "github.com/dancondo/users-api/repository/user-repository"
	"github.com/spf13/cobra"
)

var dbIndexCmd = &cobra.Command{
	Use:   "db-indexes",
	Short: "Initialize Mongo DB indexes",
	Long:  `Initialize indexes for mongo DB`,
	RunE:  dbIndexExecute,
}

func dbIndexExecute(cmd *cobra.Command, args []string) error {
	userRepository := userRepository.New()
	return userRepository.CreateIndex("username")
}

func init() {
	rootCmd.AddCommand(dbIndexCmd)
}
