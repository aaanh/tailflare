package main

import (
	"fmt"
	"sort"

	cf "github.com/aaanh/tailflare/lib/cloudflare"
	structs "github.com/aaanh/tailflare/lib/structs"
)

func Menu(cfg *structs.Config) int {
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
		4: fmt.Sprintf("Configure Cloudflare Zone ID (%s - %s)", cf.GetDomainFromZoneId(cfg), func() string {
			if cfg.States.CloudflareZoneIdExist {
				return cfg.CloudflareZoneId
			} else {
				return "Not added"
			}
		}()),
		5: "Perform Sync",
		6: "Delete synced records",
		7: "View configurations and keys",
		8: "Exit",
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
