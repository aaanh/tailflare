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
		configureTailnetOrg(config)
	case 4:
		dryRun(*config)
	case 5:
		runSync(config)
	case 6:
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
	states := States{false, false, false}

	env.Load("./config.cfg", "./.env")
	version := env.Get("version", "0.0.0-undefined")
	tailscaleApiKey := env.Get("TAILSCALE_API_KEY", "")
	cloudflareApiKey := env.Get("CLOUDFLARE_API_KEY", "")
	tailnetOrg := env.Get("TAILNET_ORG", "")

	config := Config{version, states,
		Keys{tailscaleApiKey, cloudflareApiKey},
		tailnetOrg}

	program(&config)

}
