package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["MeNA-Api/controllers/Admin:AdminController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers/Admin:AdminController"],
		beego.ControllerComments{
			"AddCharacter",
			`/add-character`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers/Admin:AdminController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers/Admin:AdminController"],
		beego.ControllerComments{
			"AddMovieTag",
			`/add-movie_tag`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers/Admin:AdminController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers/Admin:AdminController"],
		beego.ControllerComments{
			"AddMusicTag",
			`/add-music_tag`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["MeNA-Api/controllers/Admin:AdminController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers/Admin:AdminController"],
		beego.ControllerComments{
			"AddDev",
			`/add-app_key`,
			[]string{"post"},
			nil})

}
