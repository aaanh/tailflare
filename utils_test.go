package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/gofor-little/env"
)

func TestGetTailscaleHosts(t *testing.T) {
	client := &http.Client{}

	env.Load("./.env")
	apiKey := env.Get("TAILSCALE_API_KEY", "undefined")

	req, _ := http.NewRequest("GET", "https://api.tailscale.com/api/v2/tailnet/aaanh.github/devices", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	resp, _ := client.Do(req)

	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()
	t.Logf("%s", body)
}
