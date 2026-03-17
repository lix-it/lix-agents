package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/lix-it/lix-agents/internal/credentials"
	"github.com/spf13/cobra"
)

const (
	pollInterval = 3 * time.Second
	pollTimeout  = 30 * time.Minute
)

func newTokenCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "token",
		Short: "Request a temporary API token",
		Long:  "Requests a temporary API token. An approval email will be sent to the account owner.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runToken(APIBase)
		},
	}
}

// tokenRequestResponse is the JSON response from the token-request endpoint.
type tokenRequestResponse struct {
	RequestToken string `json:"request_token"`
	Message      string `json:"message"`
}

// tokenStatusResponse is the JSON response from the status polling endpoint.
type tokenStatusResponse struct {
	Status   string `json:"status"`
	APIToken string `json:"api_token"`
	Message  string `json:"message"`
}

// apiErrorResponse is the standard API error envelope.
type apiErrorResponse struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

// runToken creates a token request and polls until it is approved or denied.
func runToken(apiBase string) error {
	creds, err := credentials.Load()
	if err != nil {
		return err
	}
	if apiBase != creds.APIBase {
		creds.APIBase = apiBase
	}

	fmt.Println("Requesting temporary API token...")

	reqResp, err := createTokenRequest(creds)
	if err != nil {
		return err
	}

	fmt.Println("\n" + reqResp.Message)
	fmt.Println("Waiting for approval...")

	return pollForApproval(creds, reqResp.RequestToken)
}

// createTokenRequest calls the API to initiate a new token request and send
// the approval email.
func createTokenRequest(creds *credentials.Credentials) (*tokenRequestResponse, error) {
	reqBody := strings.NewReader(`{}`)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/cli/token-request", creds.APIBase), reqBody)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Set("Authorization", creds.Token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not reach API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var errResp apiErrorResponse
		_ = json.NewDecoder(resp.Body).Decode(&errResp)
		msg := errResp.Error.Message
		if msg == "" {
			msg = fmt.Sprintf("HTTP %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("token request failed: %s", msg)
	}

	var result tokenRequestResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("could not parse response: %w", err)
	}
	return &result, nil
}

// pollForApproval repeatedly checks the token request status until it is
// approved, denied, or the timeout is reached.
func pollForApproval(creds *credentials.Credentials, requestToken string) error {
	client := &http.Client{Timeout: 15 * time.Second}
	deadline := time.Now().Add(pollTimeout)

	for time.Now().Before(deadline) {
		time.Sleep(pollInterval)

		statusReq, err := http.NewRequest("GET",
			fmt.Sprintf("%s/cli/token-request/%s/status", creds.APIBase, requestToken), nil)
		if err != nil {
			continue
		}
		statusReq.Header.Set("Authorization", creds.Token)

		resp, err := client.Do(statusReq)
		if err != nil {
			continue
		}

		var status tokenStatusResponse
		_ = json.NewDecoder(resp.Body).Decode(&status)
		resp.Body.Close()

		switch status.Status {
		case "approved":
			fmt.Println("\nToken approved!")
			fmt.Printf("\nTemporary API Token: %s\n", status.APIToken)
			if status.Message != "" {
				fmt.Println(status.Message)
			}
			return nil
		case "denied":
			return fmt.Errorf("token request was denied")
		case "expired":
			return fmt.Errorf("token request expired")
		case "pending":
			fmt.Print(".")
		}
	}

	return fmt.Errorf("timed out waiting for approval after %v", pollTimeout)
}
