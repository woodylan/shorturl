package routers

import (
	"github.com/astaxie/beego/context"
	_ "github.com/astaxie/beego/cache/redis"
	redisDB "shorturl/db/redis"
	"shorturl/modules/util"
	"shorturl/logic"
	"shorturl/global"
)

var filterLoggedInUser = func(ctx *context.Context) {
	token := ctx.Input.Header("Token")
	if token != "" {
		tokenData := logic.TokenData{}

		redis := redisDB.RedisConnect.Get("shorturl:" + token)
		if redis == nil {
			util.ThrowApi(ctx, -1, "用户未登录")
			return
		}

		util.JsonDecode(string(redis.Val()), &tokenData)

		//存到全局
		global.TokenInfo = &tokenData
	} else {
		util.ThrowApi(ctx, -1, "缺少token")
		return
	}
}
