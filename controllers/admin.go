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
	if object == "" {
		fmt.Println("没有获取管理对象！")
		this.Redirect("/", 302)
		return
	} else {
		switch object {
		case "blog":
			this.Data["IsBlog"] = true
			articlesorts, errs := models.GetAllArticleSorts()
			if errs != nil {
				errs.Error()
				this.Redirect("/", 302)
				return
			}
			this.Data["ArticleSorts"] = articlesorts
			article, erra := models.GetAllArticles()
			if erra != nil {
				erra.Error()
				this.Redirect("/", 302)
				return
			}
			this.Data["Article"] = article
			if operating == "" {
				this.Data["IsEditArticle"] = true
			} else {
				switch operating {
				case "add":
					this.Data["IsAddArticle"] = true
				case "del":
					Aid := this.Input().Get("id")
					id, _ := strconv.ParseInt(Aid, 10, 64)
					err := models.DelArticle(id)
					if err != nil {
						err.Error()
						this.Redirect("/", 302)
						return
					}
				case "update":
					this.Data["IsUpdate"] = true
					Aid := this.Input().Get("id")
					id, _ := strconv.ParseInt(Aid, 10, 64)
					articlesorts, err := models.GetAllArticleSorts()
					if err != nil {
						err.Error()
						this.Redirect("/", 302)
						return
					}
					this.Data["ArticleSorts"] = articlesorts
					article, erra := models.GetArticle(id)
					if erra != nil {
						erra.Error()
						this.Redirect("/", 302)
						return
					}
					this.Data["ArticleId"] = Aid
					this.Data["Title"] = article.Title
					this.Data["Content"] = article.Content
					articlesort, errs := models.GetArticleSort(article.ArticleSortId)
					if errs != nil {
						errs.Error()
						this.Redirect("/", 302)
						return
					}
					this.Data["SortId"] = articlesort.Id
					this.Data["SortTitle"] = articlesort.Title
				}
			}
		case "photo":
			this.Data["IsPhoto"] = true
			this.Data["PhotoSort"], _ = models.GetAllPhotoSorts()
			this.Data["Photo"], _ = models.GetAllPhotos()
			if operating == "" {

			} else {
				switch operating {
				case "add":

				case "del":
					Pid := this.Input().Get("id")
					id, _ := strconv.ParseInt(Pid, 10, 64)
					err := models.DelPhoto(id)
					if err == nil {

					}
					err.Error()
				case "update":
				}

			}
		case "sort":
			this.Data["IsSort"] = true
			this.Data["ArticleSorts"], _ = models.GetAllArticleSorts()
			this.Data["PhotoSorts"], _ = models.GetAllPhotoSorts()
		case "pwd":
			this.Data["IsPwd"] = true
		}
		this.TplNames = "admin.html"
	}
}

func (this *AdminController) ActionAddArticle() {
	Title := this.Input().Get("ArticleTitle")
	Content := this.Input().Get("ArticleContent")
	ArticleSortId := this.Input().Get("ArticleSort")
	if Title != "" && Content != "" && ArticleSortId != "" {
		id, _ := strconv.ParseInt(ArticleSortId, 10, 64)
		models.AddArticle(Title, Content, id)
		this.Data["IsBlog"] = true
		this.Data["ArticleSorts"], _ = models.GetAllArticleSorts()
		this.Data["Article"], _ = models.GetAllArticles()
		this.TplNames = "admin.html"
	}
	this.Redirect("/admin?ob=blog", 302)
	return
}

func (this *AdminController) ActionAddArticleSort() {
	Title := this.Input().Get("articlesorttitle")
	if Title != "" {
		models.AddArticleSort(Title)
		this.Data["IsSort"] = true
		this.Data["ArticleSorts"], _ = models.GetAllArticleSorts()
		this.Data["PhotoSorts"], _ = models.GetAllPhotoSorts()
		this.TplNames = "admin.html"
	}
	this.Redirect("/admin?ob=sort", 302)
	return
}

func (this *AdminController) ActionAddPhoto() {
	Titile := this.Input().Get("phototitle")
	Url := this.Input().Get("photourl")
	Psi := this.Input().Get("photosortid")
	if Titile != "" && Url != "" && Psi != "" {
		Sid, _ := strconv.ParseInt(Psi, 10, 64)
		models.AddPhoto(Titile, Url, Sid)
		this.Data["IsPhoto"] = true
		this.Data["PhotoSort"], _ = models.GetAllPhotoSorts()
		this.Data["Photo"], _ = models.GetAllPhotos()
		this.TplNames = "admin.html"
	}
	this.Redirect("/admin?ob=sort", 302)
	return
}

func (this *AdminController) ActionAddPhotoSort() {
	Title := this.Input().Get("photosorttitle")
	if Title != "" {
		models.AddArticleSort(Title)
		this.Data["IsSort"] = true
		this.Data["ArticleSorts"], _ = models.GetAllArticleSorts()
		this.Data["PhotoSorts"], _ = models.GetAllPhotoSorts()
		this.TplNames = "admin.html"
	}
	this.Redirect("/admin?ob=sort", 302)
	return
}

func (this *AdminController) ActionUpdateArticle() {
	Aid := this.Input().Get("ArticleId")
	Title := this.Input().Get("ArticleTitle")
	Content := this.Input().Get("ArticleContent")
	ASid := this.Input().Get("ArticleSort")
	if Aid != "" && Title != "" && Content != "" && Aid != "" {
		Sid, _ := strconv.ParseInt(ASid, 10, 64)
		Id, _ := strconv.ParseInt(Aid, 10, 64)
		models.UpdateArticle(Id, Title, Content, Sid)
	}
	this.Redirect("/admin?ob=blog", 302)
	return
}

func (this *AdminController) ActionUpdateArticleSort() {
	fmt.Println("自动路由调用UpdateArticleSort方法")
	this.Redirect("/admin", 302)
	return
}

func (this *AdminController) ActionUpdatePhoto() {
	fmt.Println("自动路由调用UpdatePhoto方法")
	this.Redirect("/admin", 302)
	return
}

func (this *AdminController) ActionUpdatePhotoSort() {
	fmt.Println("自动路由调用UpdatePhotoSort方法")
	this.Redirect("/admin", 302)
	return
}
