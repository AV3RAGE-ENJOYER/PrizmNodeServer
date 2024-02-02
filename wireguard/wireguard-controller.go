package wireguard

import "node/database"

type WgController struct {
	Clients    []database.Client
	EndpointIp string
}

func NewWgController() WgController {
	wg := WgController{}
	return wg
}

func (wg *WgController) GenerateConfig() string {
	return "test"
}
