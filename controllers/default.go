package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Islogin"] = checkAccount(this.Ctx)
	this.TplNames = "index.html"
}
