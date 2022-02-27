package main

import (
	"ThinkGOGameServer/thinkutils"
	"ThinkGOGameServer/thinkutils/logger"
	"time"
)

var (
	log *logger.LocalLogger = logger.DefaultLogger()
)

func worker(ch chan string) {
	log.Info("%p", ch)
	for {
		szTxt := <-ch
		log.Info(szTxt)
	}
}

func main() {
	ch := make(chan string)
	log.Info("%p", ch)
	go worker(ch)

	for {
		ch <- thinkutils.DateTime.CurDatetime()
		time.Sleep(time.Second)
	}
}
