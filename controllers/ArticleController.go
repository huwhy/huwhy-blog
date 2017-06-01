package controllers

import (
	"demo/models/Dao"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Get() {
	o := new(dao.ArticleDao)
	id, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	article := o.Get(id)
	this.Data["article"] = article
	this.Data["tags"] = strings.Split(article.Tags, " ")
	this.TplName = "article.html"
	this.Layout = "layout/layout.html"
}
