package controllers

import (
	"MeNA-Api/common"
	// "encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	// "net"
	// "net/http"
	// "log"
	// "reflect"
)

// Operations about Users
type RegisterController struct {
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
// @router /User.Register [post]
func (this *RegisterController) Post() {
	var ret string
	var data = make(map[string]string)
	data["email"] = this.GetString("email")
	data["password"] = this.GetString("password")
	data["name"] = this.GetString("name")
	if data["email"] == "" || data["password"] == "" || data["name"] == "" {
		ret = common.Response(401, "参数不完整，请检查参数")
	} else {
		check := common.IsEmail(data["email"])
		if check {
			url := beego.AppConfig.String("UserRegister")
			ret = common.Request(url, data)
		} else {
			ret = common.Response(402, "邮箱有误")
		}
	}
	// msg, _ := common.EncodeData([]byte(ret))
	// common.WriteLog("加密消息@User.Register.EncodeData:" + string(msg))
	// if msg == nil {
	// 	ret = common.Response(505, "服务器内部错误")
	// 	this.Ctx.WriteString(ret)
	// }
	this.Ctx.WriteString(string(ret))
}
