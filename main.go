package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"net/http"
	"os"
	"path/filepath"
	"wjszm-chat/chat"
)

func initConfig() {
	WorkPath, _ := os.Getwd()
	configPath := filepath.Join(WorkPath, "conf", "config.ini")
	if err := beego.LoadAppConfig("ini", configPath); nil != err {
		panic(err)
	}
}

func init() {
	initConfig()
}

func main() {
	defer func() {
		recover()
	}()

	http.HandleFunc("/", chat.HandleConnections)

	go chat.HandleMessages()

	host := beego.AppConfig.String("host")
	logs.Info("server start success is %s", host)
	err := http.ListenAndServe(host, nil)
	if err != nil {
		logs.Error("ListenAndServe: ", err)
	}
}
