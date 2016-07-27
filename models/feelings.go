package controllers

import (
	"github.com/astaxie/beego"
)

type OneController struct {
	beego.Controller
}

type TwoController struct {
	beego.Controller
}

type ThreeController struct {
	beego.Controller
}

type FourController struct {
	beego.Controller
}

type FiveController struct {
	beego.Controller
}

type SixController struct {
	beego.Controller
}

type SevenController struct {
	beego.Controller
}

type EightController struct {
	beego.Controller
}

type NineController struct {
	beego.Controller
}

func (c *MainController) Get() { //重载了一下他的原来的get方法
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// func (c *OneController) Get() {
// 	fmt.Println("hello Ryan")
// }
