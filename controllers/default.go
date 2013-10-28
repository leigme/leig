package controllers

import (
	"github.com/astaxie/beego"
	"leig/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Islogin"] = checkAccount(this.Ctx)
	this.Data["Photos"], _ = models.GetAllPhotos()
	this.TplNames = "index.html"
}
