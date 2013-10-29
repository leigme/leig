package controllers

import (
	"github.com/astaxie/beego"
	"leig/models"
	"strconv"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	var isEdit string
	var i int64
	isEdit = this.Input().Get("options")
	i, _ = strconv.ParseInt(isEdit, 10, 64)
	switch i {
	case 1:
		this.Data["ArticleSort"], _ = models.GetAllArticleSorts()
		this.Data["Article"], _ = models.GetAllArticles()
		this.Data["IsBlog"] = true
	case 2:
		this.Data["PhotoSort"], _ = models.GetAllPhotoSorts()
		this.Data["Photo"], _ = models.GetAllPhotos()
		this.Data["IsPhoto"] = true
	case 3:
		this.Data["IsSort"] = true
		this.Data["ArticleSort"], _ = models.GetAllArticleSorts()
		this.Data["PhotoSort"], _ = models.GetAllPhotoSorts()
	case 4:
		this.Data["IsPwd"] = true
	}
	this.TplNames = "admin.html"
	return
}
