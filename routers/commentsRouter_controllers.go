package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["sabtuu/controllers:PesertaController"] = append(beego.GlobalControllerRouter["sabtuu/controllers:PesertaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["sabtuu/controllers:PesertaController"] = append(beego.GlobalControllerRouter["sabtuu/controllers:PesertaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sabtuu/controllers:PesertaController"] = append(beego.GlobalControllerRouter["sabtuu/controllers:PesertaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sabtuu/controllers:PesertaController"] = append(beego.GlobalControllerRouter["sabtuu/controllers:PesertaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["sabtuu/controllers:PesertaController"] = append(beego.GlobalControllerRouter["sabtuu/controllers:PesertaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["sabtuu/controllers:SabtuuController"] = append(beego.GlobalControllerRouter["sabtuu/controllers:SabtuuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["sabtuu/controllers:SabtuuController"] = append(beego.GlobalControllerRouter["sabtuu/controllers:SabtuuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sabtuu/controllers:SabtuuController"] = append(beego.GlobalControllerRouter["sabtuu/controllers:SabtuuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["sabtuu/controllers:SabtuuController"] = append(beego.GlobalControllerRouter["sabtuu/controllers:SabtuuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["sabtuu/controllers:SabtuuController"] = append(beego.GlobalControllerRouter["sabtuu/controllers:SabtuuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

}
