package chat

import (
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许跨域请求
	},
}
var clientsMutex = sync.RWMutex{}

type ClientInfo struct {
	Conn bool
	Uid  int64
}

var Clients = make(map[*websocket.Conn]*ClientInfo) // 已连接的客户端

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logs.Error(err)
	}

	defer func() {
		if err := recover(); err != nil {
			logs.Error("HandleConnections is error %s", err)
		}
		ws.Close()
		logs.Info("ws is quit")
		RemoveClient(ws)
	}()

	logs.Info("new connect.....")
	Clients[ws] = &ClientInfo{Conn: true}

	for {
		msg := &Message{}
		err := ws.ReadJSON(msg)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}

		MsgChannel <- &InnerMsg{Ws: ws, Message: msg}
	}

}

func AddClient(ws *websocket.Conn, uid int64) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	client, ok := Clients[ws]
	if !ok {
		return
	}
	client.Uid = uid
}

func RemoveClient(ws *websocket.Conn) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	delete(Clients, ws)
}
