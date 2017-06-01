package controllers

import (
	"demo/models/Dao"
	"github.com/astaxie/beego"
	"regexp"
)

type LinkController struct {
	beego.Controller
}

func GetIco(url string) string {
	reg, _ := regexp.Compile("https?://([^\\.]+)\\.([^\\.]+)\\.([^\\./]{2,5})")
	rs := reg.Find([]byte(url))
	return string(rs)
}

func (this *LinkController) Get() {
	linkDao := new(dao.LinkDao)
	links := linkDao.Find(0, "navigation")
	this.Data["links"] = links
	this.TplName = "hao123.html"
	this.Layout = "layout/layout.html"
}
