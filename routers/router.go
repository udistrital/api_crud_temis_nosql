// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api_crud_temis_nosql/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/ipc",
			beego.NSInclude(
				&controllers.IpcController{},
			),
		),

		beego.NSNamespace("/entidadcuotaspartes",
			beego.NSInclude(
				&controllers.EntidadCuotasPartesController{},
			),
		),

		beego.NSNamespace("/rolentidadcuotaspartes",
			beego.NSInclude(
				&controllers.RolEntidadCuotasPartesController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
