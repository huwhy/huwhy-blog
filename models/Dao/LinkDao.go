package dao

import (
	"demo/models"
	"github.com/astaxie/beego/orm"
)

type LinkDao struct {
	Dao orm.Ormer
}

func (this *LinkDao) Find(uid int, linkType string) []models.Link {
	if this.Dao == nil {
		this.Dao = orm.NewOrm()
	}
	var links []models.Link
	this.Dao.Raw("select `id`, `name`, `url` from link where creator=? and `type`=?", uid, linkType).QueryRows(&links)
	return links
}
