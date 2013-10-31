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
	Views       int64
	ArticleSort string
}

func (this *MainController) Get() {
	this.Data["Islogin"] = checkAccount(this.Ctx)
	this.Data["PhotoSort"], _ = models.GetAllPhotoSorts()
	this.Data["Photos"], _ = models.GetAllPhotos()
	this.Data["ArticleSort"], _ = models.GetAllArticleSorts()
	aas, _ := models.GetAllArticles()
	mas := make(map[int]ArticleString)
	var as ArticleString
	var i int
	for _, aa := range aas {
		as.Id = aa.Id
		as.Title = aa.Title
		as.Content = aa.Content
		as.Created = aa.Created.Format("2006-01-02 15:04:05")
		as.Updated = aa.Updated.Format("2006-01-02 15:04:05")
		as.Views = aa.Views
		gas, _ := models.GetArticleSort(aa.ArticleSortId)
		as.ArticleSort = gas.Title
		mas[i] = as
		i++
	}
	this.Data["Articles"] = mas
	this.Data["Article"], _ = models.GetIndexArticle()
	this.TplNames = "index.html"
}
