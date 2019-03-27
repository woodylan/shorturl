package main

import (
	_ "shorturl/routers"

	"github.com/astaxie/beego"
	"shorturl/redisDB"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 连接redis
	redisConnect, err := redisDB.Connect()
	if err != nil {
		// 数据库连接错误
		panic(err.Error())
	}
	defer redisConnect.Close()

	beego.Run()
}
