package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	structs "github.com/aaanh/tailflare/lib/structs"
	utils "github.com/aaanh/tailflare/utils"
)

func ConfigureCloudflare(config *structs.Config) {
	fmt.Println("Enter Cloudflare API key")
	fmt.Printf("How to get your Cloudflare API key: https://developers.cloudflare.com/fundamentals/api/get-started/create-token/\n\n")
	fmt.Print("> ")
	var apiKey string
	fmt.Scan(&apiKey)
	config.Keys.CloudflareApiKey = apiKey
	fmt.Printf("\n\n")
}

func ConfigureCloudflareZoneId(config *structs.Config) {
	fmt.Println("Enter Cloudflare Zone ID")
	fmt.Printf("Log in to the dashboard and select the target domain. Zone ID should be under the API section\n\n")
	fmt.Print("> ")
	var cfZoneId string
	fmt.Scan(&cfZoneId)
	config.CloudflareZoneId = cfZoneId
	fmt.Printf("\n\n")
}

func CheckCloudflareToken(config *structs.Config) {
	fmt.Println("> Checking Cloudflare token...")
	client := &http.Client{}

	endpoints := utils.GenerateEndpoints(config)
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

func ConstructCloudflareDnsPayload(device structs.Device, uri string) structs.Payload {
	deviceName := strings.Split(device.Name, ".")[0]

	payload := structs.Payload{
		Content: device.Addresses[0],
		Name:    deviceName + "." + uri,
		Proxied: false,
		Type:    "A",
		Ttl:     3600,
	}

	return payload
}

func AddCloudflareDnsRecords(config *structs.Config, devices structs.Devices) {
	fmt.Println("Enter your domain URI (e.g. devices.my-domain.com)")
	fmt.Printf("> ")
	var uri string
	fmt.Scanln(&uri)

	client := &http.Client{}

	endpoints := utils.GenerateEndpoints(config)

	fmt.Println("> NOW ADDING DEVICES...")

	for idx, device := range devices.Devices {
		fmt.Printf(">>> Adding device %d\n", idx)
		payload := ConstructCloudflareDnsPayload(device, uri)
		marshalled, err := json.Marshal(payload)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}

		req, _ := http.NewRequest("POST", endpoints.CloudflareAddRecord, bytes.NewReader(marshalled))

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "Bearer "+config.Keys.CloudflareApiKey)

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(">>> Error occurred with Cloudflare API request.")
		}

		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		fmt.Println(res)
		fmt.Println(string(body))
		fmt.Printf("\n\n")
	}
}

func GetDomainFromZoneId(config *structs.Config) string {
	client := &http.Client{}
	endpoints := utils.GenerateEndpoints(config)

	req, _ := http.NewRequest("GET", endpoints.CloudflareGetDomainFromZoneId, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+config.Keys.CloudflareApiKey)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(">>> Error occurred with Cloudflare API request.")
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var data struct {
		Result struct {
			Name string `json:"name"`
		} `json:"result"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %s", err)
		fmt.Println()
	}

	return data.Result.Name
}
