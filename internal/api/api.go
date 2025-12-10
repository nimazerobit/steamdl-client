package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"steam-lancache/internal/config"
	"time"
)

type SubscriptionDetails struct {
	SubscriptionID int    `json:"subscription_id"`
	End            string `json:"end"`
	UserIP         string `json:"user_ip"`
	Status         string `json:"status"`
}

func GetTokenInfo(token string) (string, SubscriptionDetails, error) {
	url := fmt.Sprintf(config.TokenAPI, token)
	resp, err := http.Get(url)
	if err != nil {
		return "", SubscriptionDetails{}, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", SubscriptionDetails{}, fmt.Errorf("invalid token (status: %d)", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", SubscriptionDetails{}, fmt.Errorf("failed to read API response body: %w", err)
	}

	var details SubscriptionDetails
	if err := json.Unmarshal(bodyBytes, &details); err != nil {
		return "", SubscriptionDetails{}, fmt.Errorf("failed to parse API JSON response: %w", err)
	}

	// Subscription Validation
	if details.Status != "active" {
		return "", details, fmt.Errorf("subscription is not active (status: %s)", details.Status)
	}

	const layout = "2006-01-02 15:04:05"
	endDate, err := time.Parse(layout, details.End)
	if err != nil {
		return "", details, fmt.Errorf("could not parse subscription end date '%s': %w", details.End, err)
	}
	if endDate.Before(time.Now()) {
		return "", details, fmt.Errorf("subscription has expired on %s", details.End)
	}

	// get server ip from response header
	serverIP := resp.Header.Get("x-server-ip")
	if serverIP == "" {
		return "", details, fmt.Errorf("x-server-ip header not found despite valid response")
	}

	return serverIP, details, nil
}
