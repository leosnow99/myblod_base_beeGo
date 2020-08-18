package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
	IsLogin   bool
	LoginUser interface{}
}

func (this *BaseController) Prepare() {
	loginuser := this.GetSession("loginuser")
	if loginuser != nil {
		this.IsLogin = true
		this.LoginUser = loginuser
	} else {
		this.IsLogin = false
	}
	this.Data["IsLogin"] = this.IsLogin
}
