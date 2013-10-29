package controllers

import (
	"github.com/astaxie/beego"
	"leig/models"
	"strconv"
)

type EditarticleController struct {
	beego.Controller
}

func (this *EditarticleController) Post() {
	this.Data["Article"], _ = models.GetAllArticles()
	Title := this.Input().Get("ArticleTitle")
	Content := this.Input().Get("ArticleContent")
	ArticleSortId := this.Input().Get("ArticleSort")
	id, _ := strconv.ParseInt(ArticleSortId, 10, 64)
	models.AddArticle(Title, Content, id)
	this.Redirect("/admin?options=1", 301)
	return
}
