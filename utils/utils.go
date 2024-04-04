package utils

import (
	"fmt"
	"sort"

	structs "github.com/aaanh/tailflare/lib/structs"
)

func GenerateEndpoints(cfg *structs.Config) structs.Endpoints {
	endpoints := structs.Endpoints{
		TailscaleDevices:              fmt.Sprintf("https://api.tailscale.com/api/v2/tailnet/%s/devices", cfg.TailnetOrg),
		CloudflareKeyCheck:            "https://api.cloudflare.com/client/v4/user/tokens/verify",
		CloudflareAddRecord:           fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", cfg.CloudflareZoneId),
		CloudflareGetDomainFromZoneId: fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s", cfg.CloudflareZoneId),
	}
	return endpoints
}

func Menu(cfg structs.Config) int {
	menuOptions := map[int]string{
		1: fmt.Sprintf("Configure Tailscale API key (%s)", func() string {
			if cfg.States.TailscaleKeyExist {
				return "Added"
			} else {
				return "Not added"
			}
		}()),
		2: fmt.Sprintf("Configure Tailnet Organization (%s)", func() string {
			if cfg.States.TailnetOrgExist {
				return cfg.TailnetOrg
			} else {
				return "Not added"
			}
		}()),
		3: fmt.Sprintf("Configure Cloudflare API key (%s)", func() string {
			if cfg.States.CloudflareKeyExist {
				return "Added"
			} else {
				return "Not added"
			}
		}()),
		4: fmt.Sprintf("Configure Cloudflare Zone ID (%s - %s)", structs., func() string {
			if cfg.States.CloudflareZoneIdExist {
				return cfg.CloudflareZoneId
			} else {
				return "Not added"
			}
		}()),
		5: "Dry run (What-If)",
		6: "Perform Sync",
		7: "Exit",
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
