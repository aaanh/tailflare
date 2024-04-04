package main

type Keys struct {
	TailscaleApiKey  string
	CloudflareApiKey string
}

type States struct {
	TailscaleKeyExist     bool
	TailnetOrgExist       bool
	CloudflareKeyExist    bool
	CloudflareZoneIdExist bool
}

type Config struct {
	Version          string
	States           States
	Keys             Keys
	TailnetOrg       string
	CloudflareZoneId string
}

type Payload struct {
	Content string `json:"content"`
	Name    string `json:"name"`
	Proxied bool   `json:"proxied"`
	Type    string `json:"type"`
	Ttl     int    `json:"ttl"`
}
