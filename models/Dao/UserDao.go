package dao

import (
	"demo/models"
	"github.com/astaxie/beego/orm"
)

type UserDao struct {
	Dao orm.Ormer
}

func (this *UserDao) Get(id int) models.User {
	if this.Dao == nil {
		this.Dao = orm.NewOrm()
	}
	var user models.User
	this.Dao.Raw("select id, username, password from users where id=?", id).QueryRow(&user)
	return user
}
