package serversdk

import (
	"ThinkGOGameServer/thinkutils"
	"ThinkGOGameServer/thinkutils/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net"
)

type MainServerInfo struct {
	Path string
	Port uint32
	HeartbeatTimeout uint32
}

type IMainServer interface {
	OnInitGameData() GameServerInfo
	OnInitWS() MainServerInfo
	OnWSConnect(pConn *websocket.Conn)
	OnWSMsg(pConn *websocket.Conn, msg []byte)
	OnWSClose(pConn *websocket.Conn)
	OnWSTimeout(pConn *websocket.Conn)
}

type MainServerSDK struct {
	m_pUDPHeartbeat *UDPHeartbeat
	m_pUDPServer *thinkutils.UDPServer
	m_pMainServer IMainServer
}

func (this *MainServerSDK) Init(server IMainServer)  {
	this.m_pMainServer = server

	this.initUDPHeartbeat(server)
	this.initUDPPort(server)
	this.initWebsocket(server)
}

func (this *MainServerSDK) onUDPMsg(pConn *net.UDPConn, addr net.Addr, data []byte) {
	log.Info("Receive %d bytes", len(data))
}

func (this *MainServerSDK) initUDPPort(server IMainServer)  {
	info := server.OnInitGameData()

	this.m_pUDPServer = &thinkutils.UDPServer{OnMsg: this.onUDPMsg}
	go this.m_pUDPServer.Start(int(info.Port))

	logger.Info("UDP Server started. port: %d", info.Port)
}

func (this *MainServerSDK) initUDPHeartbeat(server IMainServer)  {
	info := server.OnInitGameData()

	this.m_pUDPHeartbeat = &UDPHeartbeat{
		ServerInfo: info,
	}

	go this.m_pUDPHeartbeat.Init()
}

func (this *MainServerSDK) initWebsocket(server IMainServer) {
	serverInfo := server.OnInitWS()

	pHandler := &thinkutils.WSHandler{
		OnConnect:        server.OnWSConnect,
		OnMsg:            server.OnWSMsg,
		OnClose:          server.OnWSClose,
		OnTimeout:        server.OnWSTimeout,
		HeartbeatTimeout: serverInfo.HeartbeatTimeout,
	}

	pHandler.Init()

	eng := gin.Default()
	// 路由组1 ，处理GET请求
	eng.GET(serverInfo.Path, pHandler.Handler)

	szPort := fmt.Sprintf(":%d", serverInfo.Port)
	eng.Run(szPort)
}