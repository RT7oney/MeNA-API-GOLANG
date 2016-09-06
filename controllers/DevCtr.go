package controllers

import (
	"MeNA-Api/common"
	"MeNA-Api/models"
	"encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
	// "net"
	// "net/http"
	// "log"
	// "reflect"
)

// Operations about Users
type DevController struct {
	beego.Controller
}

// @Title Dev.GetAppKey
// @Description 开发者获取appkey
// @Param   api_token    query   string   true   "开发者在注册的时候获取的api_token"
// @Success 200 成功
// @Failure 401 参数不完整
// @Failure 403 不存在该api_token
// @router /get-app-key [post]
func (this *DevController) GetAppKey() {
	var msg string
	var data = make(map[string]string)
	data["api_token"] = this.GetString("api_token")
	if data["api_token"] == "" {
		msg = common.Response(401, "参数不完整，请检查参数", nil)
	} else {
		app_key, err := models.GetAppKey(data["api_token"])
		if err != nil || app_key == nil {
			msg = common.Response(403, "查询失败，api_token不存在", nil)
		} else {
			timenow := time.Now().Unix()
			expire_time, _ := strconv.ParseInt(app_key["expire_time"], 10, 64)
			if expire_time < timenow {
				//app_key过期，去update一个新的出来
				// fmt.Println("【@#==========是update的==========#@】")
				app_key, err = models.UpdateAppKey(data["api_token"])
			}
			if err != nil {
				msg = common.Response(500, "服务器内部错误", nil)
			} else {
				// jsonbyte, _ := json.Marshal(app_key)
				// encode_data, _ := common.EncodeData(jsonbyte)
				msg = common.Response(200, "成功", app_key)
			}
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

// @Title Dev.Register
// @Description 开发者注册
// @Param   the_id    query   string   true   "通过系统获取"
// @Param   password    query   string   true   "需要开发者输入密码验证身份"
// @Success 200 成功
// @Failure 401 参数不完整
// @router /register [post]
func (this *DevController) Register() {
	var msg string
	var data = make(map[string]string)
	data["the_id"] = this.GetString("api_token")
	data["password"] = this.GetString("password")
	if data["the_id"] == "" || data["password"] == "" {
		msg = common.Response(401, "参数不完整，请检查参数", nil)
	} else {
		url := beego.AppConfig.String("Component::DevRegister")
		res := common.Request(url, data, "")
		// fmt.Println("===res===")
		// fmt.Println(res)
		if res["code"] != "200" {
			// fmt.Println(res)
			msg = common.Response(500, res["msg"], nil)
		} else {
			// fmt.Println("===数据===")
			// fmt.Println(res["msg"])
			type Msg struct {
				Code string
				Msg  string
				Data map[string]string
			}
			var m Msg
			err := json.Unmarshal([]byte(res["msg"]), &m)
			if err != nil {
				// fmt.Println("===错误===")
				// fmt.Println(err)
				msg = common.Response(501, "服务器内部错误", nil)
			} else {
				if m.Code != "10003.201" {
					msg = common.Response(502, m.Msg, nil)
				} else {
					ret, err := models.AddAppKey(m.Data)
					if err != nil {
						msg = common.Response(503, "插入数据库出错："+err.Error(), nil)
					} else if ret == -1 {
						msg = common.Response(504, "已经存在该开发者", nil)
					} else {
						// fmt.Println(ret)
						msg = common.Response(200, "恭喜您成功成为开发者！", nil)
					}
				}
				/**********涨姿势***********/
				// fmt.Println("===json 解析===")
				// t := reflect.TypeOf(m.Data)
				// v := reflect.ValueOf(m.Data)
				// fmt.Println(m)
				// fmt.Println(m.Code)
				// switch m.Data.(type) {
				// case string:
				// 	fmt.Println("yeah")
				// default:
				// 	fmt.Println("······")
				// }
				// fmt.Println("===类型===")
				// fmt.Println(t)
				// fmt.Println("===值===")
				// fmt.Println(m.Data)
				//

				//
				// for i := 0; i < t.NumField(); i++ {
				// 	f := t.Field(i)
				// 	val := v.Field(i).Interface()
				// 	fmt.Println(f)
				// 	fmt.Println(val)
				// }
				//

				//
				// switch v := m.(type) {
				// default:
				// 	fmt.Println("===switch===")
				// 	fmt.Println(reflect.TypeOf(v))
				// } // 接口的类型选择
				/***************************/
			}
		}
	}
	this.Ctx.WriteString(string(msg))
}
