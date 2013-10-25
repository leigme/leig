package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	var isEdit string
	var i int64
	isEdit = this.Input().Get("options")
	i,_ = strconv.ParseInt(isEdit, 10, 64)
	switch i {
	case 1:this.Data["IsBlog"]=true
	case 2:this.Data["IsPhoto"]=true
	case 3:this.Data["IsSort"]=true
	case 4:this.Data["IsPwd"]=true
	}
	this.TplNames="admin.html"
	return
}