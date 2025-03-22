package database

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// SupabaseClient wraps the Supabase connection and operations
type SupabaseClient struct {
	client  *http.Client
	baseURL string
	apiKey  string
}

// NewSupabaseClient creates a new Supabase client
func NewSupabaseClient() (*SupabaseClient, error) {
	baseURL := os.Getenv("SUPABASE_URL")
	apiKey := os.Getenv("SUPABASE_KEY")

	if baseURL == "" || apiKey == "" {
		return nil, fmt.Errorf("SUPABASE_URL and SUPABASE_KEY environment variables are required")
	}

	return &SupabaseClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseURL: baseURL,
		apiKey:  apiKey,
	}, nil
}

// Query executes a Supabase query
func (c *SupabaseClient) Query(ctx context.Context, table string, query map[string]interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s/rest/v1/%s", c.baseURL, table)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Add headers
	req.Header.Set("apikey", c.apiKey)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=minimal")

	// Add query parameters
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, fmt.Sprintf("%v", value))
	}
	req.URL.RawQuery = q.Encode()

	// Execute request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read response body
	var result []byte
	if _, err := resp.Body.Read(result); err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return result, nil
}

// Insert inserts a record into a Supabase table
func (c *SupabaseClient) Insert(ctx context.Context, table string, data interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s/rest/v1/%s", c.baseURL, table)

	_, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("error marshaling data: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Add headers
	req.Header.Set("apikey", c.apiKey)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Prefer", "return=minimal")

	// Execute request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read response body
	var result []byte
	if _, err := resp.Body.Read(result); err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return result, nil
}
