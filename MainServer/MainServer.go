package main

import (
	"ThinkGOGameServer/serversdk"
	"ThinkGOGameServer/thinkutils"
	"ThinkGOGameServer/thinkutils/logger"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"gopkg.in/ini.v1"
	"runtime"
)

var (
	log *logger.LocalLogger         = logger.DefaultLogger()
	g_pSDK *serversdk.MainServerSDK = &serversdk.MainServerSDK{}
)

type Mainerver struct {
}

func (this *Mainerver) OnInitGameData() serversdk.GameServerInfo {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		log.Error("Read app.ini failed")

		return serversdk.GameServerInfo{
			AppId: serversdk.SERVER_MAIN_APPID,
			Type: serversdk.SERVER_TYPE_MAIN,
			Port: 8084,
		}
	}

	return serversdk.GameServerInfo{
		AppId: serversdk.SERVER_MAIN_APPID,
		Type: serversdk.SERVER_TYPE_MAIN,
		Port: uint32(cfg.Section("main_server").Key("udp_port").MustUint(8084)),
	}
}

func (this *Mainerver) OnInitWS() serversdk.MainServerInfo {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		log.Error("Read app.ini failed")

		return serversdk.MainServerInfo{
			Path: "game",
			Port: 8082,
			HeartbeatTimeout: 10,
		}
	}

	return serversdk.MainServerInfo{
		Path: "game",
		Port: uint32(cfg.Section("main_server").Key("ws_port").MustUint(8082)),
		HeartbeatTimeout: uint32(cfg.Section("main_server").Key("heartbeat").MustUint(10)),
	}
}

func (this *Mainerver) OnWSConnect(pConn *websocket.Conn) {
	log.Info("New Connect")
}

func (this *Mainerver) OnWSMsg(pConn *websocket.Conn, data []byte) {
	pGamePkg := &serversdk.GamePkg{}
	err := proto.Unmarshal(data, pGamePkg)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	log.Info("%s", thinkutils.JSONUtils.ToJson(pGamePkg))

	switch *pGamePkg.Type {
	case serversdk.HeadType_LOGIN_REQUEST:
		this.doLogin(pConn, pGamePkg)
	case serversdk.HeadType_HEARTBEAT_REQUEST:
		this.doHeartbeatReq(pConn, pGamePkg)
	}
}

func (this *Mainerver) OnWSClose(pConn *websocket.Conn) {

	log.Info("Conn closed")
}

func (this *Mainerver) OnWSTimeout(pConn *websocket.Conn) {
	log.Info("Heartbeat timeout")
}

func (this *Mainerver) doLogin(pConn *websocket.Conn, pReq *serversdk.GamePkg)  {
	var nCode int32 = 200
	szMsg := "success"
	pLoginResp := &serversdk.LoginResponse{
		Code: &nCode,
		Msg: &szMsg,
	}

	nType := serversdk.HeadType_LOGIN_RESPONSE
	nTimestamp := thinkutils.DateTime.Timestamp()
	pResp := &serversdk.GamePkg{
		Type: &nType,
		Uid: pReq.Uid,
		Timestamp: &nTimestamp,
		LoginResponse: pLoginResp,
	}

	pData, err := proto.Marshal(pResp)
	if err != nil {
		return
	}

	err = pConn.WriteMessage(websocket.BinaryMessage, pData)
	if err != nil {
		log.Info("write:", err.Error())
	}
}

func (this *Mainerver) doHeartbeatReq(pConn *websocket.Conn, pReq *serversdk.GamePkg)  {
	nType := serversdk.HeadType_HEARTBEAT_RESPONSE
	nTimestamp := thinkutils.DateTime.Timestamp()
	pResp := &serversdk.GamePkg{
		Type: &nType,
		Uid: pReq.Uid,
		Timestamp: &nTimestamp,
	}

	pData, err := proto.Marshal(pResp)
	if err != nil {
		return
	}

	err = pConn.WriteMessage(websocket.BinaryMessage, pData)
	if err != nil {
		log.Info("write:", err.Error())
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Info("Hello World")
	
	pMainServer := &Mainerver{}
	g_pSDK.Init(pMainServer)
}
