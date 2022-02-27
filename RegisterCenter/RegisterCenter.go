package main

import (
	"ThinkGOGameServer/thinkutils"
	"ThinkGOGameServer/thinkutils/logger"
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
	"gopkg.in/ini.v1"
	"net"
)

var (
	log *logger.LocalLogger = logger.DefaultLogger()
	SERVER_MAIN string = "main"
	SERVER_GAME string = "game"
)

type ServerNode struct {
	Type string `json:"type"`
	Host string `json:"host"`
	Port int32 `json:"port"`
	AppId uint64 `json:"appid"`
	Heartbeat uint64 `json:"-"`
}

type RegisterCenter struct {
	m_nHeartbeatTime     uint32
	m_mapGameServer *hashmap.Map
	m_mapMainServer *hashmap.Map
	m_pHeartbeatMgr *thinkutils.HeartbeatMgr
}

func (this *RegisterCenter) Init()  {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		this.m_nHeartbeatTime = 10
	} else {
		this.m_nHeartbeatTime = uint32(cfg.Section("register_center").Key("heartbeat").MustInt(10))
	}


	this.m_mapGameServer = hashmap.New()
	this.m_mapMainServer = hashmap.New()


	this.m_pHeartbeatMgr = &thinkutils.HeartbeatMgr{}
	this.m_pHeartbeatMgr.Init(this.m_nHeartbeatTime, this.onTimeout)
}

func (this *RegisterCenter) onTimeout(conn interface{}) {
	if nil == conn {
		return
	}

	pNode := conn.(*ServerNode)
	if nil == pNode {
		return
	}

	log.Info("%s", thinkutils.JSONUtils.ToJson(pNode))

	log.Info(this.serverInfo())
	szKey := fmt.Sprintf("%s:%d", pNode.Host, pNode.Port)
	if SERVER_MAIN == pNode.Type {
		this.m_mapMainServer.Remove(szKey)
	} else if SERVER_GAME == pNode.Type {
		this.m_mapGameServer.Remove(szKey)
	} else {
		return
	}
	log.Info(this.serverInfo())
}

func (this *RegisterCenter) serverInfo() string {
	szInfo := fmt.Sprintf("Main Server: %d Game Server: %d", this.m_mapMainServer.Size(), this.m_mapGameServer.Size())
	return szInfo
}

//{"type": "main/game", "port": 8084, "appid": 1}
func (this *RegisterCenter) OnMsg(addr net.Addr, data []byte) {
	log.Info("<%s> %s", addr.(*net.UDPAddr).IP, thinkutils.StringUtils.BytesToString(data))

	_pNode := &ServerNode{}
	err := thinkutils.JSONUtils.FromJson(thinkutils.StringUtils.BytesToString(data), _pNode)
	if err != nil || nil == _pNode {
		return
	}

	_pNode.Host = addr.(*net.UDPAddr).IP.String()
	szKey := fmt.Sprintf("%s:%d", _pNode.Host, _pNode.Port)

	if SERVER_MAIN == _pNode.Type {
		pNode, bFound := this.m_mapMainServer.Get(szKey)
		if bFound {
			_pNode = pNode.(*ServerNode)
			_pNode.Heartbeat = uint64(thinkutils.DateTime.Timestamp())
		} else {
			_pNode.Host = addr.(*net.UDPAddr).IP.String()
			_pNode.Heartbeat = uint64(thinkutils.DateTime.Timestamp())

			this.m_mapMainServer.Put(szKey, _pNode)
			this.m_pHeartbeatMgr.Update(pNode)
		}
	} else if SERVER_GAME == _pNode.Type {
		pNode, bFound := this.m_mapGameServer.Get(szKey)
		if bFound {
			_pNode = pNode.(*ServerNode)
			_pNode.Heartbeat = uint64(thinkutils.DateTime.Timestamp())
		} else {
			_pNode.Host = addr.(*net.UDPAddr).IP.String()
			_pNode.Heartbeat = uint64(thinkutils.DateTime.Timestamp())

			this.m_mapGameServer.Put(szKey, _pNode)
		}
	} else {
		return
	}

	this.m_pHeartbeatMgr.Update(_pNode)

	thinkutils.UDPUtils.Send(addr.(*net.UDPAddr).IP.String(), addr.(*net.UDPAddr).Port, thinkutils.StringUtils.StringToBytes(this.replyMsg(_pNode.Type)))
}

func (this *RegisterCenter) replyMsg(serverType string) string {
	szRet := ""
	if SERVER_MAIN == serverType {
		szRet = fmt.Sprintf("%s", thinkutils.JSONUtils.ToJson(this.m_mapGameServer.Values()))
	} else if SERVER_GAME == serverType {
		szRet = fmt.Sprintf("%s", thinkutils.JSONUtils.ToJson(this.m_mapMainServer.Values()))
	} else {
		szRet = ""
	}

	return szRet
}

func startRegisterCenter()  {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		return
	}

	center := &RegisterCenter{}
	center.Init()

	pServer := &thinkutils.UDPServer{OnMsg: center.OnMsg}
	pServer.Start(cfg.Section("register_center").Key("udp_port").MustInt())
}

func main() {
	log.Info("Hello World")

	startRegisterCenter()
}
