package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"leig/models"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	isExit := this.Input().Get("exit") == "true"
	if isExit {
		this.Ctx.SetCookie("uname", "", -1, "/")
		this.Ctx.SetCookie("pwd", "", -1, "/")
		this.Redirect("/", 302)
		return
	} else {
		this.Data["IsBlog"] = true
		this.Data["IsEditArticle"] = true
		this.Data["ArticleSorts"], _ = models.GetAllArticleSorts()
		this.Data["Article"], _ = models.GetAllArticles()
		this.TplNames = "admin.html"
		return
	}
}

func (this *LoginController) Post() {
	uname := this.Input().Get("name")
	pwd := this.Input().Get("pwd")
	aotulogin := this.Input().Get("aotulogin") == "on"
	adminpwd := models.Estimate(uname)
	if "nil" == adminpwd {
		this.Redirect("/", 302)
	} else {
		if pwd == adminpwd {
			maxAge := 0
			if aotulogin {
				maxAge = 1<<31 - 1
			}
			this.Ctx.SetCookie("uname", uname, maxAge, "/")
			this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
			this.Data["IsBlog"] = true
			this.Data["IsEditArticle"] = true
			this.Data["ArticleSorts"], _ = models.GetAllArticleSorts()
			this.Data["Article"], _ = models.GetAllArticles()
			this.TplNames = "admin.html"
			return
		}
		this.Redirect("/", 302)
		return
	}
}

func checkAccount(Ctx *context.Context) bool {
	ck, err := Ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}
	uname := ck.Value
	ck, err = Ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	pwd := ck.Value
	return pwd == models.Estimate(uname)
}
