package controllers

import (
	"demo/models"
	"demo/models/Dao"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

type CategoryVo struct {
	models.Category
	Children []models.Category
}

func (this *IndexController) Get() {
	articleDao := new(dao.ArticleDao)
	articleTerm := new(dao.ArticleTerm)
	articleTerm.Status = dao.DISPLAY
	articles := articleDao.Find(articleTerm)
	this.Data["articles"] = articles
	seoDao := new(dao.SeoDao)
	seo := seoDao.Get("Home", 0)
	this.Data["seo"] = seo
	this.Layout = "layout/layout.html"
	this.TplName = "index.html"

	categoryDao := new(dao.CategoryDao)
	firstCats := categoryDao.FindFirst()
	var firstCates = []CategoryVo{}
	if len(firstCats) > 0 {
		for _, cat := range firstCats {
			vo := CategoryVo{}
			vo.Id = cat.Id
			vo.Name = cat.Name
			vo.Pid = cat.Pid
			vo.Parent = cat.Parent
			vo.Children = categoryDao.FindChild(cat.Id)
			firstCates = append(firstCates, vo)
		}
	}
	this.Data["catList"] = firstCates
}
