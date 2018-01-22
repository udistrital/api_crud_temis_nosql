package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:EntidadCuotasPartesController"] = append(beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:EntidadCuotasPartesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:EntidadCuotasPartesController"] = append(beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:EntidadCuotasPartesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:EntidadCuotasPartesController"] = append(beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:EntidadCuotasPartesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:EntidadCuotasPartesController"] = append(beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:EntidadCuotasPartesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:IpcController"] = append(beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:IpcController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:IpcController"] = append(beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:IpcController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:IpcController"] = append(beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:IpcController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:IpcController"] = append(beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:IpcController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:RolEntidadCuotasPartesController"] = append(beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:RolEntidadCuotasPartesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:RolEntidadCuotasPartesController"] = append(beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:RolEntidadCuotasPartesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:RolEntidadCuotasPartesController"] = append(beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:RolEntidadCuotasPartesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:RolEntidadCuotasPartesController"] = append(beego.GlobalControllerRouter["api_crud_temis_nosql/controllers:RolEntidadCuotasPartesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

}
