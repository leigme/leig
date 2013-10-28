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
	id, _ := strconv.ParseInt(photosortid, 10, 64)
	ps, _ := models.GetPhotoSort(id)
	models.AddPhoto(phototitle, photourl, &ps)
	models.UpdataPhotoCount(id)
	this.Redirect("/admin?options=2", 301)
}

func (this *EditphotoController) Get() {

}
