package controllers

import (
	"github.com/astaxie/beego"
	"leig/models"
	"strconv"
)

type EditphotoController struct {
	beego.Controller
}

func (this *EditphotoController) Post() {
	phototitle := this.Input().Get("PhotoTitle")
	photourl := this.Input().Get("PhotoUrl")
	photosortid := this.Input().Get("PhotoSortId")
	sid, _ := strconv.ParseInt(photosortid, 10, 64)
	models.AddPhoto(phototitle, photourl, sid)
	models.UpdataPhotoCount(sid)
	this.Redirect("/admin?options=2", 301)
}

func (this *EditphotoController) Get() {

}
