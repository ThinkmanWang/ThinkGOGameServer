package main

import (
	"ThinkGOGameServer/thinkutils/logger"
	//"./thinkutils/logger"
)

var (
	log *logger.LocalLogger = logger.DefaultLogger()
)

func main() {
    log.Info("Hello World")
}
