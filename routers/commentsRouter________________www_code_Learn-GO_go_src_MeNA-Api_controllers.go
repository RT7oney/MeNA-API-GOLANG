package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["MeNA-Api/controllers:CharacterController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers:CharacterController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/getone`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers:CharacterController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers:CharacterController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/getall`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers:DevController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers:DevController"],
		beego.ControllerComments{
			Method: "GetAppKey",
			Router: `/get-app-key`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers:DevController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers:DevController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers:MovieController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers:MovieController"],
		beego.ControllerComments{
			Method: "GetAllTags",
			Router: `/get-all-tags`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers:MusicController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers:MusicController"],
		beego.ControllerComments{
			Method: "GetAllTags",
			Router: `/get-all-tags`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers:UserController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers:UserController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
