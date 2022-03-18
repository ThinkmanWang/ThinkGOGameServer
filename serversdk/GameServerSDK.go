package serversdk

import (
	"ThinkGOGameServer/thinkutils"
	"ThinkGOGameServer/thinkutils/logger"
	"github.com/emirpasic/gods/lists/arraylist"
	"net"
)

var (
	log *logger.LocalLogger = logger.DefaultLogger()
	SERVER_TYPE_MAIN string = "main"
	SERVER_TYPE_GAME string = "game"
	SERVER_MAIN_APPID uint64 = 0
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

type GameServerSDK struct {
	m_pUDPHeartbeat *UDPHeartbeat
	m_pUDPServer *thinkutils.UDPServer
	m_pGameServer IGameServer
}

func (this *GameServerSDK) Init(server IGameServer)  {
	this.m_pGameServer = server

	this.initUDPHeartbeat(server)
	this.initUDPPort(server)
}

func (this *GameServerSDK) onUDPMsg(pConn *net.UDPConn, addr net.Addr, data []byte) {
	logger.Info("Received %d bytes", len(data))
}

func (this *GameServerSDK) initUDPPort(server IGameServer)  {
	info := server.OnInitGameData()

	this.m_pUDPServer = &thinkutils.UDPServer{OnMsg: this.onUDPMsg}
	go this.m_pUDPServer.Start(int(info.Port))

	logger.Info("UDP Server started. port: %d", info.Port)
}

func (this *GameServerSDK) initUDPHeartbeat(server IGameServer)  {
	info := server.OnInitGameData()

	this.m_pUDPHeartbeat = &UDPHeartbeat{
		ServerInfo: info,
	}

	go this.m_pUDPHeartbeat.Init()
}

func (this *GameServerSDK) RandServer(szType string, nAppId uint64) *ServerNode {
	lstNode := this.GetAllServer(szType, nAppId)

	if nil == lstNode || len(lstNode) <= 0 {
		return nil
	}

	nPos := thinkutils.RandUtils.RandInt(0, len(lstNode))
	pNode := lstNode[nPos]

	return pNode
}

func (this *GameServerSDK) GetAllServer(szType string, nAppId uint64) []*ServerNode {
	lstRet := make([]*ServerNode, 0)

	pLstNode, bFound := g_mapServer.Get(szType)
	if false == bFound {
		return nil
	}

	for i := 0; i < pLstNode.(*arraylist.List).Size(); i++ {
		_pNode, bFound := pLstNode.(*arraylist.List).Get(i)
		if false == bFound {
			continue
		}

		pNode := _pNode.(*ServerNode)
		if nAppId != pNode.AppId {
			continue
		}

		lstRet = append(lstRet, pNode)
	}

	return lstRet
}

func (this *GameServerSDK) SendData(pNode *ServerNode, data []byte)  {
	if nil == pNode {
		return
	}

	go thinkutils.UDPUtils.Send(pNode.Host, int(pNode.Port), data)
}

func (this *GameServerSDK) SendToRandGameServer(nAppId uint64, data []byte) {
	pNode := this.RandServer(SERVER_TYPE_GAME, nAppId)
	if nil == pNode {
		return
	}

	this.SendData(pNode, data)
}

func (this *GameServerSDK) SendToAllGameServer(nAppId uint64, data []byte) {
	lstNode := this.GetAllServer(SERVER_TYPE_GAME, nAppId)
	if nil == lstNode || len(lstNode) <= 0 {
		return
	}

	for i := 0; i < len(lstNode); i++ {
		this.SendData(lstNode[i], data)
	}
}

func (this *GameServerSDK) SendToRandMainServer(data []byte) {
	pNode := this.RandServer(SERVER_TYPE_MAIN, SERVER_MAIN_APPID)
	if nil == pNode {
		return
	}

	this.SendData(pNode, data)
}

func (this *GameServerSDK) SendToAllMainServer(data []byte) {
	lstNode := this.GetAllServer(SERVER_TYPE_GAME, SERVER_MAIN_APPID)
	if nil == lstNode || len(lstNode) <= 0 {
		return
	}

	for i := 0; i < len(lstNode); i++ {
		this.SendData(lstNode[i], data)
	}
}

func (this *GameServerSDK) SendToClient(data []byte) {
	/*
	 1. make UDP package
	 2. send to main server
	 */
}
