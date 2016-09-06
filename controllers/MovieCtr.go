package controllers

import (
	"MeNA-Api/common"
	"MeNA-Api/models"
	"encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	// "net"
	// "net/http"
	// "log"
	// "reflect"
	"strconv"
)

// Operations about Movies
type MovieController struct {
	beego.Controller
}

// @Title Movie.GetAllTags
// @Description 获取所有电影的标签
// @Param   page   query   string    true   "页数(默认第一页)"
// @Param   count    query   string  true  "每页显示多少条(默认10条)"
// @Success 200 成功
// @Failure 401 参数不完整
// @Failure 500 查询数据库出错
// @Failure 501 服务器内部错误
// @router /get-all-tags [post]
func (this *MovieController) GetAllTags() {
	var msg string
	var data = make(map[string]string)
	data["page"] = this.GetString("page")
	data["count"] = this.GetString("count")
	api_token := this.GetString("api_token")
	if data["page"] == "" {
		data["page"] = "1"
	}
	if data["count"] == "" {
		data["count"] = "10"
	}
	page, _ := strconv.Atoi(data["page"])
	count, _ := strconv.Atoi(data["count"])
	if page < 1 || count < 1 || count > 20 {
		msg = common.Response(401, "参数中page不应该小于1，count不应该大于20小于1，请检查参数", nil)
	} else {
		limit := count
		offset := (page - 1) * count
		var opts = make(map[string]int)
		// opts["limit"] = strconv.Itoa(limit)
		// opts["offset"] = strconv.Itoa(offset)
		opts["limit"] = limit
		opts["offset"] = offset
		res, err := models.GetAllMovieTag(opts)
		if err != nil {
			msg = common.Response(500, "查询数据库出错", nil)
		} else {
			app_key, _ := models.GetAppKey(api_token)
			jsonbyte, _ := json.Marshal(res)
			// fmt.Println("===jsonbyte===")
			// fmt.Println(reflect.TypeOf(res))
			// fmt.Println(jsonbyte)
			encode_data, e := common.EncodeData(jsonbyte, app_key["app_key"])
			if e != nil {
				msg = common.Response(501, "服务器内部错误", nil)
			} else {
				msg = common.Response(200, "成功", string(encode_data))
			}
		}
	}
	this.Ctx.WriteString(string(msg))
}
