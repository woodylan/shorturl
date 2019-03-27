package util

import (
	"github.com/astaxie/beego/context"
	"shorturl/redisDB"
)

type RetData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ThrowApi(ctx *context.Context, code int, msg string) (*context.Context) {
	var retData RetData
	retData.Code = code
	retData.Msg = msg

	ctx.Output.JSON(retData, true, false)
	return ctx
}

// redis发号器key
func GetRedisNum(key string) int64 {
	existValue, _ := redisDB.RedisConnect.Exists(key).Result()
	if existValue == 0 {
		redisDB.RedisConnect.Set(key, 1, 0)
	}

	val, _ := redisDB.RedisConnect.Incr(key).Result()
	return val
}
