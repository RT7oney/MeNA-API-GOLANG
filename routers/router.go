// @APIVersion 1.0.0
// @Title MeNA-Api 文档
// @Description 一个想要改变世界的接口
// @Contact RyanTyler@gmail.com
package routers

import (
	// "MeNA-Api/common"
	"MeNA-Api/controllers/User"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	// "log"
)

func init() {
	var auth = func(ctx *context.Context) {
		headers := ctx.Request.Header
		/*****************************/
		// fmt.Println(headers["Secret-Key"])
		// fmt.Println(reflect.TypeOf(headers["Secret-Key"]))
		// fmt.Println(len(headers["Secret-Key"]))
		// fmt.Println(cap(headers["Secret-Key"]))

		// header_arr := headers["Secret-key"]
		// fmt.Println(header_arr)
		// fmt.Println(reflect.TypeOf(header_arr))
		/*****************************/
		// var msg string
		if headers["Secret-Key"] == nil {
			// msg = common.Response(400, "您没有权限")
			// ctx.Redirect(302, "/error")
			ctx.WriteString("您没有接口调用权限")
		}
	}
	// beego.InsertFilter("/*", beego.BeforeRouter, checkFilter)
	ns := beego.NewNamespace("/v1",
		// beego.NSNamespace("/object",
		// 	beego.NSInclude(
		// 		&controllers.ObjectController{},
		// 	),
		// ),
		// beego.NSNamespace("/user",
		// 	beego.NSInclude(
		// 		&controllers.UserController{},
		// 	),
		// ),
		beego.NSBefore(auth),
		beego.NSNamespace("/user",
			beego.NSRouter("/register", &controllers.RegisterController{}),
		),
	)
	beego.AddNamespace(ns)
}
