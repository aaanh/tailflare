package tailscale

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	structs "github.com/aaanh/tailflare/lib/structs"
	utils "github.com/aaanh/tailflare/utils"
)

func GetTailscaleDevices(config *structs.Config) structs.Devices {
	fmt.Printf("\n> Step: Querying Tailscale devices\n\n")
	client := &http.Client{}

	// var devices Devices
	endpoints := utils.GenerateEndpoints(config)

	req, _ := http.NewRequest("GET", endpoints.TailscaleDevices, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Keys.TailscaleApiKey))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Unable to perform getTailscaleDevices request")
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	var data structs.Devices
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

func ConfigureTailscale(config *structs.Config) {
	fmt.Println("Enter Tailscale API key")
	fmt.Printf("Get your Tailscale API key: https://login.tailscale.com/login?next_url=%%2Fadmin%%2Fsettings%%2Fkeys\n\n")
	fmt.Print("> ")
	var apiKey string
	fmt.Scan(&apiKey)
	config.Keys.TailscaleApiKey = apiKey
	fmt.Printf("\n\n")
}

func ConfigureTailnetOrg(config *structs.Config) {
	fmt.Println("Enter Tailnet organization")
	fmt.Printf("Should be under the Organization section at https://login.tailscale.com/admin/settings/general\n\n")
	fmt.Print("> ")
	var tailnetOrg string
	fmt.Scan(&tailnetOrg)
	config.TailnetOrg = tailnetOrg
	fmt.Printf("\n\n")
}
