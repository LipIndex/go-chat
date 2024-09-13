package chat

import (
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
)

type InnerMsg struct {
	Ws      *websocket.Conn
	Message *Message
}

type Message struct {
	Action  int32 `json:"action"`
	Uid     int64 `json:"uid"`
	Payload any   `json:"payload"`
}

var MsgChannel = make(chan *InnerMsg) // 消息通道

func HandleMessages() {
	defer func() {
		if err := recover(); err != nil {
			logs.Error("handleMessages is error err:%s", err)
		}
	}()
	for {
		select {
		case msg := <-MsgChannel:
			messageHandle(msg)
		}
	}
}

func messageHandle(msg *InnerMsg) {
	switch msg.Message.Action {
	case LOGIN:
		login(msg)
	case SEND_MSG:
		sendMsg(msg)
	}
}

func sendMsg(msg *InnerMsg) {
	logs.Info("new msg %v", *msg.Message)
	for client, info := range Clients {
		if client == msg.Ws || info.Uid == 0 {
			continue
		}
		err := client.WriteJSON(msg)
		if err != nil {
			logs.Error("error: %v", err)
			client.Close()
			RemoveClient(client)
		}
	}
}
