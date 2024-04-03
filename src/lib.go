package main

import (
	"fmt"
	"sort"
)

func updateStates(config *Config) {
	if len(config.Keys.CloudflareApiKey) > 0 {
		config.States.CloudflareKeyExist = true
	} else {
		config.States.CloudflareKeyExist = false
	}

	if len(config.Keys.TailscaleApiKey) > 0 {
		config.States.TailscaleKeyExist = true
	} else {
		config.States.TailscaleKeyExist = false
	}

	if len(config.TailnetOrg) > 0 {
		config.States.TailnetOrgExist = true
	} else {
		config.States.TailnetOrgExist = false
	}
}

func configureTailscale(config *Config) {
	fmt.Println("Enter Tailscale API key")
	fmt.Printf("Get your Tailscale API key: https://login.tailscale.com/login?next_url=%%2Fadmin%%2Fsettings%%2Fkeys\n\n")
	fmt.Print("> ")
	var apiKey string
	fmt.Scan(&apiKey)
	config.Keys.TailscaleApiKey = apiKey
	fmt.Printf("\n\n")
}

func configureCloudflare(config *Config) {
	fmt.Println("Enter Cloudflare API key")
	fmt.Printf("How to get your Cloudflare API key: https://developers.cloudflare.com/fundamentals/api/get-started/create-token/\n\n")
	fmt.Print("> ")
	var apiKey string
	fmt.Scan(&apiKey)
	config.Keys.CloudflareApiKey = apiKey
	fmt.Printf("\n\n")
}

func configureTailnetOrg(config *Config) {
	fmt.Println("Enter Tailnet organization")
	fmt.Printf("Should be under the Organization section at https://login.tailscale.com/admin/settings/general\n\n")
	fmt.Print("> ")
	var tailnetOrg string
	fmt.Scan(&tailnetOrg)
	config.TailnetOrg = tailnetOrg
	fmt.Printf("\n\n")
}

func menu(cfg Config) int {
	menuOptions := map[int]string{
		1: fmt.Sprintf("Configure Tailscale API key (%s)", func() string {
			if cfg.States.TailscaleKeyExist {
				return "Added"
			} else {
				return "Not added"
			}
		}()),
		2: fmt.Sprintf("Configure Cloudflare API key (%s)", func() string {
			if cfg.States.CloudflareKeyExist {
				return "Added"
			} else {
				return "Not added"
			}
		}()),
		3: fmt.Sprintf("Configure Tailnet Organization (%s)", func() string {
			if cfg.States.TailnetOrgExist {
				return cfg.TailnetOrg
			} else {
				return "Not added"
			}
		}()),
		4: "Dry run (What-If)",
		5: "Perform Sync",
		6: "Exit",
	}

	// Solve the misordered printing by sorting the keys in the map
	// ensuring that the function execution time doesn't affect
	// the ordering 🤷
	var keys []int
	for k := range menuOptions {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	// finish correcting sort

	for _, option := range keys {
		fmt.Printf("%d. %s\n", option, menuOptions[option])
	}

	fmt.Printf("\n\nChoose an option: ")
	var choice int
	fmt.Scan(&choice)
	return choice
}
