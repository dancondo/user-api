package cmd

import (
	"sync"

	"github.com/dancondo/users-api/common"
	"github.com/dancondo/users-api/domain/user"
	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Create Seed Data",
	Long:  `Create seed data in the database`,
	RunE:  seedExecute,
}

func seedExecute(cmd *cobra.Command, args []string) error {
	var wg sync.WaitGroup

	userService := user.NewService()

	for _, u := range user.MockUsersRequest {
		wg.Add(1)
		go func(user *user.UserRequestDto) {
			defer wg.Done()
			_, err := userService.CreateUser(user)

			if err != nil {
				common.Log.Errorf("[SEED ERROR] %v", err.Error())
			}

		}(u)
	}

	wg.Wait()

	return nil
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
