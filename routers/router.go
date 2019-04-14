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
)

func init() {
	beego.Router("/api/v1/create", &controllers.CreateController{}, "*:Run")
	beego.Router("/api/v1/query", &controllers.QueryController{}, "*:Run")
	beego.Router("/?:url", &controllers.JumpController{}, "*:Run")

	// 中间件过滤
	beego.InsertFilter("/api/*", beego.BeforeRouter, filterLoggedInUser)
}
