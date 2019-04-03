package routers

import (
	"github.com/astaxie/beego/context"
	_ "github.com/astaxie/beego/cache/redis"
	"shorturl/modules/util"
	"shorturl/global"
	"shorturl/models/tokenModel"
)

var filterLoggedInUser = func(ctx *context.Context) {
	token := ctx.Input.Header("Token")
	if token != "" {
		tokenInfo := shorturlModel.TokenModel{}

		model, err := shorturlModel.GetByToken(token)
		if err != nil {
			util.ThrowApi(ctx, -1, "Token不存在")
			return
		}

		tokenInfo.Id = model.Id
		tokenInfo.Name = model.Name
		tokenInfo.Token = model.Token

		//存到全局
		global.TokenInfo = &tokenInfo
	} else {
		util.ThrowApi(ctx, -1, "Header缺少Token")
		return
	}
}
