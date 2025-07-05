package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func NotifyDiscord(message string) error {
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")

	if webhookURL == "" {
		return nil // Not configured
	}
	payload := fmt.Sprintf(`{"content":"%s"}`, message)
	_, err := http.Post(webhookURL, "application/json", bytes.NewBuffer([]byte(payload)))
	return err
}
