package routers

import (
	"github.com/astaxie/beego/context"
	"shorturl/modules/util"
	"shorturl/global"
	"shorturl/models/tokenModel"
	"shorturl/define/retcode"
)

var filterLoggedInUser = func(ctx *context.Context) {
	token := ctx.Input.Header("Token")
	if token != "" {
		tokenInfo := tokenModel.TokenModel{}

		model, err := tokenModel.GetByToken(token)
		if err != nil {
			util.ThrowApi(ctx, retcode.TokenNotFound, "Token不存在")
			return
		}

		tokenInfo.Id = model.Id
		tokenInfo.Name = model.Name
		tokenInfo.Token = model.Token

		//存到全局
		global.TokenInfo = &tokenInfo
	} else {
		util.ThrowApi(ctx, retcode.TokenNotFound, "Header缺少Token")
		return
	}
}
