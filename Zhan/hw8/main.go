package main

import (
	_ "hw8/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.SetLogger(logs.AdapterFile, `{"filename":"test.log"}`)

	logs.Info("Http port:", beego.AppConfig.String("httpport"))
	logs.Info("Database path:", beego.AppConfig.String("dbpath"))

	beego.Run()
}
