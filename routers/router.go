// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"apiproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//使用注解路由
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.VedioController{})
	beego.Include(&controllers.Basecontroller{})
	beego.Include(&controllers.CommentController{})
	beego.Include(&controllers.TopController{})
}
