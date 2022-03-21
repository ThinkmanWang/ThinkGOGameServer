package main

import (
	"ThinkGOGameServer/serversdk"
	"ThinkGOGameServer/thinkutils"
	"ThinkGOGameServer/thinkutils/logger"
	"fmt"
	"github.com/golang/protobuf/proto"
	"runtime"
)

var (
	log *logger.LocalLogger = logger.DefaultLogger()
)

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    log.Info("Hello World")

    nType := serversdk.HeadType_SEND_TO_SERVER
    szUid := "A00001"
    nTimestamp := thinkutils.DateTime.Timestamp()
    pSendToServer := &serversdk.SendToServer{
    	Data: thinkutils.StringUtils.StringToBytes("Hello World"),
	}
    gamePkg := &serversdk.GamePkg{
    	Type: &nType,
		Uid: &szUid,
		Timestamp: &nTimestamp,
		SendToServer: pSendToServer,
	}

	pData, err := proto.Marshal(gamePkg)
	if err != nil {
		panic(err)
	}

	fmt.Println(pData)

	destPkg := &serversdk.GamePkg{}
	err = proto.Unmarshal(pData, destPkg)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	log.Info("%s", thinkutils.JSONUtils.ToJson(destPkg))
}
