package main

import (
	"encoding/json"
	"fmt"
	"io"

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
	Devices string
}

func generateEndpoints(tailnetOrg string) Endpoints {
	endpoints := Endpoints{
		fmt.Sprintf("https://api.tailscale.com/api/v2/tailnet/%s/devices", tailnetOrg),
	}

	return endpoints
}

func getTailscaleDevices(config *Config) {
	fmt.Printf("\n> Step: Querying Tailscale devices\n\n")
	client := &http.Client{}

	// var devices Devices
	endpoints := generateEndpoints(config.TailnetOrg)

	req, _ := http.NewRequest("GET", endpoints.Devices, nil)
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
}
