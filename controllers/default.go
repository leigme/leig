package controllers

import (
	"github.com/astaxie/beego"
	"leig/models"
)

type MainController struct {
	beego.Controller
}

type ArticleString struct {
	Id          int64
	Title       string
	Content     string
	Created     string
	Updated     string
	ArticleSort string
}

func (this *MainController) Get() {
	this.Data["Islogin"] = checkAccount(this.Ctx)
	this.Data["PhotoSort"], _ = models.GetAllPhotoSorts()
	this.Data["Photos"], _ = models.GetAllPhotos()
	this.Data["ArticleSort"], _ = models.GetAllArticleSorts()
	this.Data["Articles"], _ = models.GetAllArticles()
	this.Data["Article"], _ = models.GetIndexArticle()
	this.TplNames = "index.html"
}
