package main

type Keys struct {
	TailscaleApiKey  string
	CloudflareApiKey string
}

type States struct {
	TailscaleKeyExist  bool
	CloudflareKeyExist bool
	TailnetOrgExist    bool
}

type Config struct {
	Version    string
	States     States
	Keys       Keys
	TailnetOrg string
}
