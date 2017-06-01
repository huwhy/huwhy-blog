package dao

import (
	"demo/models"
	"github.com/astaxie/beego/orm"
)

type ArticleDao struct {
	Dao orm.Ormer
}

const (
	DISPLAY = "display"
	HIDE    = "hide"
)

type ArticleTerm struct {
	FirstId  int
	SecondId int
	ThirdId  int
	Status   string
}

func (this *ArticleDao) Get(id int) models.Article {
	if this.Dao == nil {
		this.Dao = orm.NewOrm()
	}
	var article models.Article
	this.Dao.Raw("select id, first_cid, second_cid, third_cid, title, tags, summary, "+
		"content, status, img_url, author, comments, views, url, updated, created from article where id=?", id).QueryRow(&article)
	return article
}

func (this *ArticleDao) Find(term *ArticleTerm) []models.Article {
	if this.Dao == nil {
		this.Dao = orm.NewOrm()
	}
	args := []interface{}{}
	var articles []models.Article
	sql := "select id, first_cid, second_cid, third_cid, title, tags, summary, " +
		"content, status, img_url, author, comments, views, url, updated, created from article where 1=1"
	if term.FirstId != 0 {
		sql += " and first_cid=?"
		args = append(args, term.FirstId)
	}
	if term.SecondId != 0 {
		sql += " and second_id=?"
		args = append(args, term.SecondId)
	}
	if term.ThirdId != 0 {
		sql += " and third_id=?"
		args = append(args, term.ThirdId)
	}
	if len(term.Status) != 0 {
		sql += " and status=?"
		args = append(args, term.Status)
	}

	this.Dao.Raw(sql, args).QueryRows(&articles)
	return articles
}
