package main

import (
	"demo/controllers"
	"demo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:abc123@/blog?charset=utf8", 30)
	orm.RegisterModel(new(models.User), new(models.Article), new(models.Seo), new(models.Category), new(models.Link))
	//orm.RunSyncdb("default", false, true)

}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello, world")
}

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	index := controllers.IndexController{}
	beego.Router("/", &index)
	beego.Router("/article/:id([0-9]+).html", &controllers.ArticleController{})
	beego.Router("/index", &index)
	beego.Router("/hao123.html", &controllers.LinkController{})
	beego.Router("/shop.html", &controllers.ShopController{})
	beego.Router("/item/:id([0-9]+).html", &controllers.ItemController{})
	beego.SetStaticPath("/static", "static")
	beego.AddFuncMap("ico", controllers.GetIco)
	beego.Run()
}
