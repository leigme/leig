package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//数据库配置文件获取
func mysqlconfig() string {
	mysqlconfig := beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpwd") + "@" + "/" + beego.AppConfig.String("mysqldb") + "?charset=" + beego.AppConfig.String("charset")
	fmt.Println("读取配置文件，拼接数据库连接语句！")
	return mysqlconfig
}

//系统管理员
type User struct {
	Id       int64
	Username string
	Password string
	Nickname string
}

//文章分类
type ArticleSort struct {
	Id      int64
	Title   string
	Created time.Time `orm:"index"`
	Count   int64
}

//文章表
type Article struct {
	Id            int64
	Title         string
	Content       string    `orm:"size(5000)"`
	Created       time.Time `orm:"index"`
	Updated       time.Time `orm:"index"`
	Views         int64     `orm:"index"`
	ArticleSortId int64
}

//图片分类
type PhotoSort struct {
	Id      int64
	Title   string
	Created time.Time `orm:"index"`
	Count   int64
}

//图片表
type Photo struct {
	Id          int64
	Title       string
	Created     time.Time `orm:"index"`
	Views       int64     `orm:"index"`
	Url         string
	PhotoSortId int64
}

//注册DB并连接数据库
func RegisterDB() {
	orm.RegisterModel(new(User), new(ArticleSort), new(Article), new(PhotoSort), new(Photo))
	orm.RegisterDataBase("default", "mysql", mysqlconfig(), 10)
	fmt.Println("数据库配置完毕！")
}

//---博客文章的数据库的相关操作---
//获取所有文章分类
func GetAllArticleSorts() ([]*ArticleSort, error) {
	o := orm.NewOrm()
	sorts := make([]*ArticleSort, 0)
	qs := o.QueryTable("ArticleSort")
	_, err := qs.All(&sorts)
	return sorts, err
}

//查询文章分类信息
func GetArticleSort(id int64) (ArticleSort, error) {
	o := orm.NewOrm()
	sort := ArticleSort{Id: id}
	err := o.Read(&sort)
	return sort, err
}

//添加文章分类
func AddArticleSort(name string) error {
	o := orm.NewOrm()
	sort := &ArticleSort{Title: name}
	sort.Created = time.Now()
	qs := o.QueryTable("ArticleSort")
	err := qs.Filter("title", name).One(sort)
	if err == nil {
		err.Error()
	}
	_, err = o.Insert(sort)
	if err != nil {
		err.Error()
	}
	return err
}

//删除文章分类
func DelArticleSort(id int64) error {
	o := orm.NewOrm()
	sort := &ArticleSort{Id: id}
	_, err := o.Delete(sort)
	return err
}

//获取所有文章
func GetAllArticles() ([]*Article, error) {
	o := orm.NewOrm()
	articles := make([]*Article, 0)
	_, err := o.QueryTable("Article").OrderBy("-Updated").All(&articles)
	return articles, err
}

//查询文章详细信息
func GetArticle(id int64) (Article, error) {
	o := orm.NewOrm()
	article := Article{Id: id}
	err := o.Read(&article)
	return article, err
}

//查询首页文章详细信息
func GetIndexArticle() (Article, error) {
	o := orm.NewOrm()
	var article Article
	err := o.QueryTable("Article").OrderBy("-Created").One(&article)
	return article, err
}

//添加文章
func AddArticle(Title string, Content string, ArticleSortId int64) error {
	o := orm.NewOrm()
	article := &Article{
		Title:         Title,
		Content:       Content,
		Created:       time.Now(),
		Updated:       time.Now(),
		ArticleSortId: ArticleSortId,
	}
	_, err := o.Insert(article)
	return err
}

//删除文章
func DelArticle(id int64) error {
	o := orm.NewOrm()
	article := &Article{Id: id}
	_, err := o.Delete(article)
	return err
}

//修改文章
func UpdateArticle(id int64, Title string, Content string, ArticleSortId int64) error {
	o := orm.NewOrm()
	article := &Article{
		Id:            id,
		Title:         Title,
		Content:       Content,
		ArticleSortId: ArticleSortId,
	}
	_, err := o.Update(article)
	return err
}

//---图片管理操作---
//获取图片分类
func GetAllPhotoSorts() ([]*PhotoSort, error) {
	o := orm.NewOrm()
	sorts := make([]*PhotoSort, 0)
	qs := o.QueryTable("PhotoSort")
	_, err := qs.All(&sorts)
	return sorts, err
}

//获取图片分类信息
func GetPhotoSort(id int64) (PhotoSort, error) {
	o := orm.NewOrm()
	sort := PhotoSort{Id: id}
	err := o.Read(&sort)
	return sort, err
}

//添加图片分类
func AddPhotoSort(name string) error {
	o := orm.NewOrm()
	sort := &PhotoSort{Title: name}
	sort.Created = time.Now()
	qs := o.QueryTable("PhotoSort")
	err := qs.Filter("title", name).One(sort)
	if err == nil {
		return err
	}
	_, err = o.Insert(sort)
	return err
}

//删除图片分类
func DelPhotoSort(id int64) error {
	o := orm.NewOrm()
	sort := &PhotoSort{Id: id}
	_, err := o.Delete(sort)
	return err
}

//更新图片分类统计数
func UpdataPhotoCount(id int64) {
	o := orm.NewOrm()
	sort := PhotoSort{Id: id}
	if o.Read(&sort) == nil {
		sort.Title = sort.Title
		sort.Created = sort.Created
		sort.Count = sort.Count + 1
		o.Update(&sort)
	}
}

//获取所有图片
func GetAllPhotos() ([]*Photo, error) {
	o := orm.NewOrm()
	photos := make([]*Photo, 0)
	_, err := o.QueryTable("Photo").All(&photos)
	return photos, err
}

//添加图片
func AddPhoto(PhotoTitle string, PhotoUrl string, sid int64) error {
	o := orm.NewOrm()
	photo := &Photo{
		Title:       PhotoTitle,
		Url:         PhotoUrl,
		Created:     time.Now(),
		PhotoSortId: sid,
	}
	_, err := o.Insert(photo)
	return err
}

//删除图片
func DelPhoto(id int64) error {
	o := orm.NewOrm()
	fmt.Println("调用DEL方法！")
	photo := &Photo{Id: id}
	_, err := o.Delete(photo)
	return err
}

//管理员查询
func Estimate(name string) string {
	o := orm.NewOrm()
	user := User{Username: name}
	err := o.Read(&user, "username")
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
		return "nil"
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
		return "nil"
	} else {
		fmt.Println(user.Id, user.Username)
		return user.Password
	}
}
