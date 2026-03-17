package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/lix-it/lix-agents/internal/credentials"
	"github.com/spf13/cobra"
)

const (
	loginPollInterval = 3 * time.Second
	loginPollTimeout  = 5 * time.Minute
)

func newLoginCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Log in to the Lix API",
		Long:  "Initiates a login session and provides a URL for the user to authenticate in a browser on any device.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runLogin(APIBase)
		},
	}
}

// loginInitResponse is the JSON response from the login init endpoint.
type loginInitResponse struct {
	Code     string `json:"code"`
	LoginURL string `json:"login_url"`
	Message  string `json:"message"`
}

// loginStatusResponse is the JSON response from the login status endpoint.
type loginStatusResponse struct {
	Status string `json:"status"`
	Token  string `json:"token"`
	Email  string `json:"email"`
}

// runLogin initiates a device-flow login: requests a session code from the
// server, displays the login URL for the user, then polls until the user
// completes authentication.
func runLogin(apiBase string) error {
	fmt.Println("Requesting login session...")

	client := &http.Client{Timeout: 30 * time.Second}

	// Step 1: Init the login session
	resp, err := client.Post(fmt.Sprintf("%s/cli/login/init", apiBase), "application/json", nil)
	if err != nil {
		return fmt.Errorf("could not reach server: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("login init failed: HTTP %d", resp.StatusCode)
	}

	var initResp loginInitResponse
	if err := json.NewDecoder(resp.Body).Decode(&initResp); err != nil {
		return fmt.Errorf("could not parse response: %w", err)
	}

	fmt.Printf("\nPlease visit this URL to log in:\n\n  %s\n\n", initResp.LoginURL)
	fmt.Println("Waiting for login...")

	// Step 2: Poll for completion
	return pollForLogin(client, apiBase, initResp.Code)
}

// pollForLogin polls the login status endpoint until the user completes
// authentication or the timeout is reached.
func pollForLogin(client *http.Client, apiBase, code string) error {
	deadline := time.Now().Add(loginPollTimeout)

	for time.Now().Before(deadline) {
		time.Sleep(loginPollInterval)

		resp, err := client.Get(fmt.Sprintf("%s/cli/login/%s/status", apiBase, code))
		if err != nil {
			continue
		}

		var status loginStatusResponse
		_ = json.NewDecoder(resp.Body).Decode(&status)
		resp.Body.Close()

		switch status.Status {
		case "authenticated":
			creds := credentials.New(status.Token, apiBase, status.Email)
			if err := credentials.Save(creds); err != nil {
				return fmt.Errorf("login successful but could not save credentials: %w", err)
			}
			fmt.Printf("\nLogin successful! Credentials saved.\n")
			if status.Email != "" {
				fmt.Printf("Logged in as: %s\n", status.Email)
			}
			return nil
		case "expired":
			return fmt.Errorf("login session expired — please try again")
		case "pending":
			fmt.Print(".")
		}
	}

	return fmt.Errorf("login timed out after %v", loginPollTimeout)
}
