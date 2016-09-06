package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["MeNA-Api/controllers/User:RegisterController"] = append(beego.GlobalControllerRouter["MeNA-Api/controllers/User:RegisterController"],
		beego.ControllerComments{
			"Register",
			`/register`,
			[]string{"post"},
			nil})

}
