// License: MIT Open Source

package main

import (
	"fmt"
	"os"

	"github.com/gofor-little/env"
	_ "golang.org/x/term"

	cf "github.com/aaanh/tailflare/lib/cloudflare"
	structs "github.com/aaanh/tailflare/lib/structs"
	ts "github.com/aaanh/tailflare/lib/tailscale"
)

func displayHeader(version string) {
	fmt.Printf(">>> Tailflare v%s <<<\n", version)
	fmt.Println("Sync your Tailflare hosts with Cloudflare DNS")
	fmt.Println()
}

func choiceHandler(choice int, config *structs.Config, added *structs.Devices) {
	switch choice {
	case 1:
		ts.ConfigureTailscale(config)
	case 2:
		ts.ConfigureTailnetOrg(config)
	case 3:
		cf.ConfigureCloudflare(config)
	case 4:
		cf.ConfigureCloudflareZoneId(config)
	case 5:
		performSync(config, added)
	case 6:
		cf.DeleteAddedDnsRecords(config, added)
	case 7:
		cf.DeleteAllTailscaleRecords(config)
	case 8:
		viewCurrentConfigs(config)
	case 9:
		{
			fmt.Println("\n=== Thanks for using Tailflare :> ===")
			os.Exit(0)
		}
	default:
		{
			fmt.Printf("\n\n> Invalid choice :<\n\n")
		}
	}
}

func program(config *structs.Config, added *structs.Devices) {
	for {
		updateStates(config)
		displayHeader(config.Version)
		choice := Menu(config)
		choiceHandler(choice, config, added)
	}
}

func main() {
	states := structs.States{
		TailscaleKeyExist:     false,
		TailnetOrgExist:       false,
		CloudflareKeyExist:    false,
		CloudflareZoneIdExist: false}

	var version string
	var tailscaleApiKey string
	var cloudflareApiKey string
	var tailnetOrg string
	var cloudflareZoneId string

	var added structs.Devices

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic occurred during environment variable load.")
			fmt.Println(err)

			version := "0.0.0-undefined"
			tailscaleApiKey := ""
			tailnetOrg := ""
			cloudflareApiKey := ""
			cloudflareZoneId := ""

			config := structs.Config{
				Version:          version,
				States:           states,
				TailnetOrg:       tailnetOrg,
				CloudflareZoneId: cloudflareZoneId,
				Keys: structs.Keys{
					TailscaleApiKey:  tailscaleApiKey,
					CloudflareApiKey: cloudflareApiKey,
				},
			}

			program(&config, &added)
		}
	}()
	env.Load("./config.cfg", "./.env")

	version = env.Get("version", "0.0.0-undefined")
	tailscaleApiKey = env.Get("TAILSCALE_API_KEY", "")
	tailnetOrg = env.Get("TAILNET_ORG", "")
	cloudflareApiKey = env.Get("CLOUDFLARE_API_KEY", "")
	cloudflareZoneId = env.Get("CLOUDFLARE_ZONE_ID", "")

	config := structs.Config{
		Version:          version,
		States:           states,
		TailnetOrg:       tailnetOrg,
		CloudflareZoneId: cloudflareZoneId,
		Keys: structs.Keys{
			TailscaleApiKey:  tailscaleApiKey,
			CloudflareApiKey: cloudflareApiKey,
		},
	}

	program(&config, &added)

}
