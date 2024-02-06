package wireguard

import (
	"fmt"
	"os/exec"
	"strconv"
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

	dump := strings.Split(string(out), "\n")

	fmt.Println(dump)

	interfaceInfo := strings.Fields(dump[0])

	wgController.Interface = WgInterface{
		PrivateKey: interfaceInfo[0],
		PublicKey:  interfaceInfo[1],
		ListenPort: func() int {
			port, _ := strconv.Atoi(interfaceInfo[2])
			return port
		}(),
		FwMark: interfaceInfo[3],
	}

	dump = dump[1:]

	for _, client := range dump {
		clientInfo := strings.Fields(client)
		wgController.Clients = append(wgController.Clients, WgClient{
			PublicKey:    clientInfo[0],
			PresharedKey: clientInfo[1],
			Endpoint:     clientInfo[2],
			AllowedIPs:   clientInfo[3],
			LatestHandshake: func() time.Time {
				unixTime, _ := strconv.Atoi(clientInfo[4])
				return time.Unix(int64(unixTime), 0)
			}(),
			TransferRx: func() int {
				transferRx, _ := strconv.Atoi(clientInfo[5])
				return transferRx
			}(),
			TransferTx: func() int {
				transferTx, _ := strconv.Atoi(clientInfo[6])
				return transferTx
			}(),
			PersistentKeepAlive: clientInfo[7],
		})
	}

	return nil
}
