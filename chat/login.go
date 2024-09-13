package chat

import "github.com/astaxie/beego/logs"

func login(msg *InnerMsg) {
	logs.Info("%d login success", msg.Message.Uid)
	AddClient(msg.Ws, msg.Message.Uid)
}
