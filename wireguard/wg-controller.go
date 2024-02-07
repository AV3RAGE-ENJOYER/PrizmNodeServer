package wireguard

import (
	"fmt"
	"node/utils"
	"os/exec"
	"strings"
	"time"
)

type WgController struct {
	Interface WgInterface
	Clients   []WgClient
}

func NewWgController() WgController {
	wg := WgController{}
	return wg
}

func (wgController *WgController) DumpClients() error {
	out, err := exec.Command("wg", "show", "wg0", "dump").Output()

	if err != nil {
		fmt.Println(err)
		return err
	}

	dump := strings.Fields(string(out))

	wgController.Interface = WgInterface{
		PrivateKey: dump[0],
		PublicKey:  dump[1],
		ListenPort: utils.ToInt(dump[2]),
		FwMark:     dump[3],
	}

	dump = dump[4:]

	for i := 0; i < len(dump)/8; i++ {
		if dump != nil {
			client := WgClient{
				PublicKey:    dump[0],
				PresharedKey: dump[1],
				Endpoint:     dump[2],
				AllowedIPs:   dump[3],
				LatestHandshake: time.Unix(
					int64(utils.ToInt(
						dump[4])), 0),
				TransferRx:          utils.ToInt(dump[5]),
				TransferTx:          utils.ToInt(dump[6]),
				PersistentKeepAlive: dump[7],
			}

			wgController.Clients = append(wgController.Clients, client)
			dump = dump[8:]
		}

	}

	return nil
}

func (wgController *WgController) AddClient(id int) (string, string, error) {
	clientConf, err := exec.Command("python3", "/PrizmNodeServer/create_client.py", fmt.Sprintf("%v", id)).Output()

	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	publicKey, _ := exec.Command("cat", fmt.Sprintf("/etc/wireguard/clients/%v_public.key", id)).Output()

	return string(clientConf), string(publicKey), nil
}
