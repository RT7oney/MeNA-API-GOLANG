package controllers

import (
	"MeNA-api/models"
	"encoding/json"
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (this *TestController) Get() {
	test := models.tryModel()
	this.Data["json"] = test
	this.ServeJSON()
}
