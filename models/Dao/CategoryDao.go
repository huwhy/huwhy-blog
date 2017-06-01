package dao

import (
	"demo/models"
	"github.com/astaxie/beego/orm"
)

type CategoryDao struct {
	Dao orm.Ormer
}

func (this *CategoryDao) Get(id int) models.Category {
	if this.Dao == nil {
		this.Dao = orm.NewOrm()
	}
	var category models.Category
	this.Dao.Raw(" SELECT id, pid, name, parent, level, updated, created from article_catalog where id=?", id).QueryRow(&category)
	return category
}

func (this *CategoryDao) FindFirst() []models.Category {
	if this.Dao == nil {
		this.Dao = orm.NewOrm()
	}
	var categories []models.Category
	this.Dao.Raw(" SELECT id, pid, name, parent, level, updated, created from article_catalog where level=?", 1).QueryRows(&categories)
	return categories
}

func (this *CategoryDao) FindChild(id int) []models.Category {
	if this.Dao == nil {
		this.Dao = orm.NewOrm()
	}
	var categories []models.Category
	this.Dao.Raw("select id, pid, name, parent, level, updated, created from article_catalog where pid=?", id).QueryRows(&categories)
	return categories
}
