package main

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"

	"net/http"
	"strings"
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
		// Regular expression to match digits
		digitRegex := regexp.MustCompile("[0-9]")

		// Replace all digits with "*"
		maskedIPv4Address := digitRegex.ReplaceAllStringFunc(data.Devices[i].Addresses[0], func(s string) string {
			return "*"
		})

		fmt.Print("Device ", i, ": ")
		fmt.Print(maskedIPv4Address, " - ")
		fmt.Print(strings.Replace(data.Devices[i].Name, "bobcat-atlas", "*******", -1))
		fmt.Println()
	}
	fmt.Println()
}
