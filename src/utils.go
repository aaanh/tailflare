package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"net/http"
)

type Device struct {
	Addresses []string `json:"addresses"`
	Name      string   `json:"name"`
	Id        string   `json:"id"`
}

type Devices struct {
	Devices []Device `json:"devices"`
}
type Endpoints struct {
	TailscaleDevices    string
	CloudflareKeyCheck  string
	CloudflareAddRecord string
}

func generateEndpoints(cfg *Config) Endpoints {
	endpoints := Endpoints{
		fmt.Sprintf("https://api.tailscale.com/api/v2/tailnet/%s/devices", cfg.TailnetOrg),
		"https://api.cloudflare.com/client/v4/user/tokens/verify",
		fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", cfg.CloudflareZoneId),
	}

	return endpoints
}

func getTailscaleDevices(config *Config) Devices {
	fmt.Printf("\n> Step: Querying Tailscale devices\n\n")
	client := &http.Client{}

	// var devices Devices
	endpoints := generateEndpoints(config)

	req, _ := http.NewRequest("GET", endpoints.TailscaleDevices, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Keys.TailscaleApiKey))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Unable to perform getTailscaleDevices request")
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var data Devices
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %s", err)
		fmt.Println()
	}

	for i := 0; i < len(data.Devices); i++ {
		fmt.Print("Device ", i, ": ")
		fmt.Print(data.Devices[i].Addresses[0], " - ")
		fmt.Print(data.Devices[i].Name)
		fmt.Println()
	}
	fmt.Println()

	return data
}

func checkCloudflareToken(config *Config) {
	fmt.Println("> Checking Cloudflare token...")
	client := &http.Client{}

	endpoints := generateEndpoints(config)
	req, _ := http.NewRequest("GET", endpoints.CloudflareKeyCheck, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Keys.CloudflareApiKey))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Unable to perform Cloudflare API key check request")
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var data struct {
		Result struct {
			Status string
		}
		Success  bool
		Messages []struct {
			Message string
		}
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %s", err)
		fmt.Println()
	}

	fmt.Printf("\n\nCheck results:\n> key_status: %s\n> success: %t\n> msg: %s\n\n",
		data.Result.Status, data.Success, data.Messages[0].Message)
}

func constructCloudflareDnsPayload(device Device, uri string) Payload {

	deviceName := strings.Split(device.Name, ".")[0]

	payload := Payload{
		Content: device.Addresses[0],
		Name:    deviceName + "." + uri,
		Proxied: false,
		Type:    "A",
		Ttl:     3600,
	}

	return payload
}

func addCloudflareDnsRecords(config *Config, devices Devices) {
	fmt.Println("Enter your domain URI (e.g. devices.my-domain.com)")
	fmt.Printf("> ")
	var uri string
	fmt.Scanln(&uri)

	client := &http.Client{}

	endpoints := generateEndpoints(config)

	fmt.Println("> NOW ADDING DEVICES...")

	for idx, device := range devices.Devices {
		fmt.Printf(">> Adding device %d\n", idx)
		payload := constructCloudflareDnsPayload(device, uri)
		marshalled, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}

		req, _ := http.NewRequest("POST", endpoints.CloudflareAddRecord, bytes.NewReader(marshalled))

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+config.Keys.CloudflareApiKey)

		res, _ := client.Do(req)

		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		fmt.Println(res)
		fmt.Println(string(body))
		fmt.Println()
		fmt.Println()
	}

}
