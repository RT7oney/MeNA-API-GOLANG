package controllers

import (
	"MeNA-Api/common"
	"MeNA-Api/models"
	"encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	"strconv"
	// "net"
	// "net/http"
	// "log"
	// "reflect"
)

// Operations about Users
type CharacterController struct {
	beego.Controller
}

// @router /getone [post]
func (this *CharacterController) GetOne() {
	var msg string
	var data = make(map[string]string)
	id := this.GetString("id")
	api_token := this.GetString("api_token")
	if id == "" {
		msg = common.Response(401, "参数不完整，请检查参数", nil)
	} else {
		character_id, _ := strconv.ParseInt(id, 10, 64)
		ret, err := models.GetOneCharacter(character_id)
		if err != nil {
			msg = common.Response(500, "查询数据库出错："+err.Error(), nil)
		} else {
			// fmt.Println(ret)
			app_key, _ := models.GetAppKey(api_token)
			key := strconv.FormatInt(ret.Id, 10)
			data[key] = ret.Introduction
			jsonbyte, _ := json.Marshal(data)
			encode_data, err := common.EncodeData(jsonbyte, app_key["app_key"])
			if err != nil {
				msg = common.Response(200, "查询错误："+err.Error(), nil)
			} else {
				msg = common.Response(200, "成功", string(encode_data))
			}
		}
	}
	this.Ctx.WriteString(msg)
}

// @router /getall [post]
func (this *CharacterController) GetAll() {
	var msg string
	//todo
	this.Ctx.WriteString(msg)
}
