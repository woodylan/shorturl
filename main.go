package main

import (
	_ "shorturl/routers"

	"github.com/astaxie/beego"
	redisDB "shorturl/db/redis"
	mysqlDB "shorturl/db/mysql"
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

	// 连接MySQL数据库
	mysqlConnect, err := mysqlDB.Connect()
	if err != nil {
		panic(err.Error())
	}
	defer mysqlConnect.Close()

	beego.Run()
}
