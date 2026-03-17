package cmd

import "github.com/spf13/cobra"

// APIBase is the base URL for the Lix API, set via persistent flag.
var APIBase string

// rootCmd is the base command for the CLI.
var rootCmd = &cobra.Command{
	Use:   "lix-agents",
	Short: "Lix CLI authentication tool for issuing temporary API tokens",
	Long:  "A CLI tool that allows users to authenticate with the Lix API and issue temporary tokens with email-based approval.",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&APIBase, "api-base", "https://lix-it.com", "Base URL of the Lix web app")

	authCmd := &cobra.Command{
		Use:   "auth",
		Short: "Authentication commands",
	}
	authCmd.AddCommand(newLoginCmd(), newTokenCmd(), newStatusCmd())
	rootCmd.AddCommand(authCmd)
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}
