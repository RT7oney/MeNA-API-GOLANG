// @APIVersion 1.0.0
// @Title MeNA-Api 文档
// @Description 一个想要让世界有一点点不一样的接口
// @Contact RyanTyler@gmail.com
package routers

import (
	"MeNA-Api/common"
	"MeNA-Api/controllers"
	"MeNA-Api/controllers/Admin"
	"MeNA-Api/models"
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	// "net/http"
)

// var r *http.Request //加了星号会爆出一个恐慌
var check = func(ctx *context.Context) {
	form := ctx.Request.Form
	if form["api_token"] == nil || form["sign"] == nil || form["timestamp"] == nil {
		ctx.WriteString("请求不合法")
	} else {
		check := make(map[string]string)
		api_token := form["api_token"][0]
		check["sign"] = form["sign"][0]
		check["timestamp"] = form["timestamp"][0]
		app_key, err := models.GetAppKey(api_token)
		if err != nil {
			ctx.WriteString("没有查询到开发者信息") //查询出错
		} else {
			ret := common.CheckSign(check, app_key)
			switch ret {
			// case 1:
			// 	ctx.WriteString("签名检查成功")
			case 0:
				ctx.WriteString("签名不匹配")
			case -1:
				ctx.WriteString("app_key过期")
			}
		}
	}
}
var auth = func(ctx *context.Context) {
	headers := ctx.Request.Header
	// fmt.Println(headers)
	if headers["AdminRT7oney"] == nil {
		// msg = common.Response(400, "您没有权限")
		// ctx.Redirect(302, "/error")
		ctx.WriteString("您没有管理员权限")
	}
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
	// fmt.Println(headers)
	// if headers["Secret-Key"] == nil {
	// 	// msg = common.Response(400, "您没有权限")
	// 	// ctx.Redirect(302, "/error")
	// 	ctx.WriteString("您没有接口调用权限")
	// }
}

func init() {
	// beego.InsertFilter("/*", beego.BeforeRouter, checkFilter)
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/dev",
			beego.NSInclude(
				&controllers.DevController{},
			),
			// beego.NSRouter("/register", &controllers.RegisterController{}),
		),
		beego.NSNamespace("/user",
			beego.NSBefore(check),
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/character",
			beego.NSBefore(check),
			beego.NSInclude(
				&controllers.CharacterController{},
			),
		),
		beego.NSNamespace("/movie",
			beego.NSBefore(check),
			beego.NSInclude(
				&controllers.MovieController{},
			),
		),
		beego.NSNamespace("/music",
			beego.NSBefore(check),
			beego.NSInclude(
				&controllers.MusicController{},
			),
		),
	)
	//管理员权限
	admin_ns := beego.NewNamespace("/admin",
		beego.NSBefore(auth),
		beego.NSInclude(
			&admin.AdminController{},
		),
	)
	beego.AddNamespace(ns, admin_ns)
}
