package serversdk

import (
	"ThinkGOGameServer/thinkutils/logger"
)

var (
	log *logger.LocalLogger = logger.DefaultLogger()
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

type ServerNode struct {
	Type string `json:"type"`
	Host string `json:"host"`
	Port int32 `json:"port"`
	AppId uint64 `json:"appid"`
	Heartbeat uint64 `json:"-"`
}

type ServerSDK struct {
	m_pUDPHeartbeat *UDPHeartbeat
}

func (this *ServerSDK) Init(server IGameServer)  {
	this.initUDPHeartbeat(server)
}

func (this *ServerSDK) initUDPHeartbeat(server IGameServer)  {
	info := server.OnInitGameData()

	this.m_pUDPHeartbeat = &UDPHeartbeat{
		ServerInfo: info,
	}

	go this.m_pUDPHeartbeat.Init()
}

func (this *ServerSDK) RandServerByType(szType string) *ServerNode {
	return nil
}

func (this *ServerSDK) RandServer(szType string, szApId uint64) *ServerNode {
	return nil
}
