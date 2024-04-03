// License: MIT Open Source

package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/gofor-little/env"
	_ "golang.org/x/term"
)

type Keys struct {
	TailscaleApiKey  string
	CloudflareApiKey string
}

type States struct {
	TailscaleKeyExist  bool
	CloudflareKeyExist bool
}

type Config struct {
	Version    string
	States     States
	Keys       Keys
	TailnetOrg string
}

func displayHeader(version string) {
	fmt.Printf(">>> Tailflare v%s <<<\n", version)
	fmt.Println("Sync your Tailflare hosts with Cloudflare DNS")
	fmt.Println()
}

func menu(states States) int {
	menuOptions := map[int]string{
		1: fmt.Sprintf("Configure Tailscale API key (%s)", func() string {
			if states.TailscaleKeyExist {
				return "Added"
			} else {
				return "Not added"
			}
		}()),
		2: fmt.Sprintf("Configure Cloudflare API key (%s)", func() string {
			if states.CloudflareKeyExist {
				return "Added"
			} else {
				return "Not added"
			}
		}()),
		3: "Dry run (What-If)",
		4: "Perform Sync",
		5: "Exit",
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

func dryRun(config Config) any {
	fmt.Printf("\n\nPerforming dry run and display what-if results.\n\n")
	return nil
}

func runSync(config *Config) {
	getTailscaleDevices(config)
}

func choiceHandler(choice int, config *Config) {
	switch choice {
	case 1:
		configureTailscale(config)
	case 2:
		configureCloudflare(config)
	case 3:
		dryRun(*config)
	case 4:
		runSync(config)
	case 5:
		{
			fmt.Println("\n=== Thanks for using Tailflare :) ===")
			os.Exit(0)
		}
	default:
		{
			fmt.Printf("\n\n> Invalid choice :(\n\n")
		}
	}
}

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
}

func program(config *Config) {
	for {
		updateStates(config)
		displayHeader(config.Version)
		choice := menu(config.States)
		choiceHandler(choice, config)
	}
}

func main() {
	states := States{false, false}

	env.Load("./config.cfg", "./.env")
	version := env.Get("version", "0.0.0-undefined")
	tailscaleApiKey := env.Get("TAILSCALE_API_KEY", "")
	cloudflareApiKey := env.Get("CLOUDFLARE_API_KEY", "")

	config := Config{version, states, Keys{
		tailscaleApiKey, cloudflareApiKey,
	}, env.Get("TAILNET_ORG", "undefined")}

	program(&config)

}
