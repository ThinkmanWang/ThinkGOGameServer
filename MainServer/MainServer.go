package main

import (
	"ThinkGOGameServer/serversdk"
	"ThinkGOGameServer/thinkutils/logger"
	"gopkg.in/ini.v1"
	"runtime"
	"sync"
)

var (
	log *logger.LocalLogger = logger.DefaultLogger()
)

type Mainerver struct {
	
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Info("Hello World")

	cfg, err := ini.Load("app.ini")
	if err != nil {
		log.Error("Read app.ini failed")
		return
	}

	pUDPHeartbeat := &serversdk.UDPHeartbeat{
		ServerInfo: serversdk.GameServerInfo{AppId: 0,
			Type: "main",
			Port: uint32(cfg.Section("main_server").Key("udp_port").MustUint(8084))},
	}
	go pUDPHeartbeat.Init()

	wg := sync.WaitGroup{}
	wg.Add(1)

	wg.Wait()
}
