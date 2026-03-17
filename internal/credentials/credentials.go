package credentials

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	dirName  = ".lix"
	fileName = "credentials.json"
)

// Credentials holds the stored authentication data for the CLI.
type Credentials struct {
	Token   string `json:"token"`
	APIBase string `json:"api_base"`
	Email   string `json:"email"`
	SavedAt string `json:"saved_at"`
}

// Path returns the full path to the credentials file.
func Path() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine home directory: %w", err)
	}
	return filepath.Join(home, dirName, fileName), nil
}

// Save persists the credentials to disk with restrictive permissions.
func Save(creds *Credentials) error {
	path, err := Path()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return fmt.Errorf("could not create credentials directory: %w", err)
	}
	data, err := json.MarshalIndent(creds, "", "  ")
	if err != nil {
		return fmt.Errorf("could not marshal credentials: %w", err)
	}
	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("could not write credentials file: %w", err)
	}
	return nil
}

// Load reads the stored credentials from disk.
func Load() (*Credentials, error) {
	path, err := Path()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("not logged in — run 'lix-agents auth login' first: %w", err)
	}
	var creds Credentials
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, fmt.Errorf("corrupted credentials file: %w", err)
	}
	return &creds, nil
}

// New creates a new Credentials value with the current timestamp.
func New(token, apiBase, email string) *Credentials {
	return &Credentials{
		Token:   token,
		APIBase: apiBase,
		Email:   email,
		SavedAt: time.Now().UTC().Format(time.RFC3339),
	}
}
