package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["MeNA-Api/controllers:CharacterController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers:CharacterController"],
		beego.ControllerComments{
			"GetOne",
			`/getone`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers:CharacterController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers:CharacterController"],
		beego.ControllerComments{
			"GetAll",
			`/getall`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers:DevController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers:DevController"],
		beego.ControllerComments{
			"GetAppKey",
			`/get-app-key`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers:DevController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers:DevController"],
		beego.ControllerComments{
			"Register",
			`/register`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers:MovieController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers:MovieController"],
		beego.ControllerComments{
			"GetAllTags",
			`/get-all-tags`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers:MusicController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers:MusicController"],
		beego.ControllerComments{
			"GetAllTags",
			`/get-all-tags`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers:UserController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers:UserController"],
		beego.ControllerComments{
			"Register",
			`/register`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers:UserController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login`,
			[]string{"post"},
			nil})

}
