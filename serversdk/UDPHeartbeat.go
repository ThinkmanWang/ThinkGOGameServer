package serversdk

import (
	"ThinkGOGameServer/thinkutils"
	"ThinkGOGameServer/thinkutils/logger"
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/maps/hashmap"
	"gopkg.in/ini.v1"
	"time"
)

type UDPHeartbeat struct {
	ServerInfo GameServerInfo
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

	var lstServer []ServerNode
	err = thinkutils.JSONUtils.FromJson(string(buf[0:n]), &lstServer)
	if err != nil {
		return
	}

	this.updateServer(lstServer)
}

func (this *UDPHeartbeat) updateServer(lstServer []ServerNode)  {
	if nil == lstServer {
		return
	}

	g_mapServer.Clear()

	for i := 0; i < len(lstServer); i++ {
		item := lstServer[i]
		lstNode, bFound := g_mapServer.Get(item.Type)
		if false == bFound {
			lstNode := arraylist.New()
			g_mapServer.Put(item.Type, lstNode)

			lstNode.Add(item)
		} else {
			lstNode.(*arraylist.List).Add(item)
		}
	}

	this.Print()
}

func (this *UDPHeartbeat) Print()  {
	lstKeys := g_mapServer.Keys()
	for i := 0; i < len(lstKeys); i++ {
		szKey := lstKeys[i]
		logger.Info("%s", szKey)

		lstNode, _ := g_mapServer.Get(szKey)
		for j := 0; j < lstNode.(*arraylist.List).Size(); j++ {
			node, _ := lstNode.(*arraylist.List).Get(j)
			logger.Info("%d. %s", j, thinkutils.JSONUtils.ToJson(node))
		}
	}

}


func (this *UDPHeartbeat) Init() {

	g_mapServer = hashmap.New()

	for {
		go this.heartbeat()

		time.Sleep(5 * time.Second)
	}
}

