package serversdk

import (
	serversdk "ThinkGOGameServer/serversdk/netutils"
	"gopkg.in/ini.v1"
)

type GameServerInfo struct {
	AppId uint64 `json:"appid"`
	Type string `json:"type"`
	Port uint32 `json:"port"`
}

type Room struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

type IGameServer interface {
	OnInitGameData() GameServerInfo
	OnReceiveFromClient(user interface{}, data []byte)
	OnJoinRoom(user interface{}, room Room)
	OnCreateRoom(user interface{}, room Room)
	OnLeaveRoom(user interface{}, room Room)
	OnRemovePlayer(user interface{}, room Room)
	OnDestroyRoom(room Room)
	OnChangeRoom(room Room)
	OnChangeCustomPlayerStatus(user interface{}, data []byte)
	OnChangePlayerNetworkState(user interface{}, data []byte)
}

type GameServer struct {

}

func (this *GameServer) Init()  {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		log.Error("Read app.ini failed")
		return
	}

	pUDPHeartbeat := &serversdk.UDPHeartbeat{
		ServerInfo: GameServerInfo{AppId: 0,
			Type: "main",
			Port: uint32(cfg.Section("game_server").Key("udp_port").MustUint(8085))},
	}

	go pUDPHeartbeat.Init()
}