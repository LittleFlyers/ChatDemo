// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"room/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/signin", &controllers.UserController{}, "get:SignIn")
	beego.Router("/user/login", &controllers.UserController{}, "get:LogIn")
	beego.Router("/chartroom", &controllers.ChartroomController{})
	beego.Router("/chartroom/add", &controllers.ChartroomController{}, "get:Add")
	beego.Router("/chart", &controllers.ChartController{})
	beego.Router("/chart/add", &controllers.ChartController{}, "get:Add")
	beego.Router("/chart/getchat", &controllers.ChartController{}, "get:GetChat")
}
