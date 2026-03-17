package cmd

import (
	"fmt"

	"github.com/lix-it/lix-agents/internal/credentials"
	"github.com/spf13/cobra"
)

func newStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show current login status",
		Long:  "Checks whether saved credentials exist and displays the logged-in email.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runStatus()
		},
	}
}

func runStatus() error {
	creds, err := credentials.Load()
	if err != nil {
		fmt.Println("Not logged in.")
		return nil
	}

	fmt.Println("Logged in.")
	if creds.Email != "" {
		fmt.Printf("Email: %s\n", creds.Email)
	}
	if creds.SavedAt != "" {
		fmt.Printf("Since: %s\n", creds.SavedAt)
	}
	return nil
}
