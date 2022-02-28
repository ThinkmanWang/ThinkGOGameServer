package netutils

import (
	"ThinkGOGameServer/thinkutils"
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
	"gopkg.in/ini.v1"
	"ThinkGOGameServer/thinkutils/logger"
	"time"
)

type UDPHeartbeat struct {
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

	szMsg := fmt.Sprintf(`{"type": "main", "port": %d}`, cfg.Section("main_server").Key("udp_port").MustInt(8084))
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

