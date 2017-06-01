package controllers

import "github.com/astaxie/beego"

type ItemController struct {
	beego.Controller
}

func (this *ItemController) Get() {
	id := this.Ctx.Input.Param(":id")
	this.Data["id"] = id
	this.TplName = "item.html"
}
