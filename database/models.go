package database

type ApiKey struct {
	ApiKey     string `json:"api_key"`
	AllowedIPs string `json:"allowed_ips"`
	Name       string `json:"name"`
}

type Client struct {
	Id              int    `json:"id"`
	PublicKey       string `json:"public_key"`
	WireguardConfig string `json:"wireguard_config"`
	ExpiryDate      int    `json:"expiry_date"`
}
