package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"leig/controllers"
	"leig/models"
)

func init() {
	//连接数据库
	models.RegisterDB()
}

func main() {
	//数据库建表
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	//注册路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.AutoRouter(&controllers.AdminController{})

	//运行站点
	beego.Run()
}
