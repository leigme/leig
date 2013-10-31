package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"leig/models"
	"strconv"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Get() {
	object := this.Input().Get("ob")
	operating := this.Input().Get("op")
	id, _ := strconv.ParseInt(this.Input().Get("id"), 10, 64)
	if object == "" {
		fmt.Println("没有获取管理对象！")
		this.Redirect("/", 302)
		return
	} else {
		switch object {
		case "blog":
			this.Data["IsBlog"] = true
			this.Data["ArticleSorts"], _ = models.GetAllArticleSorts()
			this.Data["Article"], _ = models.GetAllArticles()
			if operating == "" {
				this.Data["IsEditArticle"] = true
				this.TplNames = "admin.html"
				fmt.Println("切换到指定管理页面！")
			} else {
				switch operating {
				case "add":
					this.Data["IsAddArticle"] = true
					this.TplNames = "admin.html"
				case "del":
					err := models.DelArticle(id)
					if err == nil {
						err.Error()
					}
				case "update":
				}
				this.TplNames = "admin.html"
			}
		case "photo":
			this.Data["IsPhoto"] = true
			this.Data["PhotoSort"], _ = models.GetAllPhotoSorts()
			this.Data["Photo"], _ = models.GetAllPhotos()
			if operating == "" {
				fmt.Println("切换到指定管理页面！")
			} else {
				switch operating {
				case "add":

					this.TplNames = "admin.html"

				case "del":
					err := models.DelPhoto(id)
					if err == nil {
						err.Error()
					}
				case "update":
				}
				this.TplNames = "admin.html"
			}
		case "sort":
			this.Data["IsSort"] = true
			this.Data["ArticleSorts"], _ = models.GetAllArticleSorts()
			this.Data["PhotoSorts"], _ = models.GetAllPhotoSorts()
			this.TplNames = "admin.html"
		case "pwd":
			this.Data["IsPwd"] = true
		}
		this.TplNames = "admin.html"

	}
}

func (this *AdminController) AddArticle() {
	Title := this.Input().Get("ArticleTitle")
	Content := this.Input().Get("ArticleContent")
	ArticleSortId := this.Input().Get("ArticleSort")
	id, _ := strconv.ParseInt(ArticleSortId, 10, 64)
	if Title != "" && Content != "" && id != 0 {
		models.AddArticle(Title, Content, id)
	}
	this.Redirect("/admin?ob=blog", 302)
	return
}

func (this *AdminController) AddArticleSort() {
	fmt.Println("自动路由调用AddArticleSort方法")
	this.Redirect("/admin?ob=sort", 302)
	return
}

func (this *AdminController) AddPhoto() {
	fmt.Println("自动路由调用AddPhotot方法")
	this.Redirect("/admin?ob=sort", 302)
	return
}

func (this *AdminController) AddPhotoSort() {
	fmt.Println("自动路由调用AddPhotoSort方法")
}

func (this *AdminController) UpdateArticle() {
	fmt.Println("自动路由调用UpdateArticle方法")
}

func (this *AdminController) UpdateArticleSort() {
	fmt.Println("自动路由调用UpdateArticleSort方法")
}

func (this *AdminController) UpdatePhoto() {
	fmt.Println("自动路由调用UpdatePhoto方法")
}

func (this *AdminController) UpdatePhotoSort() {
	fmt.Println("自动路由调用UpdatePhotoSort方法")
}
