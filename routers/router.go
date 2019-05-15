// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"shorturl/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"shorturl/modules/util"
)

func init() {
	beego.Get("/", func(ctx *context.Context) {
		retData := util.RetData{Code: 200, Msg: "Hello"}
		ctx.Output.JSON(retData, true, false)
	})

	beego.Router("/api/v1/create", &controllers.CreateController{}, "*:Run")
	beego.Router("/api/v1/query", &controllers.QueryController{}, "*:Run")
	beego.Router("/:url([a-zA-z0-9]{6,})", &controllers.JumpController{}, "*:Run")

	beego.Get("/*", func(ctx *context.Context) {
		retData := util.RetData{Code: 404, Msg: "Not Found"}
		ctx.Output.JSON(retData, true, false)
	})

	// 中间件过滤
	beego.InsertFilter("/api/*", beego.BeforeRouter, filterLoggedInUser)
}
