package main

import (
	"ThinkGOGameServer/netutils"
	"ThinkGOGameServer/thinkutils/logger"
	"runtime"
	"sync"
)

var (
	log *logger.LocalLogger = logger.DefaultLogger()
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Info("Hello World")

	pUDPHeartbeat := &netutils.UDPHeartbeat{}
	go pUDPHeartbeat.Init()

	wg := sync.WaitGroup{}
	wg.Add(1)

	wg.Wait()
}
