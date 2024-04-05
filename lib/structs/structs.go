package lib

type Keys struct {
	TailscaleApiKey  string
	CloudflareApiKey string
}

type Device struct {
	Addresses []string `json:"addresses"`
	Name      string   `json:"name"`
	Id        string   `json:"id"`
}

type Devices struct {
	Devices []Device `json:"devices"`
}
type Endpoints struct {
	TailscaleDevices               string
	CloudflareKeyCheck             string
	CloudflareAddRecord            string
	CloudflareGetDomainFromZoneId  string
	CloudflareGetRecordsFromZoneId string
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

type CloudflareDnsRecord struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Type    string `json:"type"`
}

type CloudflareDnsRecordsResponse struct {
	Result []CloudflareDnsRecord `json:"result"`
}
