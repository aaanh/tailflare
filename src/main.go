// License: MIT Open Source

package main

import (
	"fmt"
	"os"

	"github.com/gofor-little/env"
	_ "golang.org/x/term"
)

func displayHeader(version string) {
	fmt.Printf(">>> Tailflare v%s <<<\n", version)
	fmt.Println("Sync your Tailflare hosts with Cloudflare DNS")
	fmt.Println()
}

func choiceHandler(choice int, config *Config) {
	switch choice {
	case 1:
		configureTailscale(config)
	case 2:
		configureTailnetOrg(config)
	case 3:
		configureCloudflare(config)
	case 4:
		configureCloudflareZoneId(config)
	case 5:
		dryRun(*config)
	case 6:
		performSync(config)
	case 7:
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

func program(config *Config) {
	for {
		updateStates(config)
		displayHeader(config.Version)
		choice := menu(*config)
		choiceHandler(choice, config)
	}
}

func main() {
	states := States{false, false, false, false}

	var version string
	var tailscaleApiKey string
	var cloudflareApiKey string
	var tailnetOrg string
	var cloudflareZoneId string

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic occurred during environment variable load.")
			fmt.Println(err)

			version := "0.0.0-undefined"
			tailscaleApiKey := ""
			tailnetOrg := ""
			cloudflareApiKey := ""
			cloudflareZoneId := ""

			config := Config{version, states,
				Keys{tailscaleApiKey, cloudflareApiKey},
				tailnetOrg, cloudflareZoneId}

			program(&config)
		}
	}()
	env.Load("./config.cfg", "./.env")

	version = env.Get("version", "0.0.0-undefined")
	tailscaleApiKey = env.Get("TAILSCALE_API_KEY", "")
	tailnetOrg = env.Get("TAILNET_ORG", "")
	cloudflareApiKey = env.Get("CLOUDFLARE_API_KEY", "")
	cloudflareZoneId = env.Get("CLOUDFLARE_ZONE_ID", "")

	config := Config{version, states,
		Keys{tailscaleApiKey, cloudflareApiKey},
		tailnetOrg, cloudflareZoneId}

	program(&config)

}
