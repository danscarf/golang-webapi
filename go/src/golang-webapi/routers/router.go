// @APIVersion 1.0.0
// @Title Golang-Docker-MongoDB Sample API
// @Description Sample API that performs basic CRUD operations on 'users'.
// @Contact dscarf@gmail.com
// @TermsOfServiceUrl http://www.danscarf.com/
// @License MIT
// @LicenseUrl https://opensource.org/licenses/MIT
package routers

import (
	"golang-webapi/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
