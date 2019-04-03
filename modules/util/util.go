package util

import (
	"github.com/astaxie/beego/context"
	redisDB "shorturl/db/redis"
	"encoding/json"
	"strings"
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

func JsonDecode(jsonStr string, structModel interface{}) error {
	decode := json.NewDecoder(strings.NewReader(jsonStr))
	err := decode.Decode(structModel)
	return err
}

func JsonEncode(structModel interface{}) (string, error) {
	jsonStr, err := json.Marshal(structModel)
	return string(jsonStr), err
}
