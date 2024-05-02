package main

import (
	"fmt"

	cf "github.com/aaanh/tailflare/lib/cloudflare"
	structs "github.com/aaanh/tailflare/lib/structs"
	ts "github.com/aaanh/tailflare/lib/tailscale"
)

func updateStates(config *structs.Config) {
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

	if len(config.CloudflareZoneId) > 0 {
		config.States.CloudflareZoneIdExist = true
	} else {
		config.States.CloudflareZoneIdExist = false
	}
}

func exportEnvironmentVariables(config *structs.Config) {
	// To be implemented
	return
}

func viewCurrentConfigs(cfg *structs.Config) {
	fmt.Printf("\n\n")
	fmt.Println("=== Configurations and Keys ===")
	fmt.Println("- Tailscale API key: ", cfg.Keys.TailscaleApiKey)
	fmt.Println("- Tailnet Organization: ", cfg.TailnetOrg)
	fmt.Println("- Cloudflare API key: ", cfg.Keys.CloudflareApiKey)
	fmt.Println("- Cloudflare Zone ID (domain ID): ", cfg.CloudflareZoneId)
	fmt.Printf("\n\n")
}

// TODO: To be implemented
// func dryRun(config structs.Config) any {
// 	fmt.Printf("\n\nPerforming dry run and display what-if results.\n\n")
// 	fmt.Printf("\n\nCurrently under development.\n\n")
// 	return nil
// }

func performSync(config *structs.Config) {
	cf.CheckCloudflareToken(config)
	cf.AddCloudflareDnsRecords(config, ts.GetTailscaleDevices(config))
}
