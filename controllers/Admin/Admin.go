package admin

import (
	"MeNA-Api/common"
	// "encoding/json"
	"MeNA-Api/models"
	// "fmt"
	"github.com/astaxie/beego"
	"strconv"
	// "net"
	// "net/http"
	// "log"
	// "reflect"
)

// Operations about Users
type AdminController struct {
	beego.Controller
}

// @router /add-character [post]
func (this *AdminController) AddCharacter() {
	var msg string
	// var data = make(map[string]string)
	intro := this.GetString("intro")
	if intro == "" {
		msg = common.Response(401, "参数不完整，请检查参数", nil)
	} else {
		ret, err := models.AddCharacter(intro)
		if err != nil {
			msg = common.Response(500, "插入数据库出错："+err.Error(), nil)
		} else if ret == -1 {
			msg = common.Response(501, "九型人格已经添加完成，不能继续添加了", nil)
		} else {
			// fmt.Println(ret)
			id := strconv.FormatInt(ret, 10)
			msg = common.Response(200, "插入第"+id+"型人格成功", nil)
		}
	}
	this.Ctx.WriteString(msg)
}

// @router /add-movie_tag [post]
func (this *AdminController) AddMovieTag() {
	var msg string
	// var data = make(map[string]string)
	tag_name := this.GetString("tag_name")
	if tag_name == "" {
		msg = common.Response(401, "参数不完整，请检查参数", nil)
	} else {
		ret, err := models.AddMovieTag(tag_name)
		if err != nil {
			msg = common.Response(500, "插入数据库出错："+err.Error(), nil)
		} else if ret == -1 {
			msg = common.Response(501, "已经存在这个标签了", nil)
		} else {
			// fmt.Println(ret)
			id := strconv.FormatInt(ret, 10)
			msg = common.Response(200, "插入第"+id+"个电影标签`"+tag_name+"`成功", nil)
		}
	}
	this.Ctx.WriteString(msg)
}

// @router /add-music_tag [post]
func (this *AdminController) AddMusicTag() {
	var msg string
	// var data = make(map[string]string)
	tag_name := this.GetString("tag_name")
	if tag_name == "" {
		msg = common.Response(401, "参数不完整，请检查参数", nil)
	} else {
		ret, err := models.AddMusicTag(tag_name)
		if err != nil {
			msg = common.Response(500, "插入数据库出错："+err.Error(), nil)
		} else if ret == -1 {
			msg = common.Response(501, "已经存在这个标签了", nil)
		} else {
			// fmt.Println(ret)
			id := strconv.FormatInt(ret, 10)
			msg = common.Response(200, "插入第"+id+"个音乐标签`"+tag_name+"`成功", nil)
		}
	}
	this.Ctx.WriteString(msg)
}

// @router /add-app_key [post]
func (this *AdminController) AddDev() {
	var msg string
	content := make(map[string]string)
	// var data = make(map[string]string)
	// TODO 之后调用组件然后返回的时候添加一个app_key
	content["api_token"] = this.GetString("api_token")
	content["dev_name"] = this.GetString("dev_name")

	if content["api_token"] == "" || content["dev_name"] == "" {
		msg = common.Response(401, "参数不完整，请检查参数", nil)
	} else {
		ret, err := models.AddAppKey(content)
		if err != nil {
			msg = common.Response(500, "插入数据库出错："+err.Error(), nil)
		} else if ret == -1 {
			msg = common.Response(501, "已经存在该开发者", nil)
		} else {
			// fmt.Println(ret)
			id := strconv.FormatInt(ret, 10)
			msg = common.Response(200, "插入第"+id+"个开发者信息成功", nil)
		}
	}
	this.Ctx.WriteString(msg)
}
