package serversdk

import (
	"ThinkGOGameServer/serversdk"
	"ThinkGOGameServer/thinkutils"
	"ThinkGOGameServer/thinkutils/logger"
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
	"gopkg.in/ini.v1"
	"time"
)

type UDPHeartbeat struct {
	ServerInfo serversdk.GameServerInfo
}

var (
	g_mapServer *hashmap.Map
	log *logger.LocalLogger = logger.DefaultLogger()
)

func (this *UDPHeartbeat) heartbeat()  {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		log.Error("Read app.ini failed")
		return
	}

	szMsg := thinkutils.JSONUtils.ToJson(this.ServerInfo)
	pConn := thinkutils.UDPUtils.Send(cfg.Section("register_center").Key("host").String(),
		cfg.Section("register_center").Key("udp_port").MustInt(8083),
		thinkutils.StringUtils.StringToBytes(szMsg))

	if nil == pConn {
		return
	}

	buf := make([]byte, 256)
	n, addr, err := pConn.ReadFromUDP(buf)
	if err != nil {
		return
	}
	fmt.Println("Received ", string(buf[0:n]), " from ", addr)
}

func (this *UDPHeartbeat) Init() {

	g_mapServer = hashmap.New()

	for {
		go this.heartbeat()

		time.Sleep(5 * time.Second)
	}
}
