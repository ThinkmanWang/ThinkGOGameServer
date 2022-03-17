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
	g_pSDK *serversdk.ServerSDK = &serversdk.ServerSDK{}
)

type Mainerver struct {
	serversdk.IGameServer
}

func (this *Mainerver) OnInitGameData() serversdk.GameServerInfo {
	cfg, err := ini.Load("app.ini")
	if err != nil {
		log.Error("Read app.ini failed")
		
		return serversdk.GameServerInfo{AppId: 0,
			Type: "main",
			Port: 0,
		}
	}
	
	return serversdk.GameServerInfo{AppId: 0,
		Type: "main",
		Port: uint32(cfg.Section("main_server").Key("udp_port").MustUint(8084)),
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Info("Hello World")
	
	pMainServer := &Mainerver{}
	g_pSDK.Init(pMainServer)

	wg := sync.WaitGroup{}
	wg.Add(1)

	wg.Wait()
}
