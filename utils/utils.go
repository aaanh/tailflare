package utils

import (
	"fmt"

	structs "github.com/aaanh/tailflare/lib/structs"
)

func GenerateEndpoints(cfg *structs.Config) structs.Endpoints {
	endpoints := structs.Endpoints{
		TailscaleDevices:               fmt.Sprintf("https://api.tailscale.com/api/v2/tailnet/%s/devices", cfg.TailnetOrg),
		CloudflareKeyCheck:             "https://api.cloudflare.com/client/v4/user/tokens/verify",
		CloudflareAddRecord:            fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", cfg.CloudflareZoneId),
		CloudflareGetDomainFromZoneId:  fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s", cfg.CloudflareZoneId),
		CloudflareGetRecordsFromZoneId: fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", cfg.CloudflareZoneId),
		CloudflareDeleteRecordById:     fmt.Sprintf("https://api.cloudflare.com/client/v4/zones/%s/dns_records", cfg.CloudflareZoneId),
	}
	return endpoints
}
