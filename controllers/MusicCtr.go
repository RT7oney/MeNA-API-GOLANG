package controllers

import (
	// "MeNA-Api/common"
	// "encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	// "net"
	// "net/http"
	// "log"
	// "reflect"
)

// Operations about Musics
type MusicController struct {
	beego.Controller
}

// @Title Music.GetAllTags
// @Description 获取所有电影的标签
// @Success 200 成功
// @Failure 401 参数不完整
// @Failure 402 邮箱有误
// @router /get-all-tags [post]
func (this *MusicController) GetAllTags() {
	var msg string

	this.Ctx.WriteString(string(msg))
}
