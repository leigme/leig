package controllers

import (
	"github.com/astaxie/beego"
	"leig/models"
	"strconv"
)

type EditblogsortController struct {
	beego.Controller
}

func (this *EditblogsortController) Post() {
	name := this.Input().Get("name")
	if len(name) == 0 {
		this.Redirect("/admin?options=3", 301)
		return
	}
	err := models.AddArticleSort(name)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/admin?options=3", 301)
	return
}

func (this *EditblogsortController) Get() {
	id := this.Input().Get("id")
	if len(id) == 0 {
		return
	}
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {

	}
	err = models.DelArticleSort(cid)
	this.Redirect("/admin?options=3", 301)
}
