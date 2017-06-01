package dao

import (
	"demo/models"
	"github.com/astaxie/beego/orm"
)

type SeoDao struct {
	Dao orm.Ormer
}

func (this *SeoDao) Get(key string, id int) models.Seo {
	if this.Dao == nil {
		this.Dao = orm.NewOrm()
	}
	var seo models.Seo
	this.Dao.Raw("select * from seo where seo_key=? and target_id=?", key, id).QueryRow(&seo)
	return seo
}
