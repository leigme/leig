package controllers

import (
"github.com/astaxie/beego"
"fmt"
)

type EditphotosortController struct {
	beego.Controller
}

func (this *EditphotosortController) Post() {
	fmt.Println(this.Input().Get("Photo"))
	fmt.Println(this.Input().Get("Url"))
	return
}

func (this *EditphotosortController) Get() {
	
}