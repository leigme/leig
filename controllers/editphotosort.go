package controllers

import (
	"github.com/astaxie/beego"
	"leig/models"
	"strconv"
)

type EditphotosortController struct {
	beego.Controller
}

func (this *EditphotosortController) Post() {
	name := this.Input().Get("name")
	if len(name) == 0 {
		this.Redirect("/admin?options=3", 301)
		return
	}
	err := models.AddPhotoSort(name)
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/admin?options=3", 301)
	return
}

func (this *EditphotosortController) Get() {
	id := this.Input().Get("id")
	if len(id) == 0 {
		return
	}
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {

	}
	err = models.DelPhotoSort(cid)
	this.Redirect("/admin?options=3", 301)
}
