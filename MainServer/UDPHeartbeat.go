package main

import (
	"ThinkGOGameServer/thinkutils"
	"fmt"
	"github.com/emirpasic/gods/maps/hashmap"
	"gopkg.in/ini.v1"
	"net"
	"time"
)

type UDPHeartbeat struct {
}

var (
	g_mapServer *hashmap.Map
)

func (this *UDPHeartbeat) heartbeat()  {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		log.Error("Read app.ini failed")
		return
	}

	serverAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", cfg.Section("register_center").Key("host").String(), cfg.Section("register_center").Key("udp_port").MustInt(8083)))
	if err != nil {
		return
	}

	localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		return
	}

	pConn, err := net.DialUDP("udp", localAddr, serverAddr)
	if err != nil {
		return
	}
	defer pConn.Close()

	szMsg := fmt.Sprintf(`{"type": "main", "port": %d}`, cfg.Section("main_server").Key("udp_port").MustInt(8084))
	_, err = pConn.Write(thinkutils.StringUtils.StringToBytes(szMsg))
	if err != nil {
		return
	}

	buf := make([]byte, 1024)
	n, addr, err := pConn.ReadFromUDP(buf)
	fmt.Println("Received ", string(buf[0:n]), " from ", addr)
}

func (this *UDPHeartbeat) Init() {

	g_mapServer = hashmap.New()

	for {
		go this.heartbeat()

		time.Sleep(5 * time.Second)
	}
}

