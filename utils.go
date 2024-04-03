package main

import (
	"fmt"
	"io"
	"net/http"
)

type Device struct {
	addresses struct {
		ipv4 string
		ipv6 string
	}
	name string
	id   string
}

type Devices []Device

type Endpoints struct {
	Devices string
}

func generateEndpoints(tailnetOrg string) Endpoints {
	endpoints := Endpoints{
		fmt.Sprintf("https://api.tailscale.com/api/v2/tailnet/%s/devices", tailnetOrg),
	}

	return endpoints
}

func getTailscaleDevices(config *Config) {
	client := &http.Client{}

	// var devices Devices
	endpoints := generateEndpoints(config.TailnetOrg)

	req, _ := http.NewRequest("GET", endpoints.Devices, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Keys.TailscaleApiKey))
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("No response")
	}
	fmt.Println(string(body))
}
