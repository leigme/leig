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
			mas := make(map[int]ArticleString)
			var as ArticleString
			var i int
			for _, aa := range article {
				as.Id = aa.Id
				as.Title = aa.Title
				as.Content = aa.Content
				as.Created = aa.Created.Format("2006-01-02 15:04:05")
				as.Updated = aa.Updated.Format("2006-01-02 15:04:05")
				as.Views = aa.Views
				gas, _ := models.GetArticleSort(aa.ArticleSortId)
				as.ArticleSort = gas.Title
				mas[i] = as
				i++
			}
			this.Data["Article"] = mas
			if operating == "" {
				this.Data["IsEditArticle"] = true
				this.TplNames = "admin.html"
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
					this.Redirect("/admin?ob=blog", 302)
					return
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
				this.TplNames = "admin.html"
			}
		case "photo":
			this.Data["IsPhoto"] = true
			photosort, errs := models.GetAllPhotoSorts()
			photo, errp := models.GetAllPhotos()
			if errs != nil {
				errs.Error()
				this.Redirect("/", 302)
				return
			}
			if errp != nil {
				errp.Error()
				this.Redirect("/", 302)
				return
			}
			this.Data["PhotoSort"] = photosort
			this.Data["Photo"] = photo
			if operating == "" {
				this.Data["IsEditPhoto"] = true
				this.TplNames = "admin.html"
			} else {
				switch operating {
				case "add":

				case "del":
					Pid := this.Input().Get("id")
					id, _ := strconv.ParseInt(Pid, 10, 64)
					err := models.DelPhoto(id)
					if err != nil {
						err.Error()
						this.Redirect("/", 302)
						return
					}

				case "update":
				}

			}
		case "sort":
			this.Data["IsSort"] = true
			articlesorts, erra := models.GetAllArticleSorts()
			photosorts, errp := models.GetAllPhotoSorts()
			if erra != nil {
				erra.Error()
				this.Redirect("/", 302)
				return
			}
			if errp != nil {
				errp.Error()
				this.Redirect("/", 302)
				return
			}
			this.Data["ArticleSorts"] = articlesorts
			this.Data["PhotoSorts"] = photosorts
			op := this.Input().Get("op")
			if op == "" {
				this.TplNames = "admin.html"
			}
			switch op {
			case "delarticle":
				aid := this.Input().Get("id")
				if aid == "" {
					this.Redirect("/", 302)
					return
				}
				id, _ := strconv.ParseInt(aid, 10, 64)
				err := models.DelArticleSort(id)
				if err != nil {
					this.Redirect("/", 302)
					return
				}
				this.Redirect("/admin?ob=sort", 302)
				return
			case "delphoto":
				pid := this.Input().Get("id")
				if pid == "" {
					this.Redirect("/", 302)
					return
				}
				id, _ := strconv.ParseInt(pid, 10, 64)
				err := models.DelPhotoSort(id)
				if err != nil {
					this.Redirect("/", 302)
				}
				this.Redirect("/admin?ob=sort", 302)
				return
			}
			this.TplNames = "admin.html"
		case "pwd":
			this.Data["IsPwd"] = true
		}

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
		artilcesorts, errs := models.GetAllArticleSorts()
		article, erra := models.GetAllArticles()
		if errs != nil {
			errs.Error()
			this.Redirect("/", 302)
			return
		}
		if erra != nil {
			erra.Error()
			this.Redirect("/", 302)
			return
		}
		this.Data["ArticleSorts"] = artilcesorts
		this.Data["Article"] = article
		this.TplNames = "admin.html"
	}
	this.Redirect("/admin?ob=blog", 302)
	return
}

func (this *AdminController) ActionAddArticleSort() {
	Title := this.Input().Get("articlesorttitle")
	if Title == "" {
		this.Redirect("/admin?ob=sort", 302)
		return
	}
	err := models.AddArticleSort(Title)
	if err != nil {
		err.Error()
		this.Redirect("/", 302)
		return
	}
	this.Data["IsSort"] = true
	articlesorts, erra := models.GetAllArticleSorts()
	photosorts, errp := models.GetAllPhotoSorts()
	if erra != nil {
		erra.Error()
		this.Redirect("/", 302)
		return
	}
	if errp != nil {
		errp.Error()
		this.Redirect("/", 302)
		return
	}
	this.Data["ArticleSorts"] = articlesorts
	this.Data["PhotoSorts"] = photosorts
	this.TplNames = "admin.html"
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
		photosort, errs := models.GetAllPhotoSorts()
		photo, errp := models.GetAllPhotos()
		if errs != nil {
			errs.Error()
			this.Redirect("/", 302)
			return
		}
		if errp != nil {
			errp.Error()
			this.Redirect("/", 302)
			return
		}
		this.Data["PhotoSort"] = photosort
		this.Data["Photo"] = photo
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
		articlesorts, erra := models.GetAllArticleSorts()
		photosorts, errp := models.GetAllPhotoSorts()
		if erra != nil {
			erra.Error()
			this.Redirect("/", 302)
			return
		}
		if errp != nil {
			errp.Error()
			this.Redirect("/", 302)
			return
		}
		this.Data["ArticleSorts"] = articlesorts
		this.Data["PhotoSorts"] = photosorts
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
