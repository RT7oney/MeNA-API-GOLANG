package controllers

import (
	"MeNA-Api/common"
	"MeNA-Api/models"
	// "encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	// "net"
	// "strconv"
	// "net/http"
	// "log"
	// "reflect"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title User.Register
// @Description 用户注册
// @Param   email   query   string    true   "用户邮箱"
// @Param   password    query   string  true  "用户设置密码"
// @Param   name    query   string   true   "用户设置姓名"
// @Success 200 成功
// @Failure 401 参数不完整
// @Failure 402 邮箱有误
// @router /register [post]
func (this *UserController) Register() {
	var msg string
	var data = make(map[string]string)
	data["email"] = this.GetString("email")
	data["password"] = this.GetString("password")
	data["name"] = this.GetString("name")
	api_token := this.GetString("api_token")
	if data["email"] == "" || data["password"] == "" || data["name"] == "" {
		msg = common.Response(401, "参数不完整，请检查参数", nil)
	} else {
		check := common.IsEmail(data["email"])
		if check {
			app_key, _ := models.GetAppKey(api_token)
			url := beego.AppConfig.String("Component::UserRegister")
			res := common.Request(url, data, app_key["app_key"])
			if res["code"] != "200" {
				// fmt.Println(res)
				msg = common.Response(500, res["msg"], nil)
			} else {
				msg = common.Response(200, "成功", res["msg"])
			}
		} else {
			msg = common.Response(402, "邮箱有误", nil)
		}
	}
	// ret, _ := common.EncodeData([]byte(msg))
	// common.WriteLog("加密消息@User.Register.EncodeData:" + string(msg))
	// if ret == nil {
	// 	msg = common.Response(505, "服务器内部错误")
	// 	this.Ctx.WriteString(msg)
	// }
	this.Ctx.WriteString(string(msg))
}

// @Title User.Login
// @Description 用户登录
// @Param   email   query   string    true   "用户邮箱"
// @Param   password    query   string  true  "用户密码"
// @Success 200 成功
// @Failure 401 参数不完整
// @Failure 402 邮箱有误
// @router /login [post]
func (this *UserController) Login() {
	var msg string
	var data = make(map[string]string)
	data["email"] = this.GetString("email")
	data["password"] = this.GetString("password")
	api_token := this.GetString("api_token")
	if data["email"] == "" || data["password"] == "" {
		msg = common.Response(401, "参数不完整，请检查参数", nil)
	} else {
		check := common.IsEmail(data["email"])
		if check {
			app_key, _ := models.GetAppKey(api_token)
			url := beego.AppConfig.String("Component::UserLogin")
			res := common.Request(url, data, app_key["app_key"])
			if res["code"] != "200" {
				// fmt.Println(res)
				msg = common.Response(500, res["msg"], nil)
			} else {
				msg = common.Response(200, "成功", res["msg"])
			}
		} else {
			msg = common.Response(402, "邮箱有误", nil)
		}
	}
	// ret, _ := common.EncodeData([]byte(msg))
	// common.WriteLog("加密消息@User.Register.EncodeData:" + string(msg))
	// if ret == nil {
	// 	msg = common.Response(505, "服务器内部错误")
	// 	this.Ctx.WriteString(msg)
	// }
	this.Ctx.WriteString(string(msg))
}
