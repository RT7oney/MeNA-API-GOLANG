package main

import (
	"fmt"
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Ctx.WriteString("hello ryan")
}

func main() {
	beego.Router("/hello", &HomeController{})
	beego.Run()
}
