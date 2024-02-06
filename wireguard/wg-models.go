package wireguard

import (
	"fmt"
	"time"
)

type WgInterface struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	ListenPort int    `json:"listen_port"`
	FwMark     string `json:"fwmark"`
}

type WgClient struct {
	PublicKey           string    `json:"public_key"`
	PresharedKey        string    `json:"preshared_key"`
	Endpoint            string    `json:"endpoint"`
	AllowedIPs          string    `json:"allowed_ips"`
	LatestHandshake     time.Time `json:"latest_handshake"`
	TransferRx          int       `json:"transfer_rx"`
	TransferTx          int       `json:"transfer_tx"`
	PersistentKeepAlive string    `json:"persistent_keep_alive"`
}

func (WgI *WgInterface) String() string {
	return fmt.Sprintf(
		"%s %s %v %s",
		WgI.PrivateKey, WgI.PublicKey, WgI.ListenPort, WgI.FwMark)
}
