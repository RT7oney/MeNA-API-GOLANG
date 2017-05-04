package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["MeNA-Api/controllers/Admin:AdminController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers/Admin:AdminController"],
		beego.ControllerComments{
			Method: "AddCharacter",
			Router: `/add-character`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers/Admin:AdminController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers/Admin:AdminController"],
		beego.ControllerComments{
			Method: "AddMovieTag",
			Router: `/add-movie_tag`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers/Admin:AdminController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers/Admin:AdminController"],
		beego.ControllerComments{
			Method: "AddMusicTag",
			Router: `/add-music_tag`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers/Admin:AdminController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers/Admin:AdminController"],
		beego.ControllerComments{
			Method: "AddDev",
			Router: `/add-app_key`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
